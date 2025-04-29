package server

import (
	"net/http"
)

/* Implementation instead of abstraction */
type Server interface {
	// Computers
	HandleComputers(http.ResponseWriter, *http.Request)
	HandleComputer(http.ResponseWriter, *http.Request)
	HandleComputerStatuses(http.ResponseWriter, *http.Request)
	HandleComputerStatus(http.ResponseWriter, *http.Request)
	HandleComputerBookings(http.ResponseWriter, *http.Request)

	// Bookings
	HandleBookings(http.ResponseWriter, *http.Request)
	HandlePendingBookings(http.ResponseWriter, *http.Request)
	HandleFinishedBookings(http.ResponseWriter, *http.Request)

	// Users
	GetUsers(http.ResponseWriter, *http.Request)

	// Service variables
	Routes() http.Handler
}

type JSONAnswer struct {
	Message string `json:"message"`
	ID      int64  `json:"id,omitempty"`
}
