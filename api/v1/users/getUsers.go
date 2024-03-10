package users

import (
	"rcbs/internal/messages"
	"rcbs/models"

	"github.com/go-fuego/fuego"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
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

func GetUsers(c *fuego.ContextNoBody) (*GetUsersResponse, error) {
	// Convert the query parameters to a GetUsersRequest
	req := &GetUsersRequest{
		Username: c.QueryParam("username"),
		Page:     c.QueryParamInt("page", 1),
		PerPage:  c.QueryParamInt("per_page", 10),
	}

	// Make a filter for the username
	filter := bson.M{
		"username": bson.M{"$regex": req.Username, "$options": "i"},
	}

	// Find users corresponding to the filters
	users, err := models.Db.Users.Find(
		filter,
		options.Find().SetSkip(int64((req.Page-1)*req.PerPage)).SetLimit(int64(req.PerPage)),
	)
	if err != nil {
		return &GetUsersResponse{
			Status:  "error",
			Details: messages.Get("user/list/db-error"),
		}, err
	}

	// Count the total number of users corresponding to the filters
	total, err := models.Db.Users.CountDocuments(filter)
	if err != nil {
		return &GetUsersResponse{
			Status:  "error",
			Details: messages.Get("user/list/db-error"),
		}, err
	}

	// Calculate the max page
	maxPage := total / int64(req.PerPage)
	if total%int64(req.PerPage) != 0 {
		maxPage++
	}

	return &GetUsersResponse{
		Users:   users,
		Page:    req.Page,
		PerPage: req.PerPage,
		MaxPage: int(maxPage),

		Status:  "ok",
		Details: messages.Get("user/list/ok"),
	}, nil
}
