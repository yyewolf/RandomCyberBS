package auth

import (
	"rcbs/internal/mail"
	"rcbs/internal/messages"
	"rcbs/models"

	viewAuth "rcbs/templa/api/v1/auth"

	"github.com/go-fuego/fuego"
	"github.com/google/uuid"
	"github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/mongo"
)

type RegisterPayload struct {
	EmailAddress    string `json:"email_address" validate:"required,email"`
	Username        string `json:"username" validate:"required"`
	Password        string `json:"password" validate:"required"`
	PasswordConfirm string `json:"password_confirm" validate:"required"`
}

func (ur *AuthRessource) Register(c *fuego.ContextWithBody[RegisterPayload]) (fuego.Templ, error) {
	body, err := c.Body()
	if err != nil {
		logrus.WithFields(logrus.Fields{
			"error": err,
		}).Debug("Could not read body")
		return viewAuth.RegisterForm(&viewAuth.RegisterError{Message: messages.Get("register/bad-request")}), nil
	}

	logrus.WithFields(logrus.Fields{
		"username": body.Username,
	}).Debug("Creating user")

	if body.Password != body.PasswordConfirm {
		logrus.WithFields(logrus.Fields{
			"username": body.Username,
		}).Debug("Bad confirmation password")
		return viewAuth.RegisterForm(&viewAuth.RegisterError{Message: messages.Get("register/bad-request")}), nil
	}

	// Create user in database
	user := &models.User{
		Username:          body.Username,
		EmailAddress:      body.EmailAddress,
		VerificationToken: uuid.New().String(),
	}

	user.SetPassword(body.Password)

	_, err = models.Db.Users.Insert(user)
	if err != nil {
		switch {
		case mongo.IsDuplicateKeyError(err):

			logrus.WithFields(logrus.Fields{
				"username": body.Username,
			}).Debug("User already exists")

			return viewAuth.RegisterForm(&viewAuth.RegisterError{Message: messages.Get("register/already-exists")}), nil
		default:

			logrus.WithFields(logrus.Fields{
				"error": err,
			}).Error("Error creating user")

			return viewAuth.RegisterForm(&viewAuth.RegisterError{Message: messages.Get("register/db-error")}), nil
		}
	}

	go mail.SendWelcomeMail(user)

	logrus.WithFields(logrus.Fields{
		"username": body.Username,
		"id":       user.ID,
		"token":    user.VerificationToken,
	}).Debug("User created")

	c.SetHeader("HX-Redirect", "/login")

	return viewAuth.RegisterForm(nil), nil
}
