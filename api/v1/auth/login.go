package auth

import (
	"net/http"
	"rcbs/internal/auth"
	"rcbs/internal/httperrors"
	"rcbs/internal/messages"
	"rcbs/models"
	viewAuth "rcbs/templa/api/v1/auth"
	"time"

	"github.com/go-fuego/fuego"
	"github.com/golang-jwt/jwt/v5"
	"github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type LoginPayload struct {
	User     string `json:"user"` // Might be an email, a username, or anything else that identifies uniquely the user
	Password string `json:"password"`
}

func (ur *AuthRessource) LoginHandler(verifyUserInfo func(user, password string) (jwt.Claims, error)) func(*fuego.ContextWithBody[LoginPayload]) (fuego.Templ, error) {
	return func(c *fuego.ContextWithBody[LoginPayload]) (fuego.Templ, error) {
		body, err := c.Body()
		if err != nil {
			logrus.WithFields(logrus.Fields{
				"error": err,
			}).Debug("Could not read body")
			return viewAuth.LoginForm(&viewAuth.LoginError{Message: messages.Get("login/error")}), nil
		}

		claims, err := verifyUserInfo(body.User, body.Password)
		if err != nil {
			return viewAuth.LoginForm(&viewAuth.LoginError{Message: messages.Get("login/wrong")}), nil
		}

		// Send the token to the cookies
		_, err = ur.s.Security.GenerateTokenToCookies(claims, c.Response())
		if err != nil {
			return viewAuth.LoginForm(&viewAuth.LoginError{Message: messages.Get("login/error")}), nil
		}

		c.SetHeader("HX-Redirect", "/")

		// Send the token to the response
		return nil, nil
	}
}

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
		return auth.AuthClaims{}, httperrors.New(err, http.StatusUnauthorized, "Invalid username or password")
	}

	if ok, err := user.VerifyPassword(password); !ok || err != nil {
		return auth.AuthClaims{}, httperrors.New(err, http.StatusUnauthorized, "Invalid username or password")
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
