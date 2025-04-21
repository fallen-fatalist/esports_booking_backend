package implementation

import (
	"booking_api/internal/entities"
	"booking_api/internal/repository"
	"log"
)

type UserService struct {
	repository repository.UserRepository
}

func NewUserService(repository repository.UserRepository) (*UserService, error) {
	if repository == nil {
		log.Fatal("Nil repository provided to user service")
	}
	return &UserService{repository}, nil
}

func (s *UserService) GetAllUsers() ([]*entities.User, error) {
	return nil, nil
}
