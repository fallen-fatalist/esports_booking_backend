package implementation

import (
	"booking_api/internal/entities"
	"booking_api/internal/repository"
	"log"
)

type ComputerService struct {
	repository repository.ComputerRepository
}

func NewComputerService(repository repository.ComputerRepository) (*ComputerService, error) {
	if repository == nil {
		log.Fatal("Nil repository provided to computer service")
	}
	return &ComputerService{repository: repository}, nil
}

func (s *ComputerService) GetAllComputers() ([]*entities.Computer, error) {
	return nil, nil
}

func (s *ComputerService) GetAllComputerStatuses() ([]*entities.ComputerStatus, error) {
	return nil, nil
}

func (s *ComputerService) GetAllPackages() ([]*entities.Package, error) {
	return nil, nil
}
