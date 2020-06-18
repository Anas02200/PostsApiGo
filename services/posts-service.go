package services

import (
	"GoApi/entity"
	"GoApi/repository"
	"errors"
	"math/rand"
)

var (
	//posts []entity.Post
	repo repository.PostRepository = repository.NewFirestoreRepository()
)

type PostService interface {
	Validate(post *entity.Post) error
	Create(post *entity.Post) (*entity.Post, error)
	FindAll() ([]entity.Post, error)
}
type service struct {
}

func (*service) Validate(post *entity.Post) error {
	if post == nil {
		err := errors.New("empty post")
		return err
	}
	if post.Title == "" {
		err := errors.New("empty title")
		return err
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

func NewPostService() PostService {
	return &service{}
}
