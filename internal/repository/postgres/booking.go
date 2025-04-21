package postgres

import (
	"booking_api/internal/entities"
	"database/sql"
)

type BookingRepository struct {
	db *sql.DB
}

func NewBookingRepository(db *sql.DB) (*BookingRepository, error) {
	if db != nil {
		return &BookingRepository{db}, nil
	}

	panic("Empty database provided to Booking repository constructor")
}

func (r *BookingRepository) GetAll() ([]*entities.Booking, error) {
	return nil, nil
}

func (r *BookingRepository) GetByID(int64) (*entities.Booking, error) {
	return nil, nil
}

func (r *BookingRepository) Create(booking *entities.Booking) (int64, error) {
	return 0, nil
}

func (r *BookingRepository) Update(booking *entities.Booking) (int64, error) {
	return 0, nil
}

func (r *BookingRepository) Delete(int64) error {
	return nil
}

func (r *BookingRepository) GetPendingBookings() ([]*entities.Booking, error) {
	query := `
		SELECT user_id, computer_id, package_name, start_time, end_time, total_price, status, created_at 
		FROM pending_bookings 
	`

	rows, err := r.db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var bookings []*entities.Booking

	for rows.Next() {
		var booking *entities.Booking

		err = rows.Scan(
			&booking.ID,
			&booking.ComputerID,
			&booking.Package,
			&booking.StartTime,
			&booking.EndTime,
			&booking.TotalPrice,
			&booking.Status,
			&booking.CreatedAt,
		)

		if err != nil {
			return nil, err
		}

		bookings = append(bookings, booking)
	}
	if err = rows.Err(); err != nil {
		return nil, err
	}

	return bookings, nil
}

func (r *BookingRepository) GetFinishedBookings() ([]*entities.Booking, error) {
	query := `
		SELECT user_id, computer_id, package_name, start_time, end_time, total_price, status, created_at 
		FROM finished_bookings 
	`

	rows, err := r.db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var bookings []*entities.Booking

	for rows.Next() {
		var booking *entities.Booking

		err = rows.Scan(
			&booking.ID,
			&booking.ComputerID,
			&booking.Package,
			&booking.StartTime,
			&booking.EndTime,
			&booking.TotalPrice,
			&booking.Status,
			&booking.CreatedAt,
		)

		if err != nil {
			return nil, err
		}

		bookings = append(bookings, booking)
	}
	if err = rows.Err(); err != nil {
		return nil, err
	}

	return bookings, nil
}

func (r *BookingRepository) RefreshStatus() {
	return
}
