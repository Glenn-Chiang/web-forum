package repos

import (
	"cvwo-backend/models"
	"fmt"

	"gorm.io/gorm"
)

type PostRepo struct {
	DB *gorm.DB
}

func NewPostRepo(db *gorm.DB) *PostRepo {
	return &PostRepo{DB: db}
}

func (repo *PostRepo) GetAll() ([]models.Post, error) {
	var posts []models.Post
	err := repo.DB.Find(&posts).Error
	if err != nil {
		return nil, err
	}
	return posts, err
}

func (repo *PostRepo) GetByID(id string) (*models.Post, error) {
	var post models.Post
	err := repo.DB.First(&post, id).Error
	if err != nil {
		return nil, err
	}
	return &post, nil
}

func (repo *PostRepo) Create(post *models.Post) (*models.Post, error) {
	if err := repo.DB.Create(post).Error; err != nil {
		return nil, err
	}
	return post, nil
}

func (repo *PostRepo) Delete(id string) error {
	result := repo.DB.Delete(&models.User{}, id)
	if result.Error != nil {
		return result.Error
	}
	if result.RowsAffected == 0 {
		return fmt.Errorf("post not found")
	}
	return nil
}
