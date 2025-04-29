package postgres

import (
	"booking_api/internal/entities"
	"database/sql"
	"fmt"
	"time"
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
	pendingBookings, err := r.GetPendingBookings()
	if err != nil {
		return nil, err
	}

	finishedBookings, err := r.GetFinishedBookings()
	if err != nil {
		return nil, err
	}

	allBookings := make([]*entities.Booking, 0, len(pendingBookings)+len(finishedBookings))
	allBookings = append(allBookings, pendingBookings...)
	allBookings = append(allBookings, finishedBookings...)
	return allBookings, nil
}

func (r *BookingRepository) GetByID(id int64) (*entities.Booking, error) {
	var booking entities.Booking
	var startTime time.Time
	var endTime time.Time
	var createdAt time.Time

	// Try fetching from pending_bookings
	queryPending := `
		SELECT booking_id, user_id, computer_id, package_id, start_time, end_time, status, created_at
		FROM pending_bookings
		WHERE booking_id = $1
	`
	row := r.db.QueryRow(queryPending, id)
	err := row.Scan(
		&booking.ID,
		&booking.UserID,
		&booking.ComputerID,
		&booking.PackageID,
		&startTime,
		&endTime,
		&booking.Status,
		&createdAt,
	)

	if err == sql.ErrNoRows {
		// Not found in pending_bookings, try finished_bookings
		queryFinished := `
			SELECT booking_id, user_id, computer_id, package_id, start_time, end_time, status, created_at
			FROM finished_bookings
			WHERE booking_id = $1
		`
		row = r.db.QueryRow(queryFinished, id)
		err = row.Scan(
			&booking.ID,
			&booking.UserID,
			&booking.ComputerID,
			&booking.PackageID,
			&startTime,
			&endTime,
			&booking.Status,
			&createdAt,
		)

		if err == sql.ErrNoRows {
			return nil, fmt.Errorf("booking with id %d not found", id)
		}
		if err != nil {
			return nil, fmt.Errorf("failed to fetch booking from finished_bookings: %w", err)
		}
	} else if err != nil {
		return nil, fmt.Errorf("failed to fetch booking from pending_bookings: %w", err)
	}

	// Set time fields to ISO 8601
	booking.StartTime = startTime.UTC().Format(time.RFC3339)
	booking.EndTime = endTime.UTC().Format(time.RFC3339)
	booking.CreatedAt = createdAt.UTC().Format(time.RFC3339)

	return &booking, nil
}

func (r *BookingRepository) Create(booking *entities.Booking) (int64, error) {
	startTime, err := time.Parse(time.RFC3339, booking.StartTime)
	if err != nil {
		return 0, fmt.Errorf("invalid start_time: %w", err)
	}

	endTime, err := time.Parse(time.RFC3339, booking.EndTime)
	if err != nil {
		return 0, fmt.Errorf("invalid end_time: %w", err)
	}

	var query string
	var id int64

	switch booking.Status {
	case "pending", "active":
		query = `
			INSERT INTO pending_bookings 
			(user_id, computer_id, package_id, start_time, end_time, status, created_at)
			VALUES ($1, $2, $3, $4, $5, $6, NOW())
			RETURNING booking_id
		`

	case "finished", "cancelled":
		query = `
			INSERT INTO finished_bookings 
			(user_id, computer_id, package_id, start_time, end_time, status, created_at)
			VALUES ($1, $2, $3, $4, $5, $6, NOW())
			RETURNING booking_id
		`

	default:
		return 0, fmt.Errorf("invalid status: %s", booking.Status)
	}

	err = r.db.QueryRow(
		query,
		booking.UserID,
		booking.ComputerID,
		booking.PackageID,
		startTime,
		endTime,
		booking.Status,
	).Scan(&id)

	if err != nil {
		return 0, err
	}

	return id, nil
}

func (r *BookingRepository) Update(booking *entities.Booking) (int64, error) {
	return 0, nil
}

func (r *BookingRepository) Delete(int64) error {
	return nil
}

func (r *BookingRepository) GetPendingBookings() ([]*entities.Booking, error) {
	query := `
		SELECT booking_id, user_id, computer_id, package_id, start_time, end_time, status, created_at 
		FROM pending_bookings 
	`

	rows, err := r.db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var bookings []*entities.Booking

	for rows.Next() {
		var booking entities.Booking

		err = rows.Scan(
			&booking.ID,
			&booking.UserID,
			&booking.ComputerID,
			&booking.PackageID,
			&booking.StartTime,
			&booking.EndTime,
			&booking.Status,
			&booking.CreatedAt,
		)

		if err != nil {
			return nil, err
		}

		bookings = append(bookings, &booking)
	}
	if err = rows.Err(); err != nil {
		return nil, err
	}

	return bookings, nil
}

func (r *BookingRepository) GetFinishedBookings() ([]*entities.Booking, error) {
	query := `
		SELECT booking_id, user_id, computer_id, package_id, start_time, end_time, status, created_at 
		FROM finished_bookings 
	`

	rows, err := r.db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var bookings []*entities.Booking

	for rows.Next() {
		var booking entities.Booking

		err = rows.Scan(
			&booking.ID,
			&booking.UserID,
			&booking.ComputerID,
			&booking.PackageID,
			&booking.StartTime,
			&booking.EndTime,
			&booking.Status,
			&booking.CreatedAt,
		)

		if err != nil {
			return nil, err
		}

		bookings = append(bookings, &booking)
	}
	if err = rows.Err(); err != nil {
		return nil, err
	}

	return bookings, nil
}

func (r *BookingRepository) refreshStatus() {
	return
}
