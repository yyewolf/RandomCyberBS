package models

import (
	"rcbs/internal/hash"
	"rcbs/internal/mongo"
)

type User struct {
	mongo.DocWithTimestamps `bson:",inline"`

	Username string `json:"username" bson:"username"`
	Password string `json:"-" bson:"password"`
}

func (u *User) SetPassword(password string) error {
	hashed, error := hash.GenerateFromPassword(password)
	if error != nil {
		return error
	}
	u.Password = string(hashed)
	return nil
}

func (u *User) VerifyPassword(password string) (bool, error) {
	return hash.ComparePasswordAndHash(password, u.Password)
}

func (u *User) BeforeInsert() error {
	return u.DocWithTimestamps.BeforeInsert()
}

func (u *User) BeforeUpdate() error {
	return u.DocWithTimestamps.BeforeUpdate()
}
