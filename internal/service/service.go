package service

import (
	"booking_api/internal/entities"
	"errors"
	"log"
)

type Service struct {
	ComputerService ComputerService
	UserService     UserService
	BookingService  BookingService
	PackageService  PackageService
}

type ComputerService interface {
	GetAllComputers() ([]*entities.Computer, error)
	GetComputer(id int64) (*entities.Computer, error)
	CreateComputer(*entities.Computer) (int64, error)
	DeleteComputer(id int64) error
	GetAllComputerStatuses() ([]*entities.ComputerStatus, error)
	GetComputerStatus(id int64) (*entities.ComputerStatus, error)
}

type UserService interface {
	GetAllUsers() ([]*entities.User, error)
	GetUser(int64) (*entities.User, error)
	CreateUser(*entities.User) (int64, error)
	//AuthUser()
	//RegisterUser()
	//FillBalance()

}

type BookingService interface {
	GetAllBookings() ([]*entities.Booking, error)
	CreateBooking(*entities.Booking) (int64, error)
	GetComputerBookings(id int64) ([]*entities.Booking, error)
	GetActiveBookings() ([]*entities.Booking, error)
	GetPendingBookings() ([]*entities.Booking, error)
	GetFinishedBookings() ([]*entities.Booking, error)
	GetCancelledBookings() ([]*entities.Booking, error)
	GetComputersLeftOccupiedTime() ([]*entities.ComputerOccupiedLeftTime, error)
	GetComputerLeftOccupiedTime(id int64) (*entities.ComputerOccupiedLeftTime, error)
	RefreshBookings()
	CreateRandomActiveBookingEndingSoon()
}

type PackageService interface {
	GetAllPackages() ([]*entities.Package, error)
	GetPackage(id int64) (*entities.Package, error)
	CreatePackage(*entities.Package) (int64, error)
	UpdatePackage(*entities.Package) (int64, error)
	DeletePackage(int64) error
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

var (
	ErrUnhandledError = errors.New("unhandled error occured, please ask an admin to fix")
)
