package v1

import (
	"rcbs/api/v1/middlewares"
	"rcbs/api/v1/users"

	"github.com/go-fuego/fuego"
)

func Setup(s *fuego.Server) {
	// Set group for v1
	v1 := fuego.Group(s, "/v1")

	// Middlewares
	mr := middlewares.New(s)

	// Auth
	fuego.Use(v1, mr.TokenFromContext)
	fuego.Post(v1, "/auth/login", s.Security.LoginHandler(users.LoginFunc)).Tags("Auth").Summary("Login")
	fuego.PostStd(v1, "/auth/logout", s.Security.CookieLogoutHandler).Tags("Auth").Summary("Logout")
	fuego.PostStd(v1, "/auth/refresh", s.Security.RefreshHandler).Tags("Auth").Summary("Refresh token")

	// Register users
	users.MountRoutes(v1)
}
