package mongo

import (
	"context"
	"fmt"
	"rcbs/internal/env"
	"rcbs/internal/values"

	"github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var Db *Database

func Connect() {
	// Build the connection string
	cfg := env.Get()
	uri := fmt.Sprintf("mongodb://%s:%s/?%s", cfg.Mongo.Host, cfg.Mongo.Port, cfg.Mongo.Additional)

	// Set client options
	opts := options.Client().
		ApplyURI(uri).
		SetAuth(options.Credential{
			Username: cfg.Mongo.User,
			Password: cfg.Mongo.Pass,
		})

	if cfg.Mode == values.Dev {
		// Direct connection to the server, bypassing the driver's topology in a single server configuration
		opts.SetDirect(true)
	}

	// Create a new client and connect to the server
	client, err := mongo.Connect(context.TODO(), opts)
	if err != nil {
		logrus.Fatal(err)
	}

	// Check the connection
	err = client.Ping(context.TODO(), nil)
	if err != nil {
		logrus.Fatal(err)
	}

	logrus.Info("Connected to MongoDB!")

	Db = &Database{
		client: client,
		db:     client.Database(cfg.Mongo.Database),
	}

}
