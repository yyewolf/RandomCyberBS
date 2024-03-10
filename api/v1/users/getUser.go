package users

import (
	"rcbs/internal/messages"
	"rcbs/models"

	"github.com/go-fuego/fuego"
	"go.mongodb.org/mongo-driver/mongo"
)

type GetUserResponse struct {
	User *models.User `json:"user"`

	Status  string `json:"status"`
	Details string `json:"details"`
}

func GetUser(c *fuego.ContextNoBody) (*GetUserResponse, error) {
	id := c.PathParam("id")

	// Find users corresponding to the filters
	user, err := models.Db.Users.FindById(id)
	if err != nil {
		switch err {
		case mongo.ErrNoDocuments:
			return &GetUserResponse{
				Status:  "error",
				Details: messages.Get("user/get/not-found"),
			}, nil
		default:
			return &GetUserResponse{
				Status:  "error",
				Details: messages.Get("user/get/db-error"),
			}, err
		}
	}

	return &GetUserResponse{
		User: user,

		Status:  "ok",
		Details: messages.Get("user/get/ok"),
	}, nil
}
