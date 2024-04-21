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

	fuego.Get(users, "/{id}", ur.GetUser).
		QueryParam("id", "string", fuego.OpenAPIParam{Required: true, Example: "1", Type: "path"})

	fuego.Get(users, "/", ur.GetUsers).
		QueryParam("name", "string", fuego.OpenAPIParam{Required: true, Example: "1", Type: "query"}).
		QueryParam("page", "int", fuego.OpenAPIParam{Required: true, Example: "1", Type: "query"}).
		QueryParam("perPage", "int", fuego.OpenAPIParam{Required: true, Example: "1", Type: "query"})

	fuego.Get(users, "/{id}/verify/{token}", ur.VerifyEmailAddress).
		QueryParam("id", "string", fuego.OpenAPIParam{Required: true, Example: "1", Type: "path"}).
		QueryParam("token", "string", fuego.OpenAPIParam{Required: true, Example: "fsdfsd", Type: "path"})

	return ur
}
