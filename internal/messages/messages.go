package messages

func Get(key string) string {
	return msgs[key]
}

var msgs = map[string]string{
	// Create user
	"register/already-exists": "User already exists",
	"register/bad-request":    "Bad request",
	"register/db-error":       "An error occurred while creating the user",
	"register/ok":             "User created",

	"user/list/bad-request": "Bad request",
	"user/list/db-error":    "An error occurred while listing users",
	"user/list/ok":          "Users listed successfully",

	"user/get/not-found": "User not found",
	"user/get/db-error":  "An error occurred while listing users",
	"user/get/ok":        "Found user",

	"user/verify/db-error":         "An error occurred while verifying the user",
	"user/verify/not-found":        "User not found",
	"user/verify/wrong-code":       "Wrong verification code",
	"user/verify/already-verified": "User already verified",
	"user/verify/ok":               "User verified",

	"login/error": "An error occurred while logging in",
	"login/wrong": "Username or password is incorrect",
}
