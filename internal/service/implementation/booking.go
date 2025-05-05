package implementation

import (
	"booking_api/internal/entities"
	"booking_api/internal/repository"
	"booking_api/internal/service"
	"fmt"
	"log"
	"log/slog"
	"time"
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
	if bookings, err := s.repository.GetAll(); err != nil {
		return nil, err
	} else {
		return bookings, nil
	}

}

func (s *BookingService) CreateBooking(booking *entities.Booking) (int64, error) {
	if err := booking.Validate(); err != nil {
		return 0, err
	}

	if id, err := s.repository.Create(booking); err != nil {
		slog.Error("Unhandled error creating booking:", err)
		return 0, service.ErrUnhandledError
	} else {
		return id, nil
	}
}

func (s *BookingService) GetComputerBookings(id int64) ([]*entities.Booking, error) {
	if err := entities.ValidateID(id); err != nil {
		return nil, err
	}

	bookings, err := s.GetAllBookings()
	if err != nil {
		slog.Error("Unhandled Error fetching pending bookings:", err)
		return nil, service.ErrUnhandledError
	}

	comp_bookings := make([]*entities.Booking, 0, len(bookings))

	for _, booking := range bookings {
		if booking.ComputerID == id {
			comp_bookings = append(comp_bookings, booking)
		}
	}

	return comp_bookings, nil
}

func (s *BookingService) GetActiveBookings() ([]*entities.Booking, error) {
	bookings, err := s.repository.GetAll()
	if err != nil {
		return nil, err
	}

	pending_bookings := make([]*entities.Booking, 0, len(bookings))

	for _, booking := range bookings {
		startTime, err := time.Parse(time.RFC3339, booking.StartTime)
		if err != nil {
			return nil, fmt.Errorf("invalid start_time: %w", err)
		}
		endTime, err := time.Parse(time.RFC3339, booking.EndTime)
		if err != nil {
			return nil, fmt.Errorf("invalid end_time: %w", err)
		}

		if booking.Status == "active" {
			pending_bookings = append(pending_bookings, &entities.Booking{
				ID:         booking.ID,
				UserID:     booking.UserID,
				ComputerID: booking.ComputerID,
				PackageID:  booking.PackageID,
				StartTime:  startTime.Format(time.RFC3339),
				EndTime:    endTime.Format(time.RFC3339),
				Status:     booking.Status,
				CreatedAt:  booking.CreatedAt,
			})
		}
	}
	return pending_bookings, nil
}

func (s *BookingService) GetPendingBookings() ([]*entities.Booking, error) {
	bookings, err := s.repository.GetAll()
	if err != nil {
		return nil, err
	}

	pending_bookings := make([]*entities.Booking, 0, len(bookings))

	for _, booking := range bookings {
		startTime, err := time.Parse(time.RFC3339, booking.StartTime)
		if err != nil {
			return nil, fmt.Errorf("invalid start_time: %w", err)
		}
		endTime, err := time.Parse(time.RFC3339, booking.EndTime)
		if err != nil {
			return nil, fmt.Errorf("invalid end_time: %w", err)
		}

		if booking.Status == "pending" {
			pending_bookings = append(pending_bookings, &entities.Booking{
				ID:         booking.ID,
				UserID:     booking.UserID,
				ComputerID: booking.ComputerID,
				PackageID:  booking.PackageID,
				StartTime:  startTime.Format(time.RFC3339),
				EndTime:    endTime.Format(time.RFC3339),
				Status:     booking.Status,
				CreatedAt:  booking.CreatedAt,
			})
		}
	}
	return pending_bookings, nil
}

func (s *BookingService) GetFinishedBookings() ([]*entities.Booking, error) {
	bookings, err := s.repository.GetAll()
	if err != nil {
		return nil, err
	}

	finished_bookings := make([]*entities.Booking, 0, len(bookings))

	for _, booking := range bookings {
		startTime, err := time.Parse(time.RFC3339, booking.StartTime)
		if err != nil {
			return nil, fmt.Errorf("invalid start_time: %w", err)
		}
		endTime, err := time.Parse(time.RFC3339, booking.EndTime)
		if err != nil {
			return nil, fmt.Errorf("invalid end_time: %w", err)
		}

		if booking.Status == "finished" {
			finished_bookings = append(finished_bookings, &entities.Booking{
				ID:         booking.ID,
				UserID:     booking.UserID,
				ComputerID: booking.ComputerID,
				PackageID:  booking.PackageID,
				StartTime:  startTime.Format(time.RFC3339),
				EndTime:    endTime.Format(time.RFC3339),
				Status:     booking.Status,
				CreatedAt:  booking.CreatedAt,
			})
		}
	}
	return finished_bookings, nil
}

