package main

import (
	"rcbs/internal/env"
	"rcbs/internal/mongo"
)

func main() {
	// Load environment variable
	env.Load()

	// Connect to MongoDB
	mongo.Connect()
}
