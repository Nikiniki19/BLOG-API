package repository

import (
	"blog-api/internal/model"
	"fmt"
)

func (r *DBRepository) Create(post *model.BlogPost) (string, error) {
	err := r.DB.Create(post).Error
	if err != nil {
		return "", err
	}
	return "Blog post created successfully", nil
}

func (r *DBRepository) GetAllPosts() ([]*model.BlogPost, error) {
	var posts []*model.BlogPost
	result := r.DB.Find(&posts)
	if result.Error != nil {
		return nil, result.Error
	}
	return posts, nil
}

func (r *DBRepository) GetPostByID(id uint) (*model.BlogPost, error) {
	var post model.BlogPost
	result := r.DB.First(&post, id)
	if result.Error != nil {
		return nil, result.Error
	}
	return &post, nil
}


func (r *DBRepository) UpdatePostByID(post *model.BlogPost) error {
    result := r.DB.Model(&model.BlogPost{}).Where("id = ?", post.ID).Updates(post)
    if result.Error != nil {
        return result.Error
    }
    if result.RowsAffected == 0 {
        return fmt.Errorf("no post found with id %d", post.ID)
    }
    return nil
}

func (r *DBRepository) DeletePostByID(id uint) error {
	result := r.DB.Delete(&model.BlogPost{}, id)
	if result.RowsAffected == 0 {
		return fmt.Errorf("post not found")
	}
	return result.Error
}
