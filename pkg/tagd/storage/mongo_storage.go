package metathings_tagd_storage

import (
	"context"
	"time"

	log "github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"

	opt_helper "github.com/nayotta/metathings/pkg/common/option"
)

const (
	MONGO_TAG_ID  = "##id"
	MONGO_TAG_NS  = "##ns"
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

	if self.client, err = mongo.Connect(self.context(), options.Client().ApplyURI(self.opt.Uri)); err != nil {
		return err
	}

	return nil
}

func (self *MongoStorage) tag_filter_by_id(ns string, id string) bson.M {
	return bson.M{
		MONGO_TAG_ID: id,
		MONGO_TAG_NS: ns,
	}
}

func (self *MongoStorage) get_collection() *mongo.Collection {
	return self.client.Database(self.opt.Database).Collection(self.opt.Collection)
}

func (self *MongoStorage) context() context.Context {
	return context.TODO()
}

func (self *MongoStorage) get_tags_by_id(ns string, id string) ([]string, error) {
	ex, err := self.exist_tag(ns, id)
	if err != nil {
		return nil, err
	}

	if !ex {
		return nil, ErrNotFound
	}

	exclude_fields := bson.M{
		MONGO_TAG_ID: 0,
		MONGO_TAG_NS: 0,
		"_id":        0,
	}
	opts := &options.FindOptions{Projection: exclude_fields}
	cur, err := self.get_collection().Find(self.context(), self.tag_filter_by_id(ns, id), opts)
	if err != nil {
		return nil, err
	}

	var tags []string
	for cur.Next(self.context()) {
		var res bson.M
		if err = cur.Decode(&res); err != nil {
			return nil, err
		}

		for tag, _ := range res {
			tags = append(tags, tag)
		}
	}

	return tags, nil
}

func (self *MongoStorage) exist_tag(ns string, id string) (bool, error) {
	coll := self.get_collection()
	size, err := coll.CountDocuments(
		self.context(),
		bson.M{
			MONGO_TAG_ID: id,
			MONGO_TAG_NS: ns,
		},
		options.Count().SetLimit(1))
	if err != nil {
		return false, err
	}

	return size > 0, nil
}

func (self *MongoStorage) Tag(ns string, id string, tags []string) error {
	ex, err := self.exist_tag(ns, id)
	if err != nil {
		return err
	}

	if ex {
		flt := self.tag_filter_by_id(ns, id)
		doc := bson.M{}
		for _, tag := range tags {
			doc[tag] = MONGO_TAG_PAD
		}

		doc = bson.M{"$set": doc}
		_, err = self.get_collection().UpdateOne(self.context(), flt, doc)
		if err != nil {
			return err
		}
	} else {
		doc := self.tag_filter_by_id(ns, id)
		for _, tag := range tags {
			doc[tag] = MONGO_TAG_PAD
		}

		_, err = self.get_collection().InsertOne(self.context(), doc)
		if err != nil {
			return err
		}
	}

	return nil
}

func (self *MongoStorage) Untag(ns string, id string, tags []string) error {
	unsets := bson.M{}
	for _, tag := range tags {
		unsets[tag] = ""
	}
	doc := bson.M{"$unset": unsets}

	_, err := self.get_collection().UpdateOne(self.context(), self.tag_filter_by_id(ns, id), doc)
	if err != nil {
		return err
	}

	return nil
}

func (self *MongoStorage) Remove(ns string, id string) error {
	if _, err := self.get_collection().DeleteOne(self.context(), self.tag_filter_by_id(ns, id)); err != nil {
		return err
	}

	return nil
}

func (self *MongoStorage) Get(ns string, id string) ([]string, error) {
	tags, err := self.get_tags_by_id(ns, id)
	if err != nil {
		return nil, err
	}
	return tags, nil
}

func (self *MongoStorage) Query(ns string, tags []string) ([]string, error) {
	if len(tags) == 0 {
		return nil, nil
	}

	query := bson.M{}
	for _, tag := range tags {
		query[tag] = MONGO_TAG_PAD
	}
	query[MONGO_TAG_NS] = ns
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
	var logger log.FieldLogger
	var err error
	opt := NewMongoStorageOption()

	if err = opt_helper.Setopt(map[string]func(string, interface{}) error{
		"uri":        opt_helper.ToString(&opt.Uri),
		"database":   opt_helper.ToString(&opt.Database),
		"collection": opt_helper.ToString(&opt.Collection),
		"timeout":    opt_helper.ToDuration(&opt.Timeout),
		"logger":     opt_helper.ToLogger(&logger),
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
