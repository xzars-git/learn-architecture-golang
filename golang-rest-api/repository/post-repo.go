package repository

import (
	"learn-architecture-golang/golang-rest-api/golang-rest-api/entity"
)

type PostRepository interface {
	Save(post *entity.Post) (*entity.Post, error)
	FindAll() ([]entity.Post, error)
}
