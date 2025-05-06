package implementation

import (
	"esports_club_booking/internal/entities"
	"esports_club_booking/internal/repository"
	"log"
)

type PackageService struct {
	repository repository.PackageRepository
}

func NewPackageService(repository repository.PackageRepository) (*PackageService, error) {
	if repository == nil {
		log.Fatal("Nil repository provided to package service")
	}
	return &PackageService{repository}, nil
}

func (s *PackageService) GetAllPackages() ([]*entities.Package, error) {
	return nil, nil
}

func (s *PackageService) GetPackage(int64) (*entities.Package, error) {
	return nil, nil
}

func (s *PackageService) CreatePackage(*entities.Package) (int64, error) {
	return 0, nil
}

func (s *PackageService) UpdatePackage(*entities.Package) (int64, error) {
	return 0, nil
}

func (s *PackageService) DeletePackage(int64) error {
	return nil
}
