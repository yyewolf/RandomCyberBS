package messages

func Get(key string) string {
	return msgs[key]
}

var msgs = map[string]string{
	// Create user
	"user/create/already-exists": "User already exists",
	"user/create/bad-request":    "Bad request",
	"user/create/db-error":       "An error occurred while creating the user",
	"user/create/ok":             "User created",

	"user/list/bad-request": "Bad request",
	"user/list/db-error":    "An error occurred while listing users",
	"user/list/ok":          "Users listed successfully",

	"user/get/not-found": "User not found",
	"user/get/db-error":  "An error occurred while listing users",
	"user/get/ok":        "Found user",
}
