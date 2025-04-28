package implementation

import (
	"booking_api/internal/entities"
	"booking_api/internal/repository"
	"booking_api/internal/service"
	"database/sql"
	"errors"
	"fmt"
	"log"
	"strings"
)

// Errors
var (
	ErrInvalidStatus    = errors.New("invalid status for computer provided")
	ErrComputerNotFound = errors.New("computer does not exist")
)

type ComputerService struct {
	repository repository.ComputerRepository
}

func NewComputerService(repository repository.ComputerRepository) (*ComputerService, error) {
	if repository == nil {
		log.Fatal("Nil repository provided to computer service")
	}
	return &ComputerService{repository: repository}, nil
}

func (s *ComputerService) GetAllComputers() ([]*entities.Computer, error) {
	return s.repository.GetAll()
}

func (s *ComputerService) GetComputer(id int64) (*entities.Computer, error) {
	if err := service.ValidateID(id); err != nil {
		return nil, err
	}
	comp, err := s.repository.GetByID(id)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, ErrComputerNotFound
		}
		return nil, err
	}
	return comp, nil
}

func (s *ComputerService) CreateComputer(computer *entities.Computer) (int64, error) {
	if err := ValidateComputer(computer); err != nil {
		return 0, err
	}
	return s.repository.Create(computer)
}

func (s *ComputerService) DeleteComputer(id int64) error {
	if err := service.ValidateID(id); err != nil {
		return err
	}

	RowsAffected, err := s.repository.Delete(id)
	if err != nil {
		return err
	} else if RowsAffected == 0 {
		return ErrComputerNotFound
	}
	return nil
}

func (s *ComputerService) GetAllComputerStatuses() ([]*entities.ComputerStatus, error) {
	comps, err := s.GetAllComputers()
	if err != nil {
		return nil, err
	}

	statuses := make([]*entities.ComputerStatus, len(comps))

	for idx, comp := range comps {
		statuses[idx] = &entities.ComputerStatus{
			ID:        comp.ID,
			Status:    comp.Status,
			UpdatedAt: comp.UpdatedAt,
		}
	}
	return statuses, nil
}

func (s *ComputerService) GetComputerStatus(id int64) (*entities.ComputerStatus, error) {
	if err := service.ValidateID(id); err != nil {
		return nil, err
	}

	comp, err := s.repository.GetByID(id)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, ErrComputerNotFound
		}
		return nil, err
	}

	status := &entities.ComputerStatus{
		ID:        comp.ID,
		Status:    comp.Status,
		UpdatedAt: comp.UpdatedAt,
	}

	return status, nil
}

func ValidateComputer(computer *entities.Computer) error {

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
