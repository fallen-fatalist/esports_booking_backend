package api

import (
	"booking_api/internal/repository"
	"database/sql"
	"fmt"
	"log"
	"log/slog"
	"net/http"
	"os"
)

var (
	dbHost     string = os.Getenv("DB_HOST")
	dbUser     string = os.Getenv("DB_USER")
	dbPassword string = os.Getenv("DB_PASSWORD")
	dbName     string = os.Getenv("DB_NAME")
	dbPort     string = os.Getenv("DB_PORT")
	db         *sql.DB
	err        error
	port       string
)

func Run() {
	dsn := fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=disable", dbUser, dbPassword, dbHost, dbPort, dbName)

	db, err = repository.OpenDB(dsn)
	if err != nil {
		log.Fatal(err)
	}

	port = os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	mux := http.NewServeMux()
	mux.Handle("/", http.HandlerFunc(Hello))
	mux.Handle("/api/v1/computers/specs", http.HandlerFunc(ComputerSpecs))
	mux.Handle("/api/v1/computers/statuses", http.HandlerFunc(ComputerStatuses))
	mux.Handle("/api/v1/users", http.HandlerFunc(Users))
	mux.Handle("/api/v1/packages", http.HandlerFunc(Packages))
	mux.Handle("/api/v1/bookings/pending", http.HandlerFunc(PendingBookings))
	mux.Handle("/api/v1/bookings/finished", http.HandlerFunc(FinishedBookings))

	// Middlewares
	recoveredMux := recoverPanic(mux)
	loggedMux := requestLogger(recoveredMux)

	slog.Info("Starting server on: " + port + " port")
	http.ListenAndServe(fmt.Sprintf(":%s", port), loggedMux)
}
