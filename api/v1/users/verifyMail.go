package users

import (
	"rcbs/internal/messages"
	"rcbs/models"

	"github.com/go-fuego/fuego"
	"go.mongodb.org/mongo-driver/mongo"
)

type VerifyEmailAddressResponse struct {
	Status  string `json:"status"`
	Details string `json:"details"`
}

func (ur *UserRessource) VerifyEmailAddress(c *fuego.ContextNoBody) (*VerifyEmailAddressResponse, error) {
	id := c.PathParam("id")
	token := c.PathParam("token")

	// Find users corresponding to the filters
	user, err := models.Db.Users.FindById(id)
	if err != nil {
		switch err {
		case mongo.ErrNoDocuments:
			return &VerifyEmailAddressResponse{
				Status:  "error",
				Details: messages.Get("user/verify/not-found"),
			}, nil
		default:
			return &VerifyEmailAddressResponse{
				Status:  "error",
				Details: messages.Get("user/get/db-error"),
			}, err
		}
	}

	if user.Verified {
		return &VerifyEmailAddressResponse{
			Status:  "error",
			Details: messages.Get("user/verify/already-verified"),
		}, nil
	}

	if user.VerificationToken != token {
		return &VerifyEmailAddressResponse{
			Status:  "error",
			Details: messages.Get("user/verify/wrong-code"),
		}, nil
	}

	user.Verified = true
	err = models.Db.Users.UpdateById(user.ID, user)
	if err != nil {
		return &VerifyEmailAddressResponse{
			Status:  "error",
			Details: messages.Get("user/get/db-error"),
		}, err
	}

	return &VerifyEmailAddressResponse{
		Status:  "ok",
		Details: messages.Get("user/verify/ok"),
	}, nil
}
