package service

import (
	"booking_api/internal/entities"
)

type ComputerBooking struct {
	ComputerID int64              `json:"id"`
	Status     string             `json:"status"`
	CPU        string             `json:"cpu"`
	GPU        string             `json:"gpu"`
	RAM        string             `json:"ram"`
	SSD        string             `json:"ssd"`
	HDD        string             `json:"hdd"`
	Monitor    string             `json:"monitor"`
	Keyboard   string             `json:"keyboard"`
	Headset    string             `json:"headset"`
	Mouse      string             `json:"mouse"`
	CreatedAt  string             `json:"created_at"`
	UpdatedAt  string             `json:"updated_at"`
	Bookings   []entities.Booking `json:"bookings"`
}

type ComputerOccupiedLeftTime struct {
	ComputerID  int64 `json:"id"`
	LeftSeconds int64 `json:"left_time"`
}
