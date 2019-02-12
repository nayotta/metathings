package metathings_tagd_storage

import (
	"context"
	"time"

	"github.com/mongodb/mongo-go-driver/bson"
	"github.com/mongodb/mongo-go-driver/mongo"
	"github.com/mongodb/mongo-go-driver/mongo/options"
	log "github.com/sirupsen/logrus"

	opt_helper "github.com/nayotta/metathings/pkg/common/option"
)

const (
	MONGO_TAG_ID  = "##id"
	MONGO_TAG_PAD = "x"
)

type MongoStorageOption struct {
	Uri        string
	Database   string
	Collection string
	Timeout    time.Duration
}

func NewMongoStorageOption() *MongoStorageOption {
	return &MongoStorageOption{
		Timeout: 10 * time.Second,
	}
}

type MongoStorage struct {
	client *mongo.Client
	logger log.FieldLogger

	opt *MongoStorageOption
}

func (self *MongoStorage) connect() error {
	var err error

	if self.client, err = mongo.Connect(self.context(), self.opt.Uri); err != nil {
		return err
	}

	return nil
}

func (self *MongoStorage) tag_filter_by_id(id string) bson.M {
	return bson.M{MONGO_TAG_ID: id}
}

func (self *MongoStorage) get_collection() *mongo.Collection {
	return self.client.Database(self.opt.Database).Collection(self.opt.Collection)
}

func (self *MongoStorage) context() context.Context {
	return context.TODO()
}

func (self *MongoStorage) get_tags_by_id(id string) ([]string, error) {
	exclude_fields := bson.M{
		MONGO_TAG_ID: 0,
		"_id":        0,
	}
	opts := &options.FindOptions{Projection: exclude_fields}
	cur, err := self.get_collection().Find(self.context(), self.tag_filter_by_id(id), opts)
	if err != nil {
		return nil, err
	}

	found := false
	var tags []string
	for cur.Next(self.context()) {
		var res bson.M
		if err = cur.Decode(&res); err != nil {
			return nil, err
		}

		for tag, _ := range res {
			tags = append(tags, tag)
		}
		found = true
	}
	if !found {
		return nil, ErrNotFound
	}

	return tags, nil
}

func (self *MongoStorage) Tag(id string, tags []string) error {
	m := bson.M{}
	m[MONGO_TAG_ID] = id
	for _, tag := range tags {
		m[tag] = MONGO_TAG_PAD
	}

	_, err := self.get_collection().InsertOne(self.context(), m)
	if err != nil {
		return err
	}

	return nil
}

func (self *MongoStorage) Untag(id string, tags []string) error {
	unsets := bson.M{}
	for _, tag := range tags {
		unsets[tag] = ""
	}
	update := bson.M{"$unset": unsets}

	_, err := self.get_collection().UpdateOne(self.context(), self.tag_filter_by_id(id), update)
	if err != nil {
		return err
	}

	return nil
}

func (self *MongoStorage) Remove(id string) error {
	if _, err := self.get_collection().DeleteOne(self.context(), self.tag_filter_by_id(id)); err != nil {
		return err
	}

	return nil
}

func (self *MongoStorage) Get(id string) ([]string, error) {
	tags, err := self.get_tags_by_id(id)
	if err != nil {
		return nil, err
	}
	return tags, nil
}

func (self *MongoStorage) Query(tags []string) ([]string, error) {
	if len(tags) == 0 {
		return nil, nil
	}

	query := bson.M{}
	for _, tag := range tags {
		query[tag] = MONGO_TAG_PAD
	}
	opts := &options.FindOptions{
		Projection: bson.M{MONGO_TAG_ID: 1},
	}

	cur, err := self.get_collection().Find(self.context(), query, opts)
	if err != nil {
		return nil, err
	}

	var ids []string
	for cur.Next(self.context()) {
		var res bson.M
		if err = cur.Decode(&res); err != nil {
			return nil, err
		}
		id, ok := res[MONGO_TAG_ID]
		if !ok {
			continue
		}
		ids = append(ids, id.(string))
	}

	return ids, nil
}

func NewMongoStorage(args ...interface{}) (Storage, error) {
	var ok bool
	var logger log.FieldLogger
	var err error
	opt := NewMongoStorageOption()

	if err = opt_helper.Setopt(map[string]func(string, interface{}) error{
		"uri":        opt_helper.ToString(&opt.Uri),
		"database":   opt_helper.ToString(&opt.Database),
		"collection": opt_helper.ToString(&opt.Collection),
		"timeout":    opt_helper.ToDuration(&opt.Timeout),
		"logger": func(key string, val interface{}) error {
			if logger, ok = val.(log.FieldLogger); !ok {
				return opt_helper.ErrInvalidArguments
			}
			return nil
		},
	})(args...); err != nil {
		return nil, err
	}
	stor := &MongoStorage{
		opt:    opt,
		logger: logger,
	}
	if err = stor.connect(); err != nil {
		return nil, err
	}

	return stor, nil
}

func init() {
	register_storage_factory("mongo", NewMongoStorage)
}