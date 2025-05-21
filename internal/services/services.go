package services

import (
	"blog-api/internal/model"
	"blog-api/internal/repository"
)

type Service struct {
	Repo repository.RepositoryInterface
}

type ServiceInterface interface {
	CreatePost(post *model.BlogPost) (string, error)
	GetAllPosts() ([]*model.BlogPost, error)
	GetPostByID(id uint) (*model.BlogPost, error)
	UpdatePostByID(post *model.BlogPost) error
	DeletePostByID(id uint) error
}

func NewService(repo repository.RepositoryInterface) (ServiceInterface, error) {
	return &Service{
		Repo: repo,
	}, nil
}
