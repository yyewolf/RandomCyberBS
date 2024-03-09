package models

import "rcbs/internal/mongo"

type User struct {
	mongo.DocWithTimestamps `bson:",inline"`

	Username string `json:"username" bson:"username"`
	Password string `json:"password" bson:"password"`
}

func (u *User) BeforeInsert() error {
	return u.DocWithTimestamps.BeforeInsert()
}

func (u *User) BeforeUpdate() error {
	return u.DocWithTimestamps.BeforeUpdate()
}
