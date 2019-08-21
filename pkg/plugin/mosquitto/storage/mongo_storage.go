package metathings_moqusitto_plugin_storage

import (
	"context"
	"fmt"

	log "github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"

	mongo_helper "github.com/nayotta/metathings/pkg/common/mongo"
	opt_helper "github.com/nayotta/metathings/pkg/common/option"
	pool_helper "github.com/nayotta/metathings/pkg/common/pool"
)

type MongoStorageOption struct {
	Mongo struct {
		Uri        string
		Database   string
		Collection string
	}
	Pool struct {
		Initial int
		Max     int
	}
}

type MongoStorage struct {
	opt    *MongoStorageOption
	logger log.FieldLogger
	pool   pool_helper.Pool
}

func (s *MongoStorage) context() context.Context {
	return context.TODO()
}

func (s *MongoStorage) get_logger() log.FieldLogger {
	return s.logger
}

func (s *MongoStorage) get_client() (*mongo_helper.MongoClientWrapper, error) {
	cli, err := s.pool.Get()
	if err != nil {
		return nil, err
	}

	return cli.(*mongo_helper.MongoClientWrapper), nil
}

func (s *MongoStorage) get_collection() (*mongo.Collection, func(), error) {
	cli, err := s.get_client()
	if err != nil {
		s.get_logger().WithError(err).Debugf("failed to get mongo collection")
		return nil, nil, err
	}

	close := func() {
		s.pool.Put(cli)
	}

	db := cli.Database(s.opt.Mongo.Database)
	coll := db.Collection(s.opt.Mongo.Collection)

	return coll, close, nil
}

func (s *MongoStorage) add_user(coll *mongo.Collection, u *User) (err error) {
	m := bson.M{"superuser": false}

	m["username"] = *u.Username
	m["password"] = *u.Password
	if u.Superuser != nil {
		m["superuser"] = *u.Superuser
	}

	ts := bson.M{}
	for _, p := range u.Permissions {
		ts[*p.Topic] = *p.Mask
	}
	m["topics"] = ts

	_, err = coll.InsertOne(s.context(), m)
	return
}

func (s *MongoStorage) AddUser(u *User) error {
	coll, close, err := s.get_collection()
	if err != nil {
		return err
	}
	defer close()

	logger := s.get_logger().WithField("username", *u.Username)

	err = s.add_user(coll, u)
	if err != nil {
		logger.WithError(err).Debugf("failed to add user to mongo")
		return err
	}

	logger.Debugf("add user")
	return nil
}

func (s *MongoStorage) remove_user(coll *mongo.Collection, u string) (err error) {
	_, err = coll.DeleteOne(s.context(), bson.M{"username": u})
	return
}

func (s *MongoStorage) RemoveUser(u string) error {
	coll, close, err := s.get_collection()
	if err != nil {
		return err
	}
	defer close()

	logger := s.get_logger().WithField("username", u)

	err = s.remove_user(coll, u)
	if err != nil {
		logger.WithError(err).Debugf("failed to remove user from mongo")
		return err
	}

	logger.Debugf("remove user")
	return nil
}

func (s *MongoStorage) add_permission(coll *mongo.Collection, u string, p *Permission) (err error) {
	filter := bson.M{"username": u}
	update := bson.M{"$set": bson.M{fmt.Sprintf("topics.%v", *p.Topic): *p.Mask}}
	_, err = coll.UpdateOne(s.context(), filter, update)
	return
}

func (s *MongoStorage) AddPermission(u string, p *Permission) error {
	coll, close, err := s.get_collection()
	if err != nil {
		return err
	}
	defer close()

	logger := s.get_logger().WithField("username", u)

	err = s.add_permission(coll, u, p)
	if err != nil {
		logger.WithError(err).Debugf("failed to add permission to mongo")
		return err
	}

	logger.Debugf("add permission")
	return nil
}

