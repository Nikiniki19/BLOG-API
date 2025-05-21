package server

import (
	"blog-api/internal/handlers"
	_ "blog-api/docs"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/recover"
	swagger "github.com/gofiber/swagger"  
)

func StartApplicationServer(app *fiber.App, h handlers.HandlerInterface) {
	app.Use(logger.New())
	app.Use(recover.New())

	api := app.Group("/api")
	blog := api.Group("/blog-post")
	blog.Post("/createPost", h.CreatePost)
	blog.Get("/getAllPosts", h.GetAllPosts)
	blog.Get("/:id", h.GetPostByID)
	blog.Patch("/:id", h.UpdatePostByID)
	blog.Delete("/:id", h.DeletePostByID)

	app.Get("/swagger/*", swagger.HandlerDefault)
}
