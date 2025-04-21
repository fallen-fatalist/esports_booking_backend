package postgres

import (
	"booking_api/internal/utils"
	"context"
	"database/sql"
	"fmt"
	"log/slog"
	"os"
	"time"

	_ "github.com/lib/pq"
)

const (
	maxRetries    = 3               // Try 3 times (15 seconds total if we wait 5 seconds between retries)
	retryInterval = 5 * time.Second // Retry every 5 seconds
)

var (
	dbHost     string = os.Getenv("DB_HOST")
	dbUser     string = os.Getenv("DB_USER")
	dbPassword string = os.Getenv("DB_PASSWORD")
	dbName     string = os.Getenv("DB_NAME")
	dbPort     string = os.Getenv("DB_PORT")
	dbVars            = []string{dbHost, dbName, dbPassword, dbHost, dbPort, dbUser}
)

func OpenDB() (*sql.DB, error) {
	if utils.IsAnyEmpty(dbVars) {
		panic("Some database environment variable is empty")
	}

	var db *sql.DB
	var err error

	dsn := fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=disable", dbUser, dbPassword, dbHost, dbPort, dbName)
	// Try connecting up to maxRetries times
	for range maxRetries {
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
