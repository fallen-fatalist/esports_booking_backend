package httpserver

import (
	"booking_api/internal/entities"
	"booking_api/internal/server"
	"booking_api/internal/service/implementation"
	"encoding/json"
	"log/slog"
	"net/http"
	"strconv"
)

func (s *Server) HandleComputers(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
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
	case http.MethodPost:
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
	default:
		// Handle Method Not Allowed
		w.Header().Set("Allow", "GET, POST")
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)

	}
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
	case http.MethodDelete:
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
	default:
		// Handle Method Not Allowed
		w.Header().Set("Allow", "GET, DELETE")
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)

	}

}

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
