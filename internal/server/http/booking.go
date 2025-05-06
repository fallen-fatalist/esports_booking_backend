package httpserver

import (
	"encoding/json"
	"esports_club_booking/internal/entities"
	"esports_club_booking/internal/server"
	"esports_club_booking/internal/service"
	"log/slog"
	"net/http"
	"strconv"
)

// @Summary Get all bookings with time format "2006-01-02T15:04:05Z07:00"
// @Description Returns a list of bookings with various statuses
// @Tags Bookings
// @Accept json
// @Produce json
// @Success 200 {array} entities.Booking
// @Failure 500 {object} server.JSONAnswer "Internal server error"
// @Router /api/bookings [get]
func (s *Server) getBookings(w http.ResponseWriter, r *http.Request) {
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

	if err := json.NewEncoder(w).Encode(bookings); err != nil {
		slog.Error("Error encoding JSON:", slog.Any("error", err))
		http.Error(w, "Failed to encode response", http.StatusInternalServerError)
		return
	}
}

// @Summary Create a new booking
// @Description Creates a new booking with the given details
// @Tags Bookings
// @Accept json
// @Produce json
// @Param booking body entities.Booking true "Booking details"
// @Success 201 {object} server.JSONAnswer "Booking created successfully"
// @Failure 400 {object} server.JSONAnswer "Invalid request body"
// @Failure 500 {object} server.JSONAnswer "Internal server error"
// @Router /api/bookings [post]
func (s *Server) postBookings(w http.ResponseWriter, r *http.Request) {
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

}
func (s *Server) HandleBookings(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		s.getBookings(w, r)
	case http.MethodPost:
		s.postBookings(w, r)
	default:
		// Handle Method Not Allowed
		w.Header().Set("Allow", "GET, POST")
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}

// @Summary Get bookings for a specific computer with time format "2006-01-02T15:04:05Z07:00"
// @Description Returns a list of bookings for the specified computer ID
// @Tags Bookings
// @Accept json
// @Produce json
// @Param id path int64 true "Computer ID"
// @Success 200 {array} entities.Booking "List of bookings for the computer"
// @Failure 400 {object} server.JSONAnswer "Invalid computer ID"
// @Failure 500 {object} server.JSONAnswer "Internal server error"
// @Router /api/bookings/{id} [get]
func (s *Server) getComputerBookings(w http.ResponseWriter, r *http.Request) {
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

}
func (s *Server) HandleComputerBookings(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		s.getComputerBookings(w, r)
	default:
		// Handle Method Not Allowed
		w.Header().Set("Allow", "GET")
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)

	}
}

