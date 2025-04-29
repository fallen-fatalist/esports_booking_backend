package httpserver

import (
	"booking_api/internal/entities"
	"booking_api/internal/server"
	"booking_api/internal/service"
	"encoding/json"
	"log/slog"
	"net/http"
	"strconv"
)

func (s *Server) HandleBookings(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		bookings, err := s.service.BookingService.GetAllBookings()
		if err != nil {
			slog.Error("Error fetching bookings:", err)
			http.Error(w, "Internal server error", http.StatusInternalServerError)
			return
		} else if len(bookings) == 0 {
			jsonPayload := server.JSONAnswer{
				Message: "No bookings",
			}
			w.Header().Set("Content-Type", "application/json")
			if err := json.NewEncoder(w).Encode(jsonPayload); err != nil {
				slog.Error("Error encoding JSON:", slog.Any("error", err))
				http.Error(w, "Failed to encode response", http.StatusInternalServerError)
				return
			}
		}

		w.Header().Set("Content-Type", "application/json")
		if err := json.NewEncoder(w).Encode(bookings); err != nil {
			slog.Error("Error encoding JSON:", slog.Any("error", err))
			http.Error(w, "Failed to encode response", http.StatusInternalServerError)
			return
		}
	case http.MethodPost:
		var booking entities.Booking

		// Parse the JSON body into the struct
		if err := json.NewDecoder(r.Body).Decode(&booking); err != nil {
			http.Error(w, "Invalid request body: "+err.Error(), http.StatusBadRequest)
			return
		}
		defer r.Body.Close()

		id, err := s.service.BookingService.CreateBooking(&booking)
		if err != nil {
			if err == service.ErrUnhandledError {
				http.Error(w, "Internal server error", http.StatusInternalServerError)
				return
			} else {
				jsonPayload := server.JSONAnswer{
					Message: err.Error(),
				}

				w.Header().Set("Content-Type", "application/json")
				w.WriteHeader(http.StatusBadRequest)
				err = json.NewEncoder(w).Encode(jsonPayload)
				if err != nil {
					slog.Error("Error encoding JSON:", err)
					http.Error(w, "Failed to encode response", http.StatusInternalServerError)
					return
				}
				return
			}
		}

		jsonPayload := server.JSONAnswer{
			Message: "Booking created successfully",
			ID:      id,
		}

		// For now, just respond with success
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusCreated)
		err = json.NewEncoder(w).Encode(jsonPayload)
		if err != nil {
			slog.Error("Error encoding JSON:", err)
			http.Error(w, "Failed to encode response", http.StatusInternalServerError)
			return
		}
	default:
		// Handle Method Not Allowed
		w.Header().Set("Allow", "GET, POST")
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)

	}
}

func (s *Server) HandleComputerBookings(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		idStr := r.PathValue("id")
		id, err := strconv.ParseInt(idStr, 10, 64)
		if err != nil {
			jsonPayload := server.JSONAnswer{
				Message: "Invalid computer ID",
			}
			w.WriteHeader(http.StatusBadRequest)
			w.Header().Set("Content-Type", "application/json")
			if err := json.NewEncoder(w).Encode(jsonPayload); err != nil {
				slog.Error("Error encoding JSON:", slog.Any("error", err))
				http.Error(w, "Failed to encode response", http.StatusInternalServerError)
				return
			}
			return
		}

		bookings, err := s.service.BookingService.GetComputerBookings(id)
		if err != nil {
			slog.Error("Error fetching bookings:", err)
			http.Error(w, "Internal server error", http.StatusInternalServerError)
			return
		} else if len(bookings) == 0 {
			jsonPayload := server.JSONAnswer{
				Message: "No bookings",
			}
			w.Header().Set("Content-Type", "application/json")
			if err := json.NewEncoder(w).Encode(jsonPayload); err != nil {
				slog.Error("Error encoding JSON:", slog.Any("error", err))
				http.Error(w, "Failed to encode response", http.StatusInternalServerError)
				return
			}
		}

		w.Header().Set("Content-Type", "application/json")
		if err := json.NewEncoder(w).Encode(bookings); err != nil {
			slog.Error("Error encoding JSON:", slog.Any("error", err))
			http.Error(w, "Failed to encode response", http.StatusInternalServerError)
			return
		}
	default:
		// Handle Method Not Allowed
		w.Header().Set("Allow", "GET")
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)

	}
}

func (s *Server) HandleFinishedBookings(w http.ResponseWriter, r *http.Request) {
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

func (s *Server) HandlePendingBookings(w http.ResponseWriter, r *http.Request) {
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
