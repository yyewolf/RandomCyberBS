package httperrors

import (
	"net/http"

	"github.com/go-fuego/fuego"
)

func New(err error, status int, detail string) fuego.HTTPError {
	return fuego.HTTPError{
		Err:    err,
		Detail: detail,
		Title:  http.StatusText(status),
		Status: status,
	}
}
