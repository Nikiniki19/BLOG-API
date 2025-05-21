package services

import (
	"blog-api/internal/model"
	"time"
)


func (s *Service) CreatePost(post *model.BlogPost) (string, error) {
	post.CreatedAt = time.Now()
	post.UpdatedAt = time.Now()

	msg, err := s.Repo.Create(post)
	if err != nil {
		return "", err
	}
	return msg, nil
}

func (s *Service) GetAllPosts() ([]*model.BlogPost, error) {
    return s.Repo.GetAllPosts()
}

func (s *Service) GetPostByID(id uint) (*model.BlogPost, error) {
    return s.Repo.GetPostByID(id)
}

func (s *Service) UpdatePostByID(post *model.BlogPost) error {
	return s.Repo.UpdatePostByID(post)
}

func (s *Service) DeletePostByID(id uint) error {
	return s.Repo.DeletePostByID(id)
}