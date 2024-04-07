package models

import (
	"rcbs/internal/mongo"

	"go.mongodb.org/mongo-driver/bson"
)

type Challenge struct {
	mongo.DocWithTimestamps `bson:",inline"`

	Name        string `json:"name" bson:"name"`
	Slug        string `json:"slug" bson:"slug"`
	Difficulty  int    `json:"difficulty" bson:"difficulty"`
	Tags        string `json:"tags" bson:"tags"`
	Category    string `json:"category" bson:"category"`
	Points      int    `json:"points" bson:"points"`
	Description string `json:"description" bson:"description"`
	Flag        string `json:"flag" bson:"flag"`
	Upvotes     int    `json:"upvotes" bson:"upvotes"`

	AuthorName string `json:"author_name" bson:"author_name"`
	Author     string `json:"author" bson:"author"`

	// Files are shown, attachments are not but can be used in the description to display images
	Files       []File `json:"files" bson:"files"`
	Attachments []File `json:"attachments" bson:"attachments"`
}

func (c *Challenge) BeforeInsert() error {
	return c.DocWithTimestamps.BeforeInsert()
}

func (c *Challenge) BeforeUpdate() error {
	return c.DocWithTimestamps.BeforeUpdate()
}

func (c *Challenge) GetSolutions() ([]*Solution, error) {
	solutions, err := Db.Solutions.Find(
		bson.M{
			"challenge": c.ID,
		})
	return solutions, err
}
