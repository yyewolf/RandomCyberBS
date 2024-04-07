package models

import "rcbs/internal/mongo"

type Upvote struct {
	mongo.DocWithTimestamps `bson:",inline"`

	User   string `json:"user" bson:"user"`
	Entity string `json:"entity" bson:"entity"`
}

func (u *Upvote) BeforeInsert() error {
	return u.DocWithTimestamps.BeforeInsert()
}

func (u *Upvote) BeforeUpdate() error {
	return u.DocWithTimestamps.BeforeUpdate()
}
