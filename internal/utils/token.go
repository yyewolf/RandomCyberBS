package utils

import (
	"context"
	"errors"
	"rcbs/internal/auth"
	"rcbs/models"
)

type ContextKey string

const (
	TokenKey ContextKey = "jwt_claims"
	UserKey  ContextKey = "user"
)

// takes in *fuego.ContextNoBody or *fuego.ContextWithBody
type Context interface {
	Context() context.Context
}

func GetToken(c Context) (auth.AuthClaims, error) {
	claims := c.Context().Value(TokenKey)
	if claims == nil {
		return auth.AuthClaims{}, errors.New("no token found in context")
	}

	switch t := claims.(type) {
	case auth.AuthClaims:
		return t, nil
	default:
		return auth.AuthClaims{}, errors.New("invalid token in context")
	}
}

func GetUser(c Context) (*models.User, error) {
	user := c.Context().Value(UserKey)
	if user == nil {
		return nil, errors.New("no user found in context")
	}

	switch u := user.(type) {
	case *models.User:
		return u, nil
	default:
		return nil, errors.New("invalid user in context")
	}
}
