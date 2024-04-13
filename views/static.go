package views

import (
	"net/http"
	"rcbs/static"
)

func robots(w http.ResponseWriter, r *http.Request) {
	http.ServeFileFS(w, r, static.StaticFiles, "robots.txt")
}

func tailwind(w http.ResponseWriter, r *http.Request) {
	http.ServeFileFS(w, r, static.StaticFiles, "tailwind.min.css")
}
