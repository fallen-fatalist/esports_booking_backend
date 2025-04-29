package entities

import (
	"errors"
	"fmt"
	"strings"
)

type Computer struct {
	ID        int64  `json:"id,omitempty"`
	Status    string `json:"status"`
	CPU       string `json:"cpu"`
	GPU       string `json:"gpu"`
	RAM       string `json:"ram"`
	SSD       string `json:"ssd"`
	HDD       string `json:"hdd"`
	Monitor   string `json:"monitor"`
	Keyboard  string `json:"keyboard"`
	Headset   string `json:"headset"`
	Mouse     string `json:"mouse"`
	Mousepad  string `json:"mousepad"`
	CreatedAt string `json:"created_at,omitempty"`
	UpdatedAt string `json:"updated_at,omitempty"`
}

type ComputerStatus struct {
	ID        int64  `json:"id"`
	Status    string `json:"status"`
	UpdatedAt string `json:"updated_at"`
}

type Status string

const (
	Busy        Status = "busy"
	Available   Status = "available"
	Pending     Status = "pending"
	NotWorking  Status = "not working"
	UnderRepair Status = "under repair"
)

func (s Status) IsValid() bool {
	switch s {
	case "busy", "available", "pending":
		return true
	case "not working", "under repair":
		return true
	default:
		return false
	}
}

var (
	ErrInvalidStatus = errors.New("invalid status for computer provided")
)

func (computer *Computer) Validate() error {

	allowedStatuses := map[string]bool{
		"available":    true,
		"pending":      true,
		"busy":         true,
		"not working":  true,
		"under repair": true,
	}

	// Validate Status
	if !allowedStatuses[computer.Status] {
		return ErrInvalidStatus
	}

	// Validate all other fields
	fields := map[string]string{
		"cpu":      computer.CPU,
		"gpu":      computer.GPU,
		"ram":      computer.RAM,
		"ssd":      computer.SSD,
		"hdd":      computer.HDD,
		"monitor":  computer.Monitor,
		"keyboard": computer.Keyboard,
		"headset":  computer.Headset,
		"mouse":    computer.Mouse,
		"mousepad": computer.Mousepad,
	}

	for fieldName, value := range fields {
		if strings.TrimSpace(value) == "" {
			return fmt.Errorf("field '%s' must not be empty", fieldName)
		}
		if len(value) > 30 {
			return fmt.Errorf("field '%s' must not exceed 30 characters", fieldName)
		}
	}

	return nil
}
