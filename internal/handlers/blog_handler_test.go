package handlers

import (
	"blog-api/internal/mocks"
	"blog-api/internal/model"
	"bytes"
	"encoding/json"
	"errors"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/gofiber/fiber/v2"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func setupTestHandler() (*fiber.App, *Handler, *mocks.ServiceInterface) {
	app := fiber.New()
	mockService := new(mocks.ServiceInterface)
	handler := &Handler{Ser: mockService}
	return app, handler, mockService
}

func TestCreatePost_Success(t *testing.T) {
	app, handler, mockService := setupTestHandler()
	app.Post("/blog-post/createPost", handler.CreatePost)

	blog := &model.BlogPost{
		Title:       "Test Title",
		Description: "Test Description",
		Body:        "Test Body",
	}

	mockService.On("CreatePost", mock.AnythingOfType("*model.BlogPost")).
		Return("Post created Successfully", nil)

	body, _ := json.Marshal(blog)
	req := httptest.NewRequest(http.MethodPost, "/blog-post/createPost", bytes.NewReader(body))
	req.Header.Set("Content-Type", "application/json")

	resp, err := app.Test(req)
	assert.NoError(t, err)
	assert.Equal(t, fiber.StatusCreated, resp.StatusCode)
}

func TestCreatePost_ServiceError(t *testing.T) {
	app, handler, mockService := setupTestHandler()
	app.Post("/blog-post/createPost", handler.CreatePost)

	post := &model.BlogPost{
		Title:       "Test Title",
		Description: "Test Description",
		Body:        "Test Body",
	}
	body, _ := json.Marshal(post)

	mockService.On("CreatePost", mock.AnythingOfType("*model.BlogPost")).Return("", errors.New("DB error"))
	req := httptest.NewRequest(http.MethodPost, "/blog-post/createPost", bytes.NewReader(body))
	req.Header.Set("Content-Type", "application/json")
	resp, err := app.Test(req)

	assert.NoError(t, err)
	assert.Equal(t, fiber.StatusInternalServerError, resp.StatusCode)

	mockService.AssertExpectations(t)
}

func TestCreatePost_InvalidJSON(t *testing.T) {
	app, handler, _ := setupTestHandler()
	app.Post("/blog-post/createPost", handler.CreatePost)

	invalidBody := `{"title": "Bad JSON", "description": "Missing quote}`

	req := httptest.NewRequest(http.MethodPost, "/blog-post/createPost", strings.NewReader(invalidBody))
	req.Header.Set("Content-Type", "application/json")
	resp, err := app.Test(req)

	assert.NoError(t, err)
	assert.Equal(t, fiber.StatusBadRequest, resp.StatusCode)
}

func TestGetAllPosts_Success(t *testing.T) {
	app, handler, mockService := setupTestHandler()
	app.Get("/blog-post/getAllPosts", handler.GetAllPosts)

	expectedPosts := []*model.BlogPost{
		{ID: 1, Title: "Post 1", Description: "Desc", Body: "Body"},
		{ID: 2, Title: "Post 2", Description: "Desc", Body: "Body"},
	}

	mockService.On("GetAllPosts").Return(expectedPosts, nil)

	req := httptest.NewRequest(http.MethodGet, "/blog-post/getAllPosts", nil)
	resp, err := app.Test(req)

	assert.NoError(t, err)
	assert.Equal(t, fiber.StatusOK, resp.StatusCode)
}

func TestGetAllPosts_Error(t *testing.T) {
	app, handler, mockService := setupTestHandler()
	app.Get("/blog-post/getAllPosts", handler.GetAllPosts)

	mockService.On("GetAllPosts").Return(nil, errors.New("db error"))

	req := httptest.NewRequest(http.MethodGet, "/blog-post/getAllPosts", nil)
	resp, err := app.Test(req)

	assert.NoError(t, err)
	assert.Equal(t, fiber.StatusInternalServerError, resp.StatusCode)
}

func TestGetPostByID_Success(t *testing.T) {
	app, handler, mockService := setupTestHandler()
	app.Get("/blog-post/:id", handler.GetPostByID)

	expectedPost := &model.BlogPost{ID: 1, Title: "Post", Description: "Desc", Body: "Body"}

	mockService.On("GetPostByID", uint(1)).Return(expectedPost, nil)

	req := httptest.NewRequest(http.MethodGet, "/blog-post/1", nil)
	resp, err := app.Test(req)

	assert.NoError(t, err)
	assert.Equal(t, fiber.StatusOK, resp.StatusCode)
}

func TestGetPostByID_InvalidID(t *testing.T) {
	app, handler, _ := setupTestHandler()
	app.Get("/blog-post/:id", handler.GetPostByID)

	req := httptest.NewRequest(http.MethodGet, "/blog-post/abc", nil)
	resp, err := app.Test(req)

	assert.NoError(t, err)
	assert.Equal(t, fiber.StatusBadRequest, resp.StatusCode)
}

