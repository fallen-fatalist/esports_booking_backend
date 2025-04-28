package implementation

import (
	"booking_api/internal/entities"
	"booking_api/internal/repository"
	"booking_api/internal/service"
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
	return s.repository.GetAll()
}

func (s *BookingService) GetComputerBookings(id int64) ([]*entities.Booking, error) {
	if err := service.ValidateID(id); err != nil {
		return nil, err
	}
	return s.repository.GetByID(id)
}

func (s *BookingService) GetPendingBookings() ([]*entities.Booking, error) {
	return nil, nil
}

func (s *BookingService) GetFinishedBookings() ([]*entities.Booking, error) {
	return nil, nil
}

func (s *BookingService) GetComputersLeftOccupiedTime() ([]*service.ComputerOccupiedLeftTime, error) {
	return nil, nil
}

func (s *BookingService) GetComputerLeftOccupiedTime(id int64) (*service.ComputerOccupiedLeftTime, error) {
	return nil, nil
}
