package httpserver

import (
	"encoding/json"
	"log/slog"
	"net/http"
)

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
