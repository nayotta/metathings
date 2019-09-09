package metathings_identityd2_policy

import (
	"context"

	log "github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"

	mongo_helper "github.com/nayotta/metathings/pkg/common/mongo"
	opt_helper "github.com/nayotta/metathings/pkg/common/option"
	storage "github.com/nayotta/metathings/pkg/identityd2/storage"
)

type MongoBackendCacheOption struct {
	MongoUri        string
	MongoDatabase   string
	MongoCollection string
}

type MongoBackendCache struct {
	opt      *MongoBackendCacheOption
	mgo_coll *mongo.Collection
	logger   log.FieldLogger
}

func (mbc *MongoBackendCache) context() context.Context {
	return context.TODO()
}

func (mbc *MongoBackendCache) get_logger() log.FieldLogger {
	return mbc.logger
}

func (mbc *MongoBackendCache) to_bson(sub, obj *storage.Entity, act *storage.Action) bson.M {
	m := bson.M{}
	if sub != nil {
		m["subject"] = ConvertSubject(sub)
	}
	if obj != nil {
		m["object"] = ConvertObject(obj)
	}
	if act != nil {
		m["action"] = ConvertAction(act)
	}
	return m
}

func (mbc *MongoBackendCache) Get(sub, obj *storage.Entity, act *storage.Action) (ret bool, err error) {
	doc := mbc.to_bson(sub, obj, act)
	sr := mbc.mgo_coll.FindOne(mbc.context(), doc)

	err = sr.Err()
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return false, ErrNoCached
		} else {
			return false, err
		}
	}

	err = sr.Decode(&struct{ Result *bool }{&ret})
	return
}

func (mbc *MongoBackendCache) Set(sub, obj *storage.Entity, act *storage.Action, ret bool) (err error) {
	flt := mbc.to_bson(sub, obj, act)
	doc := mbc.to_bson(sub, obj, act)
	doc["result"] = ret

	_, err = mbc.mgo_coll.ReplaceOne(mbc.context(), flt, doc, options.Replace().SetUpsert(true))
	return
}

func (mbc *MongoBackendCache) Remove(vals ...interface{}) (err error) {
	siz := len(vals)
	if siz%2 != 0 {
		return ErrInvalidArguments
	}

	flt := bson.M{}
	for i := 0; i < siz; i += 2 {
		key, ok := vals[i].(string)
		if !ok {
			return ErrInvalidArguments
		}
		switch key {
		case "subject":
			flt[key] = ConvertSubject(vals[i+1].(*storage.Entity))
		case "object":
			flt[key] = ConvertObject(vals[i+1].(*storage.Entity))
		case "action":
			flt[key] = ConvertAction(vals[i+1].(*storage.Action))
		}
	}

	_, err = mbc.mgo_coll.DeleteMany(mbc.context(), flt)
	return
}

type MongoBackendCacheFactory struct{}

func (*MongoBackendCacheFactory) New(args ...interface{}) (BackendCache, error) {
	var opt MongoBackendCacheOption
	var logger log.FieldLogger

	if err := opt_helper.Setopt(map[string]func(string, interface{}) error{
		"mongo_uri":        opt_helper.ToString(&opt.MongoUri),
		"mongo_database":   opt_helper.ToString(&opt.MongoDatabase),
		"mongo_collection": opt_helper.ToString(&opt.MongoCollection),
		"logger":           opt_helper.ToLogger(&logger),
	})(args...); err != nil {
		return nil, err
	}

	mgo_cli, err := mongo_helper.NewMongoClient(opt.MongoUri)
	if err != nil {
		return nil, err
	}

	coll := mgo_cli.Database(opt.MongoDatabase).Collection(opt.MongoCollection)

	mbc := &MongoBackendCache{
		opt:      &opt,
		mgo_coll: coll,
		logger:   logger,
	}

	return mbc, nil
}

func init() {
	register_backend_cache_factory("mongo", new(MongoBackendCacheFactory))
}
