package services

import "example.com/todogo/models"

type PostService interface {
	CreatePost(post models.Post) (models.Post, error)
	GetPosts() ([]models.Post, error)
	GetPost(id string) (models.Post, error)
	UpdatePost(id string, post models.Post) (models.Post, error)
	DeletePost(id string) error
}
