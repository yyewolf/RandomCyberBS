package users

import "github.com/go-fuego/fuego"

func Setup(s *fuego.Server) {
	// Set group for users
	users := fuego.Group(s, "/users")

	fuego.Get(users, "/{id}", GetUser).
		QueryParam("id", "string", fuego.OpenAPIParam{Required: true, Example: "user", Type: "path"})

	fuego.Get(users, "/", GetUsers).
		QueryParam("page", "int", fuego.OpenAPIParam{Required: false, Example: "1", Type: "query"}).
		QueryParam("per_page", "int", fuego.OpenAPIParam{Required: false, Example: "10", Type: "query"}).
		QueryParam("username", "string", fuego.OpenAPIParam{Required: false, Example: "user", Type: "query"})

	fuego.Post(users, "/", CreateUser)
}
