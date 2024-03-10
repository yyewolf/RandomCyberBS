package models

import (
	"rcbs/internal/mongo"

	mongodb "go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Database struct {
	Users *mongo.Collection[*User]

	*mongo.Database
}

var Db *Database

func LoadDatabase() {
	Db = &Database{
		Users: mongo.GetCollection[*User](mongo.Db, "users"),

		Database: mongo.Db,
	}

	Db.Users.CreateIndex(mongodb.IndexModel{
		Keys:    map[string]int{"username": 1},
		Options: options.Index().SetUnique(true),
	})

}
