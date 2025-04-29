package httpserver

import (
	"encoding/json"
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
	mux.HandleFunc("/api/computers/{id}/bookings", s.HandleComputerBookings)

	/* Bookings */
	mux.HandleFunc("/api/bookings", s.HandleBookings)

	//mux.HandleFunc("/api/users", s.GetUsers)
	// mux.Handle("/api/packages", s.GetAllPackages)
	// mux.Handle("/api/bookings/pending", http.HandlerFunc(PendingBookings))
	// mux.Handle("/api/bookings/finished", http.HandlerFunc(FinishedBookings))

	/* Middlewares attach */
	router := recoverPanic(mux)
	router = requestLogger(router)
	return router
}

func (s *Server) GetUsers(w http.ResponseWriter, r *http.Request) {
	users, err := s.service.UserService.GetAllUsers()
	if err != nil {
		slog.Error(err.Error())
		http.Error(w, "Internal server error, Please message administrators", http.StatusInternalServerError)
	}

	w.Header().Set("Content-Type", "application/json")

	if err := json.NewEncoder(w).Encode(users); err != nil {
		slog.Error("Error encoding JSON:", err)
		http.Error(w, "Failed to encode response", http.StatusInternalServerError)
	}
}

func (s *Server) Packages(w http.ResponseWriter, r *http.Request) {
	packages, err := s.service.PackageService.GetAllPackages()
	if err != nil {
		slog.Error("Error fetching packages:", err)
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(packages); err != nil {
		slog.Error("Error encoding JSON:", err)
		http.Error(w, "Failed to encode response", http.StatusInternalServerError)
	}
}
func (s *Server) AllBookings(w http.ResponseWriter, r *http.Request) {
	bookings, err := s.service.BookingService.GetAllBookings()
	if err != nil {
		slog.Error("Error fetching pending bookings:", err)
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(bookings); err != nil {
		slog.Error("Error encoding JSON:", err)
		http.Error(w, "Failed to encode response", http.StatusInternalServerError)
	}
}

func (s *Server) FinishedBookings(w http.ResponseWriter, r *http.Request) {
	bookings, err := s.service.BookingService.GetFinishedBookings()
	if err != nil {
		slog.Error("Error fetching finished bookings:", slog.Any("error", err))
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(bookings); err != nil {
		slog.Error("Error encoding JSON:", err)
		http.Error(w, "Failed to encode response", http.StatusInternalServerError)
	}
}
