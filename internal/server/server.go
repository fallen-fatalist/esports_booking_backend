package server

import (
	"net/http"
)

type Server interface {
	// Handlers
	HandleComputers(http.ResponseWriter, *http.Request)
	HandleComputer(http.ResponseWriter, *http.Request)
	HandleComputerStatuses(http.ResponseWriter, *http.Request)
	HandleComputerStatus(http.ResponseWriter, *http.Request)
	GetUsers(http.ResponseWriter, *http.Request)
	GetAllBookings(http.ResponseWriter, *http.Request)
	GetPendingBookings(http.ResponseWriter, *http.Request)
	GetFinishedBookings(http.ResponseWriter, *http.Request)

	// Service variables
	Routes() http.Handler
}

type JSONAnswer struct {
	Message string `json:"message"`
	ID      int64  `json:"id,omitempty"`
}
