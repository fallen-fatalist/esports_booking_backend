package repository

import (
	"booking_api/internal/entities"
)

type Repository struct {
	UserRepository     UserRepository
	ComputerRepository ComputerRepository
	BookingRepository  BookingRepository
}

func NewRepository(
	userRepository UserRepository,
	computerRepository ComputerRepository,
	bookingRepository BookingRepository,

) (*Repository, error) {
	if userRepository == nil {
		panic("User repository is nil")
	} else if computerRepository == nil {
		panic("Computer repository is nil")
	} else if userRepository == nil {
		panic("User repository is nil")
	}
	return &Repository{
		UserRepository:     userRepository,
		ComputerRepository: computerRepository,
		BookingRepository:  bookingRepository,
	}, nil
}

type UserRepository interface {
	GetAll() ([]*entities.User, error)
	GetByID(int64) (*entities.User, error)
	Create(*entities.User) (int64, error)
	Update(*entities.Booking) (int64, error)
	Delete(int64) error
}

type ComputerRepository interface {
	GetAll() ([]*entities.Computer, error)
	GetByID(int64) (*entities.Computer, error)
	Create(*entities.Computer) (int64, error)
	Update(*entities.Computer) (int64, error)
	Delete(int64) (int64, error)
}

type BookingRepository interface {
	GetAll() ([]*entities.Booking, error)
	GetByID(int64) (*entities.Booking, error)
	Create(booking *entities.Booking) (int64, error)
	Update(booking *entities.Booking) (int64, error)
	Delete(int64) error
	/* Very bad and stupid solution must be optimized to priority queue */
	RefreshStatus()
	CreateRandomActiveBookingEndingSoon()
}

type PackageRepository interface {
	GetAll() ([]*entities.Package, error)
	GetByID(int64) (*entities.Package, error)
	Create(*entities.Package) (int64, error)
	Update(*entities.Package) (int64, error)
	Delete(int64) error
}
