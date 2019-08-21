package mongo_helper

import (
	"context"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type MongoClientWrapper struct {
	*mongo.Client
	cancel func()
}

func (w *MongoClientWrapper) Close() error {
	w.cancel()

	return nil
}

func NewMongoClient(uri string, opts ...*options.ClientOptions) (*MongoClientWrapper, error) {
	ctx := context.TODO()
	ctx, cancel := context.WithCancel(ctx)

	opts = append(opts, options.Client().ApplyURI(uri))
	client, err := mongo.Connect(ctx, opts...)
	if err != nil {
		cancel()
		return nil, err
	}

	return &MongoClientWrapper{
		Client: client,
		cancel: cancel,
	}, nil
}
