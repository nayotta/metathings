package mongo_helper

import (
	"context"

	"github.com/mongodb/mongo-go-driver/mongo"
	"github.com/mongodb/mongo-go-driver/mongo/options"
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

	client, err := mongo.Connect(ctx, uri, opts...)
	if err != nil {
		cancel()
		return nil, err
	}

	return &MongoClientWrapper{
		Client: client,
		cancel: cancel,
	}, nil
}
