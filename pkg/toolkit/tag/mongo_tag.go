package metathings_toolkit_tag

import (
	"context"
	"time"

	"github.com/mongodb/mongo-go-driver/bson"
	"github.com/mongodb/mongo-go-driver/mongo"
	"github.com/mongodb/mongo-go-driver/mongo/options"

	opt_helper "github.com/nayotta/metathings/pkg/common/option"
)

const (
	MONGO_TAG_ID  = "##id"
	MONGO_TAG_PAD = "x"
)

type MongoTagToolkitOption struct {
	Uri        string
	Database   string
	Collection string
	Timeout    time.Duration
}

func NewMongoTagToolkitOption() *MongoTagToolkitOption {
	return &MongoTagToolkitOption{
		Timeout: 10 * time.Second,
	}
}

type MongoTagToolkit struct {
	client *mongo.Client

	opt *MongoTagToolkitOption
}

func (self *MongoTagToolkit) connect() error {
	var err error

	if self.client, err = mongo.Connect(self.context(), self.opt.Uri); err != nil {
		return err
	}

	return nil
}

func (self *MongoTagToolkit) tag_filter_by_id(id string) bson.M {
	return bson.M{MONGO_TAG_ID: id}
}

func (self *MongoTagToolkit) get_collection() *mongo.Collection {
	return self.client.Database(self.opt.Database).Collection(self.opt.Collection)
}

func (self *MongoTagToolkit) context() context.Context {
	return context.TODO()
}

func (self *MongoTagToolkit) get_tags_by_id(id string) ([]string, error) {
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

func (self *MongoTagToolkit) Tag(id string, tags []string) error {
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

func (self *MongoTagToolkit) Untag(id string, tags []string) error {
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

func (self *MongoTagToolkit) Remove(id string) error {
	if _, err := self.get_collection().DeleteOne(self.context(), self.tag_filter_by_id(id)); err != nil {
		return err
	}

	return nil
}

func (self *MongoTagToolkit) Get(id string) ([]string, error) {
	tags, err := self.get_tags_by_id(id)
	if err != nil {
		return nil, err
	}
	return tags, nil
}

func (self *MongoTagToolkit) Query(tags []string) ([]string, error) {
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

func NewMongoTagToolkit(args ...interface{}) (TagToolkit, error) {
	var err error
	opt := NewMongoTagToolkitOption()

	if err = opt_helper.Setopt(map[string]func(string, interface{}) error{
		"uri":        opt_helper.ToString(&opt.Uri),
		"database":   opt_helper.ToString(&opt.Database),
		"collection": opt_helper.ToString(&opt.Collection),
		"timeout":    opt_helper.ToDuration(&opt.Timeout),
	})(args...); err != nil {
		return nil, err
	}

	tagtk := &MongoTagToolkit{opt: opt}
	if err = tagtk.connect(); err != nil {
		return nil, err
	}

	return tagtk, nil
}

func init() {
	register_tag_toolkit_factory("mongo", NewMongoTagToolkit)
}
