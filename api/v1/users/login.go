package users

import (
	"rcbs/internal/auth"
	"rcbs/models"
	"time"

	"github.com/go-fuego/fuego"
	"github.com/golang-jwt/jwt/v5"
	"github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

func LoginFunc(username, password string) (jwt.Claims, error) {
	// Find user corresponding
	user, err := models.Db.Users.FindOne(bson.M{"username": username})
	if err != nil {
		switch err {
		case mongo.ErrNoDocuments:
			logrus.WithFields(logrus.Fields{
				"username": username,
			}).Debug("User not found")
		default:
			logrus.WithFields(logrus.Fields{
				"username": username,
			}).Error("Error finding user")
		}
		return auth.AuthClaims{}, fuego.ErrUnauthorized
	}

	return auth.AuthClaims{
		RegisteredClaims: jwt.RegisteredClaims{
			Issuer:    username,
			Subject:   username,
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(24 * time.Hour)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			ID:        user.ID,
		},
		Roles: user.Roles,
	}, nil
}
