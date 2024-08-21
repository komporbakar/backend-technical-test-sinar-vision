package repository

import (
	"backend_technical_test/entities/models"

	"gorm.io/gorm"
)

type PostRepository interface {
	CountArticles() (int64, error)
	FindAll(limit int, offset int) ([]models.Posts, error)
	FindById(id int) (models.Posts, error)
	Save(post models.Posts) (models.Posts, error)
	Update(post models.Posts) (models.Posts, error)
	Delete(id int) (models.Posts, error)
}

type repository struct {
	db *gorm.DB
}

func NewPostRepository(db *gorm.DB) *repository {
	return &repository{db}
}

func (r *repository) CountArticles() (int64, error) {
	var count int64
	err := r.db.Model(&models.Posts{}).Count(&count).Error
	return count, err
}

func (r *repository) FindAll(limit int, offset int) ([]models.Posts, error) {
	var posts []models.Posts
	err := r.db.Debug().Limit(limit).Offset(offset).Find(&posts).Error
	return posts, err
}

func (r *repository) FindById(id int) (models.Posts, error) {
	var posts models.Posts
	err := r.db.First(&posts, id).Error
	return posts, err
}

func (r *repository) Save(post models.Posts) (models.Posts, error) {
	err := r.db.Debug().Create(&post).Error
	return post, err
}

func (r *repository) Update(post models.Posts) (models.Posts, error) {
	err := r.db.Debug().Save(&post).Error
	return post, err
}

func (r *repository) Delete(id int) (models.Posts, error) {
	err := r.db.Debug().Delete(&models.Posts{}, id).Error
	return models.Posts{}, err
}
