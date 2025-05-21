// @title Blog Post API
// @version 1.0
// @description This is a sample blog post CRUD API built with Go Fiber
// @host localhost:3000
// @BasePath /api/blog-post

package main

import (
	"blog-api/internal/database"
	"blog-api/internal/handlers"
	"blog-api/internal/repository"
	"blog-api/internal/server"
	"blog-api/internal/services"

	"github.com/gofiber/fiber/v2"
	"github.com/rs/zerolog/log"
)

func main() {
	log.Info().Msg("Connecting to the database...")
	dbConn, err := database.ConnectToDatabase()
	if err != nil {
		log.Fatal().Err(err).Msg("Database connection failed")
		return
	}

	repo, err := repository.NewRepository(dbConn)
	if err != nil {
		log.Fatal().Err(err).Msg("Failed to create repository")
		return
	}

	service, err := services.NewService(repo)
	if err != nil {
		log.Fatal().Err(err).Msg("Failed to create service")
		return
	}

	handler, err := handlers.NewHandler(service)
	if err != nil {
		log.Fatal().Err(err).Msg("Failed to create handler")
		return
	}

	app := fiber.New()
	log.Info().Msg("Starting the application...")
	server.StartApplicationServer(app, handler)


	if err := app.Listen(":3000"); err != nil {
		log.Fatal().Err(err).Msg("Failed to start server")
	}
}
