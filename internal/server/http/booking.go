package httpserver

import (
	"encoding/json"
	"log/slog"
	"net/http"
)

func (s *Server) GetAllBookings(w http.ResponseWriter, r *http.Request) {
	bookings, err := s.service.BookingService.GetAllBookings()
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

func (s *Server) GetFinishedBookings(w http.ResponseWriter, r *http.Request) {
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

func (s *Server) GetPendingBookings(w http.ResponseWriter, r *http.Request) {
	bookings, err := s.service.BookingService.GetPendingBookings()
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
