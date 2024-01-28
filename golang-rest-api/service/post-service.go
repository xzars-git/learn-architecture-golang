package service

import (
	"errors"
	"learn-architecture-golang/golang-rest-api/golang-rest-api/entity"
	"learn-architecture-golang/golang-rest-api/golang-rest-api/repository"
	"math/rand"
)

type PostService interface {
	Validate(post *entity.Post) error
	Create(post *entity.Post) (*entity.Post, error)
	FindAll() ([]entity.Post, error)
}

type service struct{}

var (
	repo repository.PostRepository
)

func NewPostService(repository repository.PostRepository) PostService {
	repo = repository
	return &service{}
}
func (*service) Validate(post *entity.Post) error {
	if post == nil {
		return errors.New("the post is empty")
	}
	if post.Title == "" {
		return errors.New("the title is empty")
	}
	return nil
}
func (*service) Create(post *entity.Post) (*entity.Post, error) {
	post.ID = rand.Int63()
	return repo.Save(post)
}
func (*service) FindAll() ([]entity.Post, error) {
	return repo.FindAll()
}
