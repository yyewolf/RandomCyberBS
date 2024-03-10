package api

import "github.com/go-fuego/fuego"

type HealthResponse struct {
	Status string `json:"status"`
}

func Health(c *fuego.ContextNoBody) (*HealthResponse, error) {
	return &HealthResponse{Status: "ok"}, nil
}
