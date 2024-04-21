package auth

import (
	"rcbs/api/v1/middlewares"

	"github.com/go-fuego/fuego"
)

type AuthRessource struct {
	s  *fuego.Server
	mr *middlewares.MiddlewareRessource
}

func MountRoutes(s *fuego.Server) *AuthRessource {
	// Set group for users
	auth := fuego.Group(s, "/auth")

	ur := &AuthRessource{s: auth, mr: middlewares.New(s)}

	fuego.Post(auth, "/login", ur.LoginHandler(LoginFunc)).Tags("Auth").Summary("Login")
	fuego.Post(auth, "/register", ur.Register).Tags("Auth").Summary("Register")

	return ur
}
