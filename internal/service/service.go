package service

import (
	"booking_api/internal/entities"
	"log"
)

type Service struct {
	ComputerService ComputerService
	UserService     UserService
	BookingService  BookingService
	PackageService  PackageService
}

func NewService(
	ComputerService ComputerService,
	UserService UserService,
	BookingService BookingService,
	PackageService PackageService,
) (*Service, error) {
	if ComputerService == nil {
		log.Fatal("Nil computer service provided to general service")
	} else if UserService == nil {
		log.Fatal("Nil user service provided to general service")
	} else if BookingService == nil {
		log.Fatal("Nil booking service provided to general service")
	} else if PackageService == nil {
		log.Fatal("Nil package service provided to general service")
	}
	return &Service{
		ComputerService: ComputerService,
		UserService:     UserService,
		BookingService:  BookingService,
		PackageService:  PackageService,
	}, nil
}

type ComputerService interface {
	GetAllComputers() ([]*entities.Computer, error)
	GetAllComputerStatuses() ([]*entities.ComputerStatus, error)
	GetAllPackages() ([]*entities.Package, error)
}

type UserService interface {
	GetAllUsers() ([]*entities.User, error)
}

type BookingService interface {
	GetAllBookings() ([]*entities.Booking, error)
	GetPendingBookings() ([]*entities.Booking, error)
	GetFinishedBookings() ([]*entities.Booking, error)
}

type PackageService interface {
	GetPackages() ([]*entities.Package, error)
	GetPackage(int64) (*entities.Package, error)
	CreatePackage(*entities.Package) (int64, error)
	UpdatePackage(*entities.Package) (int64, error)
	DeletePackage(int64) error
}
