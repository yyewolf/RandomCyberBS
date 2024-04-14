package users

import (
	"rcbs/internal/auth"
	"rcbs/internal/httperrors"
	"rcbs/models"
	"time"

	"github.com/go-fuego/fuego"
	"github.com/golang-jwt/jwt/v5"
	"github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type tokenResponse struct {
	Token string `json:"token"`
}

func (ur *UserRessource) LoginHandler(verifyUserInfo func(user, password string) (jwt.Claims, error)) func(*fuego.ContextWithBody[fuego.LoginPayload]) (tokenResponse, error) {
	return func(c *fuego.ContextWithBody[fuego.LoginPayload]) (tokenResponse, error) {
		body, err := c.Body()
		if err != nil {
			return tokenResponse{}, err
		}

		claims, err := verifyUserInfo(body.User, body.Password)
		if err != nil {
			return tokenResponse{}, err
		}

		// Send the token to the cookies
		token, err := ur.s.Security.GenerateTokenToCookies(claims, c.Response())
		if err != nil {
			return tokenResponse{}, err
		}

		c.SetHeader("HX-Redirect", "/")

		// Send the token to the response
		return tokenResponse{
			Token: token,
		}, nil
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
		return auth.AuthClaims{}, httperrors.New(err, 401, "Invalid username or password")
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
