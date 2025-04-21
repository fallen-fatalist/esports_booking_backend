package implementation

import (
	"booking_api/internal/entities"
	"booking_api/internal/repository"
	"log"
)

type BookingService struct {
	repository repository.BookingRepository
}

func NewBookingService(repository repository.BookingRepository) (*BookingService, error) {
	if repository == nil {
		log.Fatal("Nil repository provided to booking service")
	}
	return &BookingService{repository: repository}, nil
}

func (s *BookingService) GetAllBookings() ([]*entities.Booking, error) {
	return nil, nil
}

func (s *BookingService) GetPendingBookings() ([]*entities.Booking, error) {
	return nil, nil
}

func (s *BookingService) GetFinishedBookings() ([]*entities.Booking, error) {
	return nil, nil
}
