package users

import (
	"rcbs/api/v1/controllers/users"
	"rcbs/internal/messages"
	"rcbs/models"

	"github.com/go-fuego/fuego"
)

type GetUsersRequest struct {
	// Filters
	Username string `json:"username"`

	Page    int `json:"page" validate:"min=1"`
	PerPage int `json:"per_page" validate:"max=100,min=1"`
}

type GetUsersResponse struct {
	Users   []*models.User `json:"users"`
	Page    int            `json:"page"`
	PerPage int            `json:"per_page"`
	MaxPage int            `json:"max_page"`

	Status  string `json:"status"`
	Details string `json:"details"`
}

func (ur *UserRessource) GetUsers(c *fuego.ContextNoBody) (*GetUsersResponse, error) {
	// Convert the query parameters to a GetUsersRequest
	req := &GetUsersRequest{
		Username: c.QueryParam("username"),
		Page:     c.QueryParamInt("page", 1),
		PerPage:  c.QueryParamInt("per_page", 10),
	}

	r, err := users.GetUsers(req.Username, req.Page, req.PerPage)
	if err != nil {
		return &GetUsersResponse{
			Status:  "error",
			Details: messages.Get("user/list/db-error"),
		}, err
	}

	return &GetUsersResponse{
		Users:   r.Users,
		Page:    r.Page,
		PerPage: r.PerPage,
		MaxPage: r.MaxPage,

		Status:  "ok",
		Details: messages.Get("user/list/ok"),
	}, nil
}