func TestGetPostByID_NotFound(t *testing.T) {
	app, handler, mockService := setupTestHandler()
	app.Get("/blog-post/:id", handler.GetPostByID)

	mockService.On("GetPostByID", uint(99)).Return(nil, errors.New("not found"))

	req := httptest.NewRequest(http.MethodGet, "/blog-post/99", nil)
	resp, err := app.Test(req)

	assert.NoError(t, err)
	assert.Equal(t, fiber.StatusNotFound, resp.StatusCode)
}

func TestUpdatePostByID_Success(t *testing.T) {
	app, handler, mockService := setupTestHandler()
	app.Patch("/blog-post/:id", handler.UpdatePostByID)

	post := &model.BlogPost{Title: "Updated", Description: "Desc", Body: "Body"}
	body, _ := json.Marshal(post)

	mockService.On("UpdatePostByID", mock.AnythingOfType("*model.BlogPost")).Return(nil)

	req := httptest.NewRequest(http.MethodPatch, "/blog-post/1", bytes.NewReader(body))
	req.Header.Set("Content-Type", "application/json")

	resp, err := app.Test(req)
	assert.NoError(t, err)
	assert.Equal(t, fiber.StatusOK, resp.StatusCode)
}


func TestUpdatePostByID_Error(t *testing.T) {
	app, handler, mockService := setupTestHandler()
	app.Patch("/blog-post/:id", handler.UpdatePostByID)

	post := &model.BlogPost{Title: "Updated", Description: "Desc", Body: "Body"}
	body, _ := json.Marshal(post)

	mockService.On("UpdatePostByID", mock.AnythingOfType("*model.BlogPost")).Return(errors.New("update error"))

	req := httptest.NewRequest(http.MethodPatch, "/blog-post/1", bytes.NewReader(body))
	req.Header.Set("Content-Type", "application/json")

	resp, err := app.Test(req)
	assert.NoError(t, err)
	assert.Equal(t, fiber.StatusInternalServerError, resp.StatusCode)
}

func TestUpdatePostByID_InvalidBody(t *testing.T) {
	app, handler, _ := setupTestHandler()
	app.Patch("/blog-post/:id", handler.UpdatePostByID)

	invalidJSON := `{"title": "Bad JSON"`

	req := httptest.NewRequest(http.MethodPatch, "/blog-post/1", strings.NewReader(invalidJSON))
	req.Header.Set("Content-Type", "application/json")

	resp, err := app.Test(req)
	assert.NoError(t, err)
	assert.Equal(t, fiber.StatusBadRequest, resp.StatusCode)

	var result map[string]string
	json.NewDecoder(resp.Body).Decode(&result)
	assert.Equal(t, "Failed to parse request body", result["error"])
}

func TestUpdatePostByID_InvalidID(t *testing.T) {
	app, handler, _ := setupTestHandler()
	app.Patch("/blog-post/:id", handler.UpdatePostByID)

	post := &model.BlogPost{
		Title:       "Updated Title",
		Description: "Updated Description",
		Body:        "Updated Body",
	}
	body, _ := json.Marshal(post)

	req := httptest.NewRequest(http.MethodPatch, "/blog-post/invalid-id", bytes.NewReader(body))
	req.Header.Set("Content-Type", "application/json")

	resp, err := app.Test(req)

	assert.NoError(t, err)
	assert.Equal(t, fiber.StatusBadRequest, resp.StatusCode)

	var result map[string]string
	json.NewDecoder(resp.Body).Decode(&result)
	assert.Equal(t, "Invalid post ID", result["error"])
}

func TestDeletePostByID_Success(t *testing.T) {
	app, handler, mockService := setupTestHandler()
	app.Delete("/blog-post/:id", handler.DeletePostByID)

	mockService.On("DeletePostByID", uint(1)).Return(nil)

	req := httptest.NewRequest(http.MethodDelete, "/blog-post/1", nil)
	resp, err := app.Test(req)

	assert.NoError(t, err)
	assert.Equal(t, fiber.StatusOK, resp.StatusCode)
}

func TestDeletePostByID_InvalidID(t *testing.T) {
	app, handler, _ := setupTestHandler()
	app.Delete("/blog-post/:id", handler.DeletePostByID)

	req := httptest.NewRequest(http.MethodDelete, "/blog-post/abc", nil)
	resp, err := app.Test(req)

	assert.NoError(t, err)
	assert.Equal(t, fiber.StatusBadRequest, resp.StatusCode)
}

func TestDeletePostByID_NotFound(t *testing.T) {
	app, handler, mockService := setupTestHandler()
	app.Delete("/blog-post/:id", handler.DeletePostByID)

	mockService.On("DeletePostByID", uint(99)).Return(errors.New("not found"))

	req := httptest.NewRequest(http.MethodDelete, "/blog-post/99", nil)
	resp, err := app.Test(req)

	assert.NoError(t, err)
	assert.Equal(t, fiber.StatusNotFound, resp.StatusCode)
}
