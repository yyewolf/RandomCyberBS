package mongo

import (
	"context"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
)

type database struct {
	db     *mongo.Database
	client *mongo.Client
}

func (db *database) Disconnect() error {
	err := db.client.Disconnect(DefaultContext())
	db.db = nil
	return err
}

func DefaultContext() context.Context {
	ctx, _ := context.WithTimeout(context.Background(), 5*time.Second)
	return ctx
}

func GetCollection[T Document](db *database, collectionName string) *Collection[T] {
	return &Collection[T]{db.db.Collection(collectionName)}
}
