package models

import "rcbs/internal/mongo"

type Database struct {
	Users *mongo.Collection[*User]
}

var Db *Database

func LoadDatabase() {
	Db = &Database{
		Users: mongo.GetCollection[*User](mongo.Db, "users"),
	}
}
