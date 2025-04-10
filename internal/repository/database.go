package repository

import (
	"booking_api/internal/entities"
	"context"
	"database/sql"
	"log/slog"
	"time"

	_ "github.com/lib/pq"
)

func OpenDB(dsn string) (*sql.DB, error) {
	// Retry logic: attempt to connect multiple times
	maxRetries := 3                  // Try 6 times (30 seconds total if we wait 5 seconds between retries)
	retryInterval := 5 * time.Second // Retry every 5 seconds
	var db *sql.DB
	var err error

	// Try connecting up to maxRetries times
	for i := 0; i < maxRetries; i++ {
		// Use sql.Open() to create an empty connection pool, using the DSN from the config
		// struct.
		db, err = sql.Open("postgres", dsn)
		if err == nil {
			// Create a context with a timeout for the Ping operation
			ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
			defer cancel()

			// Try to ping the database to check if it's available
			err = db.PingContext(ctx)
			// if err == nil {
			// 	// If the ping is successful, use the connection
			// 	postgresDB = db
			// 	slog.Info("PostgreSQL database connection established")
			// 	return postgresDB, nil
			// }
			if err == nil {
				return db, nil
			}
		}

		// If any error occurs, log it and retry after a delay
		//slog.Errorf("Failed to connect to PostgreSQL, retrying in %v... (attempt %d/%d)", retryInterval, i+1, maxRetries)
		time.Sleep(retryInterval)
	}

	// Return the sql.DB connection pool.
	slog.Error("Failed to connect to PostgreSQL", "attempts", maxRetries, "interval", retryInterval, "error", err)
	return nil, err
}

func GetAllComputerSpecs(db *sql.DB) ([]entities.ComputerSpecs, error) {
	query := `
		SELECT computer_id, cpu, gpu, ram, ssd, hdd, monitor, keyboard, headset, mouse, created_at 
		FROM computers_specs 
	`

	rows, err := db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var computers_specs []entities.ComputerSpecs

	for rows.Next() {
		var computer_specs entities.ComputerSpecs

		err = rows.Scan(
			&computer_specs.ID,
			&computer_specs.CPU,
			&computer_specs.GPU,
			&computer_specs.RAM,
			&computer_specs.SSD,
			&computer_specs.HDD,
			&computer_specs.Monitor,
			&computer_specs.Keyboard,
			&computer_specs.Headset,
			&computer_specs.Mouse,
			&computer_specs.CreatedAt,
		)

		if err != nil {
			return nil, err
		}

		computers_specs = append(computers_specs, computer_specs)
	}
	if err = rows.Err(); err != nil {
		return nil, err
	}

	return computers_specs, nil
}

func GetAllComputerStatuses(db *sql.DB) ([]entities.ComputerStatus, error) {
	query := `
		SELECT computer_id, status, created_at 
		FROM computers_statuses 
	`

	rows, err := db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var computers_statuses []entities.ComputerStatus

	for rows.Next() {
		var computer_status entities.ComputerStatus

		err = rows.Scan(
			&computer_status.ID,
			&computer_status.Status,
			&computer_status.UpdatedAt,
		)

		if err != nil {
			return nil, err
		}

		computers_statuses = append(computers_statuses, computer_status)
	}
	if err = rows.Err(); err != nil {
		return nil, err
	}

	return computers_statuses, nil
}

func GetAllPackages(db *sql.DB) ([]entities.User, error) {
	query := `
		SELECT user_id, login, email, status, role, balance, created_at 
		FROM users 
	`

	rows, err := db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var users []entities.User

	for rows.Next() {
		var user entities.User

		err = rows.Scan(
			&user.ID,
			&user.Login,
			&user.Email,
			&user.Status,
			&user.Role,
			&user.Balance,
			&user.CreatedAt,
		)

		if err != nil {
			return nil, err
		}

		users = append(users, user)
	}
	if err = rows.Err(); err != nil {
		return nil, err
	}

	return users, nil
}

func GetPackages(db *sql.DB) ([]entities.Package, error) {
	query := `
		SELECT package_name, price, startTime, endTime, created_at 
		FROM packages 
	`

	rows, err := db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var packages []entities.Package

	for rows.Next() {
		var package_ entities.Package

		err = rows.Scan(
			&package_.Name,
			&package_.Price,
			&package_.StartTime,
			&package_.EndTime,
			&package_.CreatedAt,
		)

		if err != nil {
			return nil, err
		}

		packages = append(packages, package_)

	}
	if err = rows.Err(); err != nil {
		return nil, err
	}

	return packages, nil
}
func GetAllUsers(db *sql.DB) ([]entities.User, error) {
	query := `
		SELECT user_id, login, email, status, role, balance, created_at 
		FROM users 
	`

	rows, err := db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var users []entities.User

	for rows.Next() {
		var user entities.User

		err = rows.Scan(
			&user.ID,
			&user.Login,
			&user.Email,
			&user.Status,
			&user.Role,
			&user.Balance,
			&user.CreatedAt,
		)

		if err != nil {
			return nil, err
		}

		users = append(users, user)
	}
	if err = rows.Err(); err != nil {
		return nil, err
	}

	return users, nil
}

func GetAllPendingBookings(db *sql.DB) ([]entities.Booking, error) {
	query := `
		SELECT user_id, computer_id, package_name, start_time, end_time, total_price, status, created_at 
		FROM pending_bookings 
	`

	rows, err := db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var bookings []entities.Booking

	for rows.Next() {
		var booking entities.Booking

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

func GetAllFinishedBookings(db *sql.DB) ([]entities.Booking, error) {
	query := `
		SELECT user_id, computer_id, package_name, start_time, end_time, total_price, status, created_at 
		FROM finished_bookings 
	`

	rows, err := db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var bookings []entities.Booking

	for rows.Next() {
		var booking entities.Booking

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
