package handlers

import (
	"blog-api/internal/services"

	"github.com/gofiber/fiber/v2"
)

type Handler struct {
	Ser services.ServiceInterface
}

type HandlerInterface interface {
	CreatePost(c *fiber.Ctx) error
	GetAllPosts(c *fiber.Ctx) error
	GetPostByID(c *fiber.Ctx) error
	UpdatePostByID(c *fiber.Ctx) error
	DeletePostByID(c *fiber.Ctx) error
}

func NewHandler(ser services.ServiceInterface) (HandlerInterface, error) {
	return &Handler{
		Ser: ser,
	}, nil
}
