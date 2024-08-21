package service

import (
	"backend_technical_test/entities/models"
	"backend_technical_test/entities/request"
	"backend_technical_test/repository"
)

type PostService interface {
	CreatePost(request request.PostRequest) (models.Posts, error)
	GetAllPostsByLimitAndOffset(limit int, offset int) ([]models.Posts, int64, error)
	GetPostById(id int) (models.Posts, error)
	UpdatePost(id int, request request.PostRequest) (models.Posts, error)
	DeletePost(id int) (models.Posts, error)
}

type service struct {
	repository repository.PostRepository
}

func NewPostService(repository repository.PostRepository) *service {
	return &service{repository}
}

func (s *service) CreatePost(request request.PostRequest) (models.Posts, error) {
	post := models.Posts{
		Title:    request.Title,
		Content:  request.Content,
		Category: request.Category,
		Status:   request.Status,
	}
	newPost, err := s.repository.Save(post)
	if err != nil {
		return newPost, err
	}
	return newPost, nil

}

func (s *service) GetAllPostsByLimitAndOffset(limit int, offset int) ([]models.Posts, int64, error) {
	posts, err := s.repository.FindAll(limit, offset)
	if err != nil {
		return nil, 0, err
	}
	totalCount, err := s.repository.CountArticles()
	if err != nil {
		return nil, 0, err
	}

	return posts, totalCount, nil
}

func (s *service) GetPostById(id int) (models.Posts, error) {
	post, err := s.repository.FindById(id)
	if err != nil {
		return post, err
	}
	return post, nil
}

func (s *service) UpdatePost(id int, request request.PostRequest) (models.Posts, error) {
	post, err := s.repository.FindById(id)
	if err != nil {
		return post, err
	}
	post.Title = request.Title
	post.Content = request.Content
	post.Category = request.Category
	post.Status = request.Status
	updatedPost, err := s.repository.Update(post)
	if err != nil {
		return updatedPost, err
	}
	return updatedPost, nil
}

func (s *service) DeletePost(id int) (models.Posts, error) {
	post, err := s.repository.FindById(id)
	if err != nil {
		return post, err
	}
	deletedPost, err := s.repository.Delete(id)
	if err != nil {
		return deletedPost, err
	}
	return deletedPost, nil
}
