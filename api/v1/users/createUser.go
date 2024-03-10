package users

import (
	"rcbs/internal/messages"
	"rcbs/models"

	"github.com/go-fuego/fuego"
	"github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/mongo"
)

type CreateUserRequest struct {
	Username string `json:"username" validate:"required"`
	Password string `json:"password" validate:"required"`
}

type CreateUserResponse struct {
	Status  string `json:"status"`
	Details string `json:"details"`
}

func CreateUser(c *fuego.ContextWithBody[*CreateUserRequest]) (*CreateUserResponse, error) {
	body, err := c.Body()
	if err != nil {
		return &CreateUserResponse{
			Status:  "error",
			Details: messages.Get("user/create/bad-request"),
		}, err
	}

	logrus.WithFields(logrus.Fields{
		"username": body.Username,
	}).Info("Creating user")

	// Create user in database
	user := &models.User{
		Username: body.Username,
	}

	user.SetPassword(body.Password)

	_, err = models.Db.Users.Insert(user)
	if err != nil {
		switch {
		case mongo.IsDuplicateKeyError(err):

			logrus.WithFields(logrus.Fields{
				"username": body.Username,
			}).Debug("User already exists")

			return &CreateUserResponse{
				Status:  "error",
				Details: messages.Get("user/create/already-exists"),
			}, nil
		default:

			logrus.WithFields(logrus.Fields{
				"error": err,
			}).Error("Error creating user")

			return &CreateUserResponse{
				Status:  "error",
				Details: messages.Get("user/create/db-error"),
			}, err
		}
	}

	return &CreateUserResponse{Status: "ok", Details: messages.Get("user/create/ok")}, nil
}
