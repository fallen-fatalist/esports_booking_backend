package server

import (
	"net/http"
)

type Server interface {
	// Handlers
	GetComputers(http.ResponseWriter, *http.Request)
	GetComputerStatuses(http.ResponseWriter, *http.Request)
	GetUsers(http.ResponseWriter, *http.Request)
	GetAllBookings(http.ResponseWriter, *http.Request)
	GetPendingBookings(http.ResponseWriter, *http.Request)
	GetFinishedBookings(http.ResponseWriter, *http.Request)

	// Service variables
	Routes() http.Handler
}
