package entities

import (
	"errors"
	"fmt"
	"strings"
	"time"
)

type Booking struct {
	ID         int64  `json:"id,omitempty"`
	UserID     int64  `json:"user_id"`
	ComputerID int64  `json:"computer_id"`
	PackageID  int64  `json:"package_id"`
	StartTime  string `json:"start_time"`
	EndTime    string `json:"end_time"`
	Status     string `json:"status"`
	CreatedAt  string `json:"created_at,omitempty"`
}

type ComputerOccupiedLeftTime struct {
	ComputerID  int64 `json:"id"`
	LeftSeconds int64 `json:"left_time"`
}

func (b *Booking) Validate() error {
	allowedStatuses := map[string]bool{
		"pending":   true,
		"active":    true,
		"finished":  true,
		"cancelled": true,
	}

	// Validate IDs
	if b.UserID <= 0 {
		return errors.New("user_id must be a positive number")
	}
	if b.ComputerID <= 0 {
		return errors.New("computer_id must be a positive number")
	}

	// Validate Package
	if b.PackageID == 0 {
		return errors.New("package must not be zero")
	}
	if b.PackageID < 0 {
		return errors.New("package must not be negative")
	}

	// Parse StartTime and EndTime (ISO 8601 format)
	start, err := time.Parse(time.RFC3339, b.StartTime)
	if err != nil {
		return fmt.Errorf("start_time must be in ISO 8601 format: %w", err)
	}
	end, err := time.Parse(time.RFC3339, b.EndTime)
	if err != nil {
		return fmt.Errorf("end_time must be in ISO 8601 format: %w", err)
	}

	now := time.Now().UTC()

	if !start.After(now) {
		return errors.New("start_time must be in the future")
	}
	if !end.After(start) {
		return errors.New("end_time must be after start_time")
	}

	if end.Sub(start)%time.Hour != 0 {
		return errors.New("booking duration must be a whole number of hours")
	}

	// Validate Status
	if strings.TrimSpace(b.Status) == "" {
		return errors.New("status must not be empty")
	}
	if !allowedStatuses[b.Status] {
		return fmt.Errorf("invalid status: %s", b.Status)
	}

	return nil
}
