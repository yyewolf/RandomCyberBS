package models

import (
	"rcbs/internal/mongo"

	mongodb "go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Database struct {
	Users      *mongo.Collection[*User]
	Challenges *mongo.Collection[*Challenge]
	Solutions  *mongo.Collection[*Solution]
	Upvotes    *mongo.Collection[*Upvote]

	*mongo.Database
}

var Db *Database

func LoadDatabase() {
	Db = &Database{
		Users:      mongo.GetCollection[*User](mongo.Db, "users"),
		Challenges: mongo.GetCollection[*Challenge](mongo.Db, "challenges"),
		Solutions:  mongo.GetCollection[*Solution](mongo.Db, "solutions"),
		Upvotes:    mongo.GetCollection[*Upvote](mongo.Db, "upvotes"),

		Database: mongo.Db,
	}

	// The username should be unique
	Db.Users.CreateIndex(mongodb.IndexModel{
		Keys:    map[string]int{"username": 1},
		Options: options.Index().SetUnique(true),
	})

	// The email address should be unique
	Db.Users.CreateIndex(mongodb.IndexModel{
		Keys:    map[string]int{"email_address": 1},
		Options: options.Index().SetUnique(true),
	})

	// The combo "user" + "entity" should be unique
	Db.Upvotes.CreateIndex(mongodb.IndexModel{
		Keys:    map[string]int{"user": 1, "entity": 1},
		Options: options.Index().SetUnique(true),
	})

}
