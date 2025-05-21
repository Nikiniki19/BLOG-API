package handlers

import (
	"blog-api/internal/model"
	"strconv"
	"strings"

	"github.com/gofiber/fiber/v2"
)

// CreatePost godoc
// @Summary     Create a new blog post
// @Description Add a new post to the blog
// @Tags        Blog
// @Accept      json
// @Produce     json
// @Param       post body model.BlogPost true "Blog post data"
// @Success 200 {object} model.BlogPostResponse
// @Failure     400 {object} model.ErrorResponse
// @Router      /blog-post/createPost [post]
func (h *Handler) CreatePost(c *fiber.Ctx) error {
	post := new(model.BlogPost)
	if err := c.BodyParser(post); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "cannot parse JSON",
		})
	}

	msg, err := h.Ser.CreatePost(post)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return c.Status(fiber.StatusCreated).JSON(fiber.Map{
		"message": msg,
		"id":      post.ID,
	})
}

// GetAllPosts godoc
// @Summary     Get all blog posts
// @Description Retrieve all blog posts
// @Tags        Blog
// @Produce     json
// @Success 200 {object} model.BlogPostResponse
// @Router      /blog-post/getAllPosts [get]
func (h *Handler) GetAllPosts(c *fiber.Ctx) error {
	posts, err := h.Ser.GetAllPosts()
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return c.Status(fiber.StatusOK).JSON(posts)
}

// GetPostByID godoc
// @Summary      Get a blog post by ID
// @Description  Retrieve a single blog post by its ID
// @Tags         Blog
// @Produce      json
// @Param        id path int true "Post ID"
// @Success 200 {object} model.BlogPostResponse
// @Failure      400 {object} model.ErrorResponse
// @Failure      404 {object} model.ErrorResponse
// @Router       /blog-post/{id} [get]
func (h *Handler) GetPostByID(c *fiber.Ctx) error {
	idParam := strings.TrimSpace(c.Params("id"))
	id, err := strconv.ParseUint(idParam, 10, 64)
	if err != nil || id == 0 {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid post ID",
		})
	}
	post, err := h.Ser.GetPostByID(uint(id))
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error": "No post found with the given ID",
		})
	}

	return c.JSON(post)
}

// UpdatePostByID godoc
// @Summary      Update a blog post by ID
// @Description  Update a specific blog post's details
// @Tags         Blog
// @Accept       json
// @Produce      json
// @Param        id path int true "Post ID"
// @Param        post body model.BlogPost true "Updated blog post"
// @Success 200 {object} model.BlogPostResponse
// @Failure      400 {object} model.ErrorResponse
// @Failure      404 {object} model.ErrorResponse
// @Router       /blog-post/{id} [patch]
func (h *Handler) UpdatePostByID(c *fiber.Ctx) error {
	idParam := c.Params("id")
	id, err := strconv.ParseUint(idParam, 10, 64)
	if err != nil || id == 0 {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid post ID",
		})
	}

	post := new(model.BlogPost)
	if err := c.BodyParser(post); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Failed to parse request body",
		})
	}
	post.ID = uint(id)

	err = h.Ser.UpdatePostByID(post)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to update post",
		})
	}

	return c.JSON(fiber.Map{
		"message": "Post updated successfully",
		"id":      post.ID,
	})
}

// DeletePostByID godoc
// @Summary      Delete a blog post by ID
// @Description  Remove a blog post using its ID
// @Tags         Blog
// @Produce      json
// @Param        id path int true "Post ID"
// @Success 200 {object} model.BlogPostResponse
// @Failure      400 {object} model.ErrorResponse
// @Failure      404 {object} model.ErrorResponse
// @Router       /blog-post/{id} [delete]
func (h *Handler) DeletePostByID(c *fiber.Ctx) error {
	idParam := strings.TrimSpace(c.Params("id"))
	id, err := strconv.ParseUint(idParam, 10, 64)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid post ID",
		})
	}

	err = h.Ser.DeletePostByID(uint(id))
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error": "Post not found or could not be deleted",
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "Post deleted successfully",
		"id":      id,
	})
}
