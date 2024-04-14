package users

import (
	"rcbs/api/v1/middlewares"

	"github.com/go-fuego/fuego"
)

type UserRessource struct {
	s  *fuego.Server
	mr *middlewares.MiddlewareRessource
}

func MountRoutes(s *fuego.Server) *UserRessource {
	// Set group for users
	users := fuego.Group(s, "/users")

	ur := &UserRessource{s: users, mr: middlewares.New(s)}

	fuego.Post(s, "/auth/login", ur.LoginHandler(LoginFunc)).Tags("Auth").Summary("Login")

	fuego.Get(users, "/{id}", ur.GetUser).
		QueryParam("id", "string", fuego.OpenAPIParam{Required: true, Example: "1", Type: "path"})

	fuego.Get(users, "/{id}/verify/{token}", ur.VerifyEmailAddress).
		QueryParam("id", "string", fuego.OpenAPIParam{Required: true, Example: "1", Type: "path"}).
		QueryParam("token", "string", fuego.OpenAPIParam{Required: true, Example: "fsdfsd", Type: "path"})

	// fuego.Get(users, "/", ur.GetUsers, ur.mr.AuthWall("admin")).
	fuego.Get(users, "/", ur.GetUsers).
		QueryParam("page", "int", fuego.OpenAPIParam{Required: false, Example: "1", Type: "query"}).
		QueryParam("per_page", "int", fuego.OpenAPIParam{Required: false, Example: "10", Type: "query"}).
		QueryParam("username", "string", fuego.OpenAPIParam{Required: false, Example: "user", Type: "query"})

	fuego.Post(users, "/", ur.Register)

	return ur
}
