package middlewares

import "github.com/go-fuego/fuego"

type MiddlewareRessource struct {
	s *fuego.Server
}

func New(s *fuego.Server) *MiddlewareRessource {
	return &MiddlewareRessource{s}
}
