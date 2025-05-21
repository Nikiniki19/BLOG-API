package services

import (
	"blog-api/internal/mocks"
	"blog-api/internal/model"
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestCreatePost(t *testing.T) {
	mockRepo := new(mocks.RepositoryInterface)
	service := &Service{Repo: mockRepo}

	input := &model.BlogPost{
		Title:       "Test Title",
		Description: "Test Desc",
		Body:        "Test Body",
	}

	expectedMsg := "Post created Successfully"
	mockRepo.On("Create", mock.MatchedBy(func(post *model.BlogPost) bool {
		return !post.CreatedAt.IsZero() && !post.UpdatedAt.IsZero() &&
			post.Title == input.Title &&
			post.Description == input.Description &&
			post.Body == input.Body
	})).Return(expectedMsg, nil)

	result, err := service.CreatePost(input)

	assert.NoError(t, err)
	assert.Equal(t, expectedMsg, result)

	mockRepo.AssertExpectations(t)
}

func TestCreatePost_Error(t *testing.T) {
	mockRepo := new(mocks.RepositoryInterface)
	service := &Service{Repo: mockRepo}

	input := &model.BlogPost{
		Title:       "Test Title",
		Description: "Test Desc",
		Body:        "Test Body",
	}

	expectedErr := errors.New("database error")
	mockRepo.On("Create", mock.AnythingOfType("*model.BlogPost")).Return("", expectedErr)

	msg, err := service.CreatePost(input)

	assert.Error(t, err)
	assert.Equal(t, "", msg)
	assert.Equal(t, expectedErr, err)

	mockRepo.AssertExpectations(t)
}

func TestGetAllPosts(t *testing.T) {
	mockRepo := new(mocks.RepositoryInterface)
	service := &Service{Repo: mockRepo}

	expectedPosts := []*model.BlogPost{
		{ID: 1, Title: "Post 1"},
		{ID: 2, Title: "Post 2"},
	}

	mockRepo.On("GetAllPosts").Return(expectedPosts, nil)

	posts, err := service.GetAllPosts()

	assert.NoError(t, err)
	assert.Equal(t, expectedPosts, posts)
	mockRepo.AssertExpectations(t)
}

func TestGetPostByID(t *testing.T) {
	mockRepo := new(mocks.RepositoryInterface)
	service := &Service{Repo: mockRepo}

	expectedPost := &model.BlogPost{ID: 1, Title: "Post 1"}

	mockRepo.On("GetPostByID", uint(1)).Return(expectedPost, nil)

	post, err := service.GetPostByID(1)

	assert.NoError(t, err)
	assert.Equal(t, expectedPost, post)
	mockRepo.AssertExpectations(t)
}

func TestUpdatePostByID(t *testing.T) {
	mockRepo := new(mocks.RepositoryInterface)
	service := &Service{Repo: mockRepo}

	postToUpdate := &model.BlogPost{ID: 1, Title: "Updated Title"}

	mockRepo.On("UpdatePostByID", postToUpdate).Return(nil)

	err := service.UpdatePostByID(postToUpdate)

	assert.NoError(t, err)
	mockRepo.AssertExpectations(t)
}

func TestDeletePostByID(t *testing.T) {
	mockRepo := new(mocks.RepositoryInterface)
	service := &Service{Repo: mockRepo}

	mockRepo.On("DeletePostByID", uint(1)).Return(nil)

	err := service.DeletePostByID(1)

	assert.NoError(t, err)
	mockRepo.AssertExpectations(t)
}
