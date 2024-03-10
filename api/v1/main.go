package v1

import (
	"rcbs/api/v1/users"

	"github.com/go-fuego/fuego"
)

func Setup(s *fuego.Server) {
	// Set group for v1
	v1 := fuego.Group(s, "/v1")

	// Register users
	users.Setup(v1)
}
