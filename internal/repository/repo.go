package repository

import (
	"blog-api/internal/model"
	"errors"

	"gorm.io/gorm"
)

type DBRepository struct {
	DB *gorm.DB
}


type RepositoryInterface interface {
	Create(post *model.BlogPost) (string, error)
	GetAllPosts() ([]*model.BlogPost, error)
	GetPostByID(id uint) (*model.BlogPost, error)
	UpdatePostByID(post *model.BlogPost) error
	DeletePostByID(id uint) error
}

func NewRepository(db *gorm.DB) (RepositoryInterface, error) {
	if db == nil {
		return nil, errors.New("db cannot be nil")
	}
	return &DBRepository{
		DB: db,
	}, nil
}
