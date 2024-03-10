package mongo

import (
	"go.mongodb.org/mongo-driver/mongo"
)

func (repo *Collection[T]) CreateIndex(mod mongo.IndexModel) error {
	_, err := repo.collection.Indexes().CreateOne(DefaultContext(), mod)

	return err
}
