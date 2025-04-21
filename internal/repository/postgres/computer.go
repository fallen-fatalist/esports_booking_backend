package postgres

import (
	"booking_api/internal/entities"
	"database/sql"
)

type ComputerRepository struct {
	db *sql.DB
}

func NewComputerRepository(db *sql.DB) (*ComputerRepository, error) {
	if db != nil {
		return &ComputerRepository{db}, nil
	}

	panic("Empty database provided to Computer repository constructor")
}

func (r *ComputerRepository) GetAll() ([]*entities.Computer, error) {
	query := `
		SELECT 
			computer_id, 
			status, 
			cpu, 
			gpu, 
			ram, 
			ssd, 
			hdd, 
			monitor, 
			keyboard, 
			headset,
			mouse, 
			created_at, 
			updated_at
		FROM computers 
	`

	rows, err := r.db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var computers_specs []*entities.Computer

	for rows.Next() {
		var computer_specs *entities.Computer

		err = rows.Scan(
			&computer_specs.ID,
			&computer_specs.CPU,
			&computer_specs.GPU,
			&computer_specs.RAM,
			&computer_specs.SSD,
			&computer_specs.HDD,
			&computer_specs.Monitor,
			&computer_specs.Keyboard,
			&computer_specs.Headset,
			&computer_specs.Mouse,
			&computer_specs.CreatedAt,
		)

		if err != nil {
			return nil, err
		}

		computers_specs = append(computers_specs, computer_specs)
	}
	if err = rows.Err(); err != nil {
		return nil, err
	}

	return computers_specs, nil
}

func (r *ComputerRepository) GetByID(int64) (*entities.Computer, error) {
	return nil, nil
}

func (r *ComputerRepository) Create(*entities.Computer) (int64, error) {
	return 0, nil
}

func (r *ComputerRepository) Update(computer *entities.Computer) (int64, error) {
	return 0, nil
}

func (r *ComputerRepository) Delete(int64) error {
	return nil
}

// func (r *ComputerRepository) GetAllComputerStatuses() ([]entities.ComputerStatus, error) {
// 	query := `
// 		SELECT computer_id, status, updated_at
// 		FROM computers
// 	`

// 	rows, err := r.db.Query(query)
// 	if err != nil {
// 		return nil, err
// 	}
// 	defer rows.Close()

// 	var computers_statuses []entities.ComputerStatus

// 	for rows.Next() {
// 		var computer_status entities.ComputerStatus

// 		err = rows.Scan(
// 			&computer_status.ID,
// 			&computer_status.Status,
// 			&computer_status.UpdatedAt,
// 		)

// 		if err != nil {
// 			return nil, err
// 		}

// 		computers_statuses = append(computers_statuses, computer_status)
// 	}
// 	if err = rows.Err(); err != nil {
// 		return nil, err
// 	}

// 	return computers_statuses, nil
// }
