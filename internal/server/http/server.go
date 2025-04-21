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

	mux.HandleFunc("/api/v1/computers", s.GetComputers)
	mux.HandleFunc("/api/v1/users", s.GetUsers)
	// mux.Handle("/api/v1/packages", s.GetAllPackages)
	// mux.Handle("/api/v1/bookings/pending", http.HandlerFunc(PendingBookings))
	// mux.Handle("/api/v1/bookings/finished", http.HandlerFunc(FinishedBookings))

	// Middlewares attach
	// recoveredMux := recoverPanic(mux)
	// router := requestLogger(recoveredMux)
	return mux
}

func (s *Server) GetComputers(w http.ResponseWriter, r *http.Request) {
	computerSpecs, err := s.service.ComputerService.GetAllComputers()
	if err != nil {
		slog.Error("Error fetching computer specs:", err)
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(computerSpecs); err != nil {
		slog.Error("Error encoding JSON:", err)
		http.Error(w, "Failed to encode response", http.StatusInternalServerError)
	}
}

func (s *Server) GetComputerStatuses(w http.ResponseWriter, r *http.Request) {
	computerStatuses, err := s.service.ComputerService.GetAllComputerStatuses()
	if err != nil {
		slog.Error("Error fetching computer statuses:", err)
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(computerStatuses); err != nil {
		slog.Error("Error encoding JSON:", err)
		http.Error(w, "Failed to encode response", http.StatusInternalServerError)
	}
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
	packages, err := s.service.ComputerService.GetAllPackages()
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
