package httpserver

import (
	"encoding/json"
	"esports_club_booking/internal/entities"
	"esports_club_booking/internal/server"
	"esports_club_booking/internal/service/implementation"
	"log/slog"
	"net/http"
	"strconv"
)

// @Summary Get all computers
// @Description Returns a list of all computers
// @Tags Computers
// @Produce json
// @Success 200 {array} entities.Computer "List of computers"
// @Failure 500 {object} server.JSONAnswer "Internal server error"
// @Router /api/computers [get]
func (s *Server) getComputers(w http.ResponseWriter, r *http.Request) {
	computerSpecs, err := s.service.ComputerService.GetAllComputers()
	if err != nil {
		slog.Error("Error fetching computer specs:", err)
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(computerSpecs); err != nil {
		slog.Error("Error encoding JSON:", slog.Any("error", err))
		http.Error(w, "Failed to encode response", http.StatusInternalServerError)
		return
	}

}

// @Summary Create a new computer
// @Description Creates a new computer
// @Tags Computers
// @Accept json
// @Produce json
// @Param computer body entities.Computer true "Computer details"
// @Success 201 {object} server.JSONAnswer "Computer created successfully"
// @Failure 400 {object} server.JSONAnswer "Invalid request body"
// @Failure 500 {object} server.JSONAnswer "Internal server error"
// @Router /api/computers [post]
func (s *Server) postComputers(w http.ResponseWriter, r *http.Request) {
	var computer entities.Computer

	// Parse the JSON body into the struct
	if err := json.NewDecoder(r.Body).Decode(&computer); err != nil {
		http.Error(w, "Invalid request body: "+err.Error(), http.StatusBadRequest)
		return
	}
	defer r.Body.Close()

	id, err := s.service.ComputerService.CreateComputer(&computer)
	if err != nil {
		slog.Error("Error creating computer:", err)
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}

	jsonPayload := server.JSONAnswer{
		Message: "Computer created successfully",
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

func (s *Server) HandleComputers(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		s.getComputers(w, r)
	case http.MethodPost:
		s.postComputers(w, r)
	default:
		// Handle Method Not Allowed
		w.Header().Set("Allow", "GET, POST")
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)

	}
}

// @Summary Get a specific computer
// @Description Returns a specific computer by ID
// @Tags Computers
// @Produce json
// @Param id path int64 true "Computer ID"
// @Success 200 {object} entities.Computer "Computer details"
// @Failure 400 {object} server.JSONAnswer "Invalid computer ID"
// @Failure 404 {object} server.JSONAnswer "Computer not found"
// @Failure 500 {object} server.JSONAnswer "Internal server error"
// @Router /api/computers/{id} [get]
func (s *Server) getComputer(w http.ResponseWriter, r *http.Request, id int64) {
	computer, err := s.service.ComputerService.GetComputer(id)
	if err != nil {
		if err == implementation.ErrComputerNotFound {
			jsonPayload := server.JSONAnswer{
				Message: "Computer not found",
				ID:      id,
			}

			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusNotFound)
			if err := json.NewEncoder(w).Encode(jsonPayload); err != nil {
				slog.Error("Error encoding JSON:", err)
				http.Error(w, "Failed to encode response", http.StatusInternalServerError)
				return
			}
			return
		}
		slog.Error("Error fetching computer specs:", err)
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(computer); err != nil {
		slog.Error("Error encoding JSON:", slog.Any("error", err))
		http.Error(w, "Failed to encode response", http.StatusInternalServerError)
		return
	}

}

// @Summary Delete a specific computer
// @Description Deletes a specific computer by ID
// @Tags Computers
// @Produce json
// @Param id path int64 true "Computer ID"
// @Success 204 {object} server.JSONAnswer "Computer deleted successfully"
// @Failure 400 {object} server.JSONAnswer "Invalid computer ID"
// @Failure 404 {object} server.JSONAnswer "Computer not found"
// @Failure 500 {object} server.JSONAnswer "Internal server error"
// @Router /api/computers/{id} [delete]
func (s *Server) deleteComputer(w http.ResponseWriter, r *http.Request, id int64) {
	err := s.service.ComputerService.DeleteComputer(id)
	if err != nil {
		if err == implementation.ErrComputerNotFound {
			jsonPayload := server.JSONAnswer{
				Message: "Computer not found",
				ID:      id,
			}

			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusNotFound)
			if err := json.NewEncoder(w).Encode(jsonPayload); err != nil {
				slog.Error("Error encoding JSON:", err)
				http.Error(w, "Failed to encode response", http.StatusInternalServerError)
				return
			}
			return
		}
		slog.Error("Error deleting computer:", err)
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}

	// For now, just respond with success
	w.WriteHeader(http.StatusNoContent)
}

func (s *Server) HandleComputer(w http.ResponseWriter, r *http.Request) {
	idStr := r.PathValue("id")
	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		slog.Error("Error parsing computer ID:", slog.Any("error", err))
		http.Error(w, "Invalid computer ID", http.StatusBadRequest)
		return
	}
	switch r.Method {
	case http.MethodGet:
		s.getComputer(w, r, id)
	case http.MethodDelete:
		s.deleteComputer(w, r, id)
	default:
		// Handle Method Not Allowed
		w.Header().Set("Allow", "GET, DELETE")
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)

	}

}

// @Summary Get all computer statuses
// @Description Returns a list of all computer statuses
// @Tags Computers
// @Accept json
// @Produce json
// @Success 200 {array} entities.ComputerStatus "List of computer statuses"
// @Failure 500 {object} server.JSONAnswer "Internal server error"
// @Router /api/computers/statuses [get]
func (s *Server) HandleComputerStatuses(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		w.Header().Set("Allow", "GET")
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}

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

// @Summary Get the status of a specific computer
// @Description Returns the status of the specified computer ID
// @Tags Computers
// @Accept json
// @Produce json
// @Param id path int64 true "Computer ID"
// @Success 200 {object} entities.ComputerStatus "Computer status"
// @Failure 400 {object} server.JSONAnswer "Invalid computer ID"
// @Failure 404 {object} server.JSONAnswer "Computer not found"
// @Failure 500 {object} server.JSONAnswer "Internal server error"
// @Router /api/computers/{id}/status [get]
func (s *Server) HandleComputerStatus(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		w.Header().Set("Allow", "GET")
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}

	idStr := r.PathValue("id")
	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		slog.Error("Error parsing computer ID:", slog.Any("error", err))
		http.Error(w, "Invalid computer ID", http.StatusBadRequest)
		return
	}

	status, err := s.service.ComputerService.GetComputerStatus(id)
	if err != nil {
		if err == implementation.ErrComputerNotFound {
			jsonPayload := server.JSONAnswer{
				Message: "Computer not found",
				ID:      id,
			}

			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusNotFound)
			if err := json.NewEncoder(w).Encode(jsonPayload); err != nil {
				slog.Error("Error encoding JSON:", err)
				http.Error(w, "Failed to encode response", http.StatusInternalServerError)
				return
			}
			return
		}
		slog.Error("Error fetching computer specs:", err)
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(status); err != nil {
		slog.Error("Error encoding JSON:", slog.Any("error", err))
		http.Error(w, "Failed to encode response", http.StatusInternalServerError)
		return
	}
}
