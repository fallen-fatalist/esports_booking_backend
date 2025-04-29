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

func (s *UserService) GetUser(id int64) (*entities.User, error) {
	if err := entities.ValidateID(id); err != nil {
		return nil, err
	}

	return nil, nil
}

func (s *UserService) CreateUser(user *entities.User) (int64, error) {
	return 0, nil
}
