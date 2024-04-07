package models

import (
	"rcbs/internal/hash"
	"rcbs/internal/mongo"
)

type User struct {
	mongo.DocWithTimestamps `bson:",inline"`

	EmailAddress      string `json:"email_address" bson:"email_address"`
	VerificationToken string `json:"verification_token" bson:"verification_token"`
	Verified          bool   `json:"verified" bson:"verified"`

	Username       string `json:"username" bson:"username"`
	Slug           string `json:"slug" bson:"slug"`
	HashedPassword string `json:"-" bson:"password"`

	Points              uint64   `json:"points" bson:"points"`
	ChallengesCompleted []uint64 `json:"challenges_completed" bson:"challenges_completed"`

	Roles []string `json:"roles" bson:"roles"`
}

func (u *User) SetPassword(password string) error {
	hashed, error := hash.GenerateFromPassword(password)
	if error != nil {
		return error
	}
	u.HashedPassword = string(hashed)
	return nil
}

func (u *User) VerifyPassword(password string) (bool, error) {
	return hash.ComparePasswordAndHash(password, u.HashedPassword)
}

func (u *User) BeforeInsert() error {
	return u.DocWithTimestamps.BeforeInsert()
}

func (u *User) BeforeUpdate() error {
	return u.DocWithTimestamps.BeforeUpdate()
}
