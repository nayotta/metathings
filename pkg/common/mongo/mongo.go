package mongo_helper

import (
	"context"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type MongoClientWrapper struct {
	*mongo.Client
}

func (w *MongoClientWrapper) Close() error {
	return w.Disconnect(context.TODO())
}

func NewMongoClient(uri string, opts ...*options.ClientOptions) (*MongoClientWrapper, error) {
	opts = append(opts, options.Client().ApplyURI(uri))
	client, err := mongo.Connect(context.TODO(), opts...)
	if err != nil {
		return nil, err
	}

	return &MongoClientWrapper{
		Client: client,
	}, nil
}
