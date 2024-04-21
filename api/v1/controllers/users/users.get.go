package users

import (
	"rcbs/models"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type GetUsersControllers struct {
	Users   []*models.User `json:"users"`
	Page    int            `json:"page" validate:"min=1"`
	PerPage int            `json:"per_page" validate:"max=100,min=1"`
	MaxPage int            `json:"max_page"`
}

func GetUsers(username string, page, perPage int) (GetUsersControllers, error) {
	// Make a filter for the username
	filter := bson.M{
		"username": bson.M{"$regex": username, "$options": "i"},
	}

	if perPage > 100 {
		perPage = 100
	}
	if perPage < 1 {
		perPage = 1
	}

	// Count the total number of users corresponding to the filters
	total, err := models.Db.Users.CountDocuments(filter)
	if err != nil {
		return GetUsersControllers{}, err
	}

	// Calculate the max page
	maxPage := total / int64(perPage)
	if total%int64(perPage) != 0 {
		maxPage++
	}

	if page > int(maxPage) {
		page = int(maxPage)
	}

	// Find users corresponding to the filters
	users, err := models.Db.Users.Find(
		filter,
		options.Find().SetSkip(int64((page-1)*perPage)).SetLimit(int64(perPage)),
	)
	if err != nil {
		return GetUsersControllers{}, err
	}

	return GetUsersControllers{
		Users:   users,
		Page:    page,
		PerPage: perPage,
		MaxPage: int(maxPage),
	}, err
}
