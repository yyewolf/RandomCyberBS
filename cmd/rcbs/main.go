package main

import (
	"log/slog"
	"rcbs/api"
	"rcbs/internal/env"
	"rcbs/internal/mongo"
	"rcbs/internal/values"
	"rcbs/models"

	"github.com/rs/cors"

	chiLog "github.com/chi-middleware/logrus-logger"
	"github.com/go-fuego/fuego"
	sloglogrus "github.com/samber/slog-logrus/v2"
	"github.com/sirupsen/logrus"
)

func main() {
	// Load environment variable
	env.Load()

	// Get config
	cfg := env.Get()

	// Connect to MongoDB
	mongo.Connect()

	// Load models
	models.LoadDatabase()

	logrusLogger := logrus.New()

	// Set log level for both logrus and slog
	logrus.SetLevel(logrus.DebugLevel)
	logrusLogger.SetLevel(logrus.DebugLevel)
	slogLevel := slog.LevelDebug
	if cfg.Mode == values.Prod {
		logrus.SetLevel(logrus.InfoLevel)
		logrusLogger.SetLevel(logrus.InfoLevel)
		slogLevel = slog.LevelInfo
	}

	logger := slog.New(sloglogrus.Option{Level: slogLevel, Logger: logrusLogger}.NewLogrusHandler())

	// Create our fuego server
	s := fuego.NewServer(
		fuego.WithPort(cfg.Server.Port),
		fuego.WithLogHandler(logger.Handler()),
		fuego.WithOpenAPIConfig(fuego.OpenAPIConfig{
			DisableLocalSave: true,
			SwaggerUrl:       "/api",
			JsonUrl:          "/api/api.json",
		}),
		fuego.WithCorsMiddleware(cors.New(cors.Options{
			AllowedOrigins:   []string{"*"},
			AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
			AllowedHeaders:   []string{"*"},
			AllowCredentials: true,
		}).Handler),
	)

	s.Security = fuego.NewSecurity()

	// Setup middlewares
	fuego.Use(s, chiLog.Logger("router", logrusLogger))

	// Setup our API
	api.Setup(s)

	// Run the server
	s.Run()
}