// @Summary Get finished bookings with time format "2006-01-02T15:04:05Z07:00"
// @Description Returns a list of finished bookings
// @Tags Bookings
// @Accept json
// @Produce json
// @Success 200 {array} entities.Booking "List of finished bookings"
// @Failure 500 {object} server.JSONAnswer "Internal server error"
// @Router /api/bookings/finished [get]
func (s *Server) getFinishedBookings(w http.ResponseWriter, r *http.Request) {
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
func (s *Server) HandleFinishedBookings(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		s.getFinishedBookings(w, r)
	default:
		w.Header().Set("Allow", "GET")
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}

// @Summary Get pending bookings with time format "2006-01-02T15:04:05Z07:00"
// @Description Returns a list of pending bookings
// @Tags Bookings
// @Accept json
// @Produce json
// @Success 200 {array} entities.Booking "List of pending bookings"
// @Failure 500 {object} server.JSONAnswer "Internal server error"
// @Router /api/bookings/pending [get]
func (s *Server) getPendingBookings(w http.ResponseWriter, r *http.Request) {
	bookings, err := s.service.BookingService.GetPendingBookings()
	if err != nil {
		slog.Error("Error fetching pending bookings:", slog.Any("error", err))
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
	switch r.Method {
	case http.MethodGet:
		s.getPendingBookings(w, r)
	default:
		w.Header().Set("Allow", "GET")
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}

// @Summary Get active bookings with time format "2006-01-02T15:04:05Z07:00"
// @Description Returns a list of active bookings
// @Tags Bookings
// @Accept json
// @Produce json
// @Success 200 {array} entities.Booking "List of active bookings"
// @Failure 500 {object} server.JSONAnswer "Internal server error"
// @Router /api/bookings/active [get]
func (s *Server) getActiveBookings(w http.ResponseWriter, r *http.Request) {
	activeBookings, err := s.service.BookingService.GetActiveBookings()
	if err != nil {
		slog.Error("Error fetching active bookings:", slog.Any("error", err))
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(activeBookings); err != nil {
		slog.Error("Error encoding JSON:", err)
		http.Error(w, "Failed to encode response", http.StatusInternalServerError)
	}

}

func (s *Server) HandleActiveBookings(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		s.getActiveBookings(w, r)
	default:
		w.Header().Set("Allow", "GET")
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}

// @Summary Get cancelled bookings with time format "2006-01-02T15:04:05Z07:00"
// @Description Returns a list of cancelled bookings
// @Tags Bookings
// @Accept json
// @Produce json
// @Success 200 {array} entities.Booking "List of cancelled bookings"
// @Failure 500 {object} server.JSONAnswer "Internal server error"
// @Router /api/bookings/cancelled [get]
func (s *Server) getCancelledBookings(w http.ResponseWriter, r *http.Request) {
	cancelledBookings, err := s.service.BookingService.GetCancelledBookings()
	if err != nil {
		slog.Error("Error fetching cancelled bookings:", slog.Any("error", err))
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(cancelledBookings); err != nil {
		slog.Error("Error encoding JSON:", err)
		http.Error(w, "Failed to encode response", http.StatusInternalServerError)
	}

}

func (s *Server) HandleCancelledBookings(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		s.getCancelledBookings(w, r)
	default:
		w.Header().Set("Allow", "GET")
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}

// @Summary Get the left time for a specific computer in seconds
// @Description Returns the left time for the specified computer ID
// @Tags Bookings
// @Accept json
// @Produce json
// @Param id path int64 true "Computer ID"
// @Success 200 {object} entities.ComputerOccupiedLeftTime "Left time for the computer"
// @Failure 400 {object} server.JSONAnswer "Invalid computer ID"
// @Failure 500 {object} server.JSONAnswer "Internal server error"
// @Router /api/bookings/{id}/left [get]
func (s *Server) HandleComputerLeftTime(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		idStr := r.PathValue("id")
		id, err := strconv.ParseInt(idStr, 10, 64)
		if err != nil {
			http.Error(w, "Invalid computer ID", http.StatusBadRequest)
			return
		}
		computerLeftTime, err := s.service.BookingService.GetComputerLeftOccupiedTime(id)
		if err != nil {
			slog.Error("Error fetching computer left time:", slog.Any("error", err))
			http.Error(w, "Internal server error", http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		if err := json.NewEncoder(w).Encode(computerLeftTime); err != nil {
			slog.Error("Error encoding JSON:", err)
			http.Error(w, "Failed to encode response", http.StatusInternalServerError)
		}
	default:
		w.Header().Set("Allow", "GET")
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}

// @Summary Get computers left time in seconds
// @Description Returns a list of computers left time
// @Tags Bookings
// @Accept json
// @Produce json
// @Success 200 {array} entities.ComputerOccupiedLeftTime "List of computers left time"
// @Failure 500 {object} server.JSONAnswer "Internal server error"
// @Router /api/bookings/computers-left-time [get]
func (s *Server) HandleComputersLeftTime(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		computersLeftTime, err := s.service.BookingService.GetComputersLeftOccupiedTime()
		if err != nil {
			slog.Error("Error fetching computers left time:", slog.Any("error", err))
			http.Error(w, "Internal server error", http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		if err := json.NewEncoder(w).Encode(computersLeftTime); err != nil {
			slog.Error("Error encoding JSON:", err)
			http.Error(w, "Failed to encode response", http.StatusInternalServerError)
		}
	default:
		w.Header().Set("Allow", "GET")
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}
