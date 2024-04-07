package models

import (
	"rcbs/internal/mongo"
)

type Solution struct {
	mongo.DocWithTimestamps `bson:",inline"`

	Author     string `json:"author" bson:"author"`
	AuthorName string `json:"author_name" bson:"author_name"`

	Challenge string `json:"challenge" bson:"challenge"`

	Description string   `json:"description" bson:"description"`
	Language    string   `json:"language" bson:"language"`
	Upvotes     int      `json:"upvotes" bson:"upvotes"`
	Tags        []string `json:"tags" bson:"tags"`

	// Administration purposes
	AwaitingConfirmation bool   `json:"awaiting_confirmation" bson:"awaiting_confirmation"`
	Confirmed            bool   `json:"confirmed" bson:"confirmed"`
	DenyReason           string `json:"deny_reason" bson:"deny_reason"`
}

func (u *Solution) BeforeInsert() error {
	return u.DocWithTimestamps.BeforeInsert()
}

func (u *Solution) BeforeUpdate() error {
	return u.DocWithTimestamps.BeforeUpdate()
}
