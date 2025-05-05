package httpserver

import (
	"log"
	"net/http"

	"log/slog"

	"booking_api/internal/service"
)

type Server struct {
	logger  *slog.Logger
	service *service.Service
}

func NewServer(
	service *service.Service,
	logger *slog.Logger,
) *Server {
	if service == nil {
		log.Fatal("Nil service provided to Server")
	} else if logger == nil {
		log.Fatal("Nil logger provided to Server")
	}
	return &Server{
		logger:  logger,
		service: service,
	}
}

func (s *Server) Routes() http.Handler {
	mux := http.NewServeMux()

	/* Computers */
	mux.HandleFunc("/api/computers", s.HandleComputers)
	mux.HandleFunc("/api/computers/{id}", s.HandleComputer)
	mux.HandleFunc("/api/computers/status", s.HandleComputerStatuses)
	mux.HandleFunc("/api/computers/{id}/status", s.HandleComputerStatus)

	/* Bookings */
	mux.HandleFunc("/api/bookings/{id}", s.HandleComputerBookings)
	mux.HandleFunc("/api/bookings", s.HandleBookings)
	mux.HandleFunc("/api/bookings/pending", s.HandlePendingBookings)
	mux.HandleFunc("/api/bookings/active", s.HandleActiveBookings)
	mux.HandleFunc("/api/bookings/finished", s.HandleFinishedBookings)
	mux.HandleFunc("/api/bookings/cancelled", s.HandleCancelledBookings)
	mux.HandleFunc("/api/bookings/left", s.HandleComputersLeftTime)
	mux.HandleFunc("/api/bookings/{id}/left", s.HandleComputerLeftTime)

	//mux.HandleFunc("/api/users", s.GetUsers)
	// mux.Handle("/api/packages", s.GetAllPackages)
	// mux.Handle("/api/bookings/pending", http.HandlerFunc(PendingBookings))
	// mux.Handle("/api/bookings/finished", http.HandlerFunc(FinishedBookings))

	/* Middlewares attach */
	router := recoverPanic(mux)
	router = requestLogger(router)
	return router
}