func (s *BookingService) GetCancelledBookings() ([]*entities.Booking, error) {
	bookings, err := s.repository.GetAll()
	if err != nil {
		return nil, err
	}

	cancelled_bookings := make([]*entities.Booking, 0, len(bookings))

	for _, booking := range bookings {
		startTime, err := time.Parse(time.RFC3339, booking.StartTime)
		if err != nil {
			return nil, fmt.Errorf("invalid start_time: %w", err)
		}
		endTime, err := time.Parse(time.RFC3339, booking.EndTime)
		if err != nil {
			return nil, fmt.Errorf("invalid end_time: %w", err)
		}

		if booking.Status == "cancelled" {
			cancelled_bookings = append(cancelled_bookings, &entities.Booking{
				ID:         booking.ID,
				UserID:     booking.UserID,
				ComputerID: booking.ComputerID,
				PackageID:  booking.PackageID,
				StartTime:  startTime.Format(time.RFC3339),
				EndTime:    endTime.Format(time.RFC3339),
				Status:     booking.Status,
				CreatedAt:  booking.CreatedAt,
			})
		}
	}
	return cancelled_bookings, nil

}

// TODO: Test
func (s *BookingService) GetComputersLeftOccupiedTime() ([]*entities.ComputerOccupiedLeftTime, error) {
	bookings, err := s.repository.GetAll()
	if err != nil {
		return nil, err
	}

	now := time.Now().UTC()

	leftTimes := make([]*entities.ComputerOccupiedLeftTime, 0, len(bookings))

	for _, booking := range bookings {
		if booking.Status == "active" {
			endTime, err := time.Parse(time.RFC3339, booking.EndTime)
			if err != nil {
				return nil, fmt.Errorf("invalid end_time format for booking %d: %w", booking.ID, err)
			}

			if endTime.After(now) {
				leftDuration := endTime.Sub(now)
				leftTimes = append(leftTimes, &entities.ComputerOccupiedLeftTime{
					ComputerID:  booking.ComputerID,
					LeftSeconds: int64(leftDuration.Seconds()),
				})
			} else {
				s.repository.Update(&entities.Booking{
					ID:         booking.ID,
					UserID:     booking.UserID,
					ComputerID: booking.ComputerID,
					PackageID:  booking.PackageID,
					StartTime:  booking.StartTime,
					EndTime:    booking.EndTime,
					Status:     "finished",
					CreatedAt:  booking.CreatedAt,
				})
			}
		}
	}

	return leftTimes, nil
}

func (s *BookingService) GetComputerLeftOccupiedTime(id int64) (*entities.ComputerOccupiedLeftTime, error) {
	bookings, err := s.GetComputerBookings(id)
	if err != nil {
		return nil, err
	}

	now := time.Now().UTC()

	leftTime := entities.ComputerOccupiedLeftTime{
		ComputerID:  id,
		LeftSeconds: 0,
	}

	for _, booking := range bookings {
		if booking.Status == "active" {
			endTime, err := time.Parse(time.RFC3339, booking.EndTime)
			if err != nil {
				return nil, fmt.Errorf("invalid end_time format for booking %d: %w", booking.ID, err)
			}

			if endTime.After(now) {
				leftDuration := endTime.Sub(now)
				return &entities.ComputerOccupiedLeftTime{
					ComputerID:  booking.ComputerID,
					LeftSeconds: int64(leftDuration.Seconds()),
				}, nil
			} else {
				return &leftTime, nil
			}
		}
	}

	// No active booking found
	return &leftTime, nil
}

func (s *BookingService) RefreshBookings() {
	s.repository.RefreshStatus()
}

func (s *BookingService) CreateRandomActiveBookingEndingSoon() {
	s.repository.CreateRandomActiveBookingEndingSoon()
}
