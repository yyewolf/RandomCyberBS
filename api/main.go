package api

import (
	"log"
	v1 "rcbs/api/v1"
	"time"

	"github.com/go-fuego/fuego"
	"github.com/sethvargo/go-limiter/httplimit"
	"github.com/sethvargo/go-limiter/memorystore"
)

func Setup(s *fuego.Server) {

	// Set group for api
	api := fuego.Group(s, "/api")

	// Middleware for rate limiting
	store, err := memorystore.New(&memorystore.Config{
		// Number of tokens allowed per interval.
		Tokens: 200,

		// Interval until tokens reset.
		Interval: time.Minute,
	})
	if err != nil {
		log.Fatal(err)
	}

	middleware, err := httplimit.NewMiddleware(store, httplimit.IPKeyFunc())
	if err != nil {
		log.Fatal(err)
	}

	fuego.Use(api, middleware.Handle)

	fuego.Get(api, "/health", Health)

	v1.Setup(api)
}