func (s *MongoStorage) remove_permission(coll *mongo.Collection, u string, t string) (err error) {
	filter := bson.M{"username": u}
	update := bson.M{"$unset": bson.M{fmt.Sprintf("topics.%v", t): 1}}
	_, err = coll.UpdateOne(s.context(), filter, update)
	return
}

func (s *MongoStorage) RemovePermission(u string, t string) error {
	coll, close, err := s.get_collection()
	if err != nil {
		return err
	}
	defer close()

	logger := s.get_logger().WithField("username", u)

	err = s.remove_permission(coll, u, t)
	if err != nil {
		logger.WithError(err).Debugf("failed to remove permission from mongo")
		return err
	}

	logger.Debugf("remove permission")
	return nil
}

func (s *MongoStorage) get_user(coll *mongo.Collection, u string) (usr *User, err error) {
	res := coll.FindOne(s.context(), bson.M{"username": u})
	if err = res.Err(); err != nil {
		return
	}

	var tmp struct {
		Username  string
		Password  string
		Superuser bool
		Topics    map[string]string
	}
	if err = res.Decode(&tmp); err != nil {
		return
	}

	usr = &User{}
	usr.Username = &tmp.Username
	usr.Password = &tmp.Password
	usr.Superuser = &tmp.Superuser
	for topic, mask := range tmp.Topics {
		usr.Permissions = append(usr.Permissions, &Permission{
			Topic: &topic,
			Mask:  &mask,
		})
	}

	return
}

func (s *MongoStorage) GetUser(u string) (*User, error) {
	coll, close, err := s.get_collection()
	if err != nil {
		return nil, err
	}
	defer close()

	logger := s.get_logger().WithField("username", u)

	usr, err := s.get_user(coll, u)
	if err != nil {
		logger.WithError(err).Debugf("failed to get user from mongo")
		return nil, err
	}

	logger.Debugf("get user")
	return usr, nil
}

func (s *MongoStorage) exist_user(coll *mongo.Collection, u string) (existed bool, err error) {
	res := coll.FindOne(s.context(), bson.M{"username": u})
	if err = res.Err(); err != nil {
		return
	}

	if err = res.Decode(&struct{}{}); err != nil {
		if err == mongo.ErrNoDocuments {
			existed = false
			err = nil
		}
		return
	}

	existed = true
	return
}

func (s *MongoStorage) ExistUser(u string) (bool, error) {
	coll, close, err := s.get_collection()
	if err != nil {
		return false, err
	}
	defer close()

	logger := s.get_logger().WithField("username", u)

	existed, err := s.exist_user(coll, u)
	if err != nil {
		logger.WithError(err).Debugf("failed to check user exists in mongo")
		return false, err
	}

	logger.Debugf("check user exists in mongo")
	return existed, nil
}

type MongoStorageFactory struct{}

func (*MongoStorageFactory) New(args ...interface{}) (Storage, error) {
	var logger log.FieldLogger
	opt := &MongoStorageOption{}
	opt.Pool.Initial = 5
	opt.Pool.Max = 17

	if err := opt_helper.Setopt(map[string]func(string, interface{}) error{
		"logger":     opt_helper.ToLogger(&logger),
		"uri":        opt_helper.ToString(&opt.Mongo.Uri),
		"database":   opt_helper.ToString(&opt.Mongo.Database),
		"collection": opt_helper.ToString(&opt.Mongo.Collection),
	})(args...); err != nil {
		return nil, err
	}

	pool, err := pool_helper.NewPool(opt.Pool.Initial, opt.Pool.Max, func() (pool_helper.Client, error) {
		return mongo_helper.NewMongoClient(opt.Mongo.Uri)
	})
	if err != nil {
		return nil, err
	}

	return &MongoStorage{
		opt:    opt,
		logger: logger,
		pool:   pool,
	}, nil
}

func init() {
	register_storage_factory("mongo", new(MongoStorageFactory))
}
