package mongo

import (
	"context"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
)

type Database struct {
	db     *mongo.Database
	client *mongo.Client
}

func (db *Database) Disconnect() error {
	err := db.client.Disconnect(DefaultContext())
	db.db = nil
	return err
}

func DefaultContext() context.Context {
	ctx, _ := context.WithTimeout(context.Background(), 5*time.Second)
	return ctx
}

func GetCollection[T Document](db *Database, collectionName string) *Collection[T] {
	return &Collection[T]{db.db.Collection(collectionName)}
}
