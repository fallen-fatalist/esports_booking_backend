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
	rows, err := r.db.Query(`
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
			mousepad,
			created_at, 
			updated_at
		FROM computers
	`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var computers []*entities.Computer

	for rows.Next() {
		var c entities.Computer
		err := rows.Scan(
			&c.ID,
			&c.Status,
			&c.CPU,
			&c.GPU,
			&c.RAM,
			&c.SSD,
			&c.HDD,
			&c.Monitor,
			&c.Keyboard,
			&c.Headset,
			&c.Mouse,
			&c.Mousepad,
			&c.CreatedAt,
			&c.UpdatedAt,
		)
		if err != nil {
			return nil, err
		}
		computers = append(computers, &c)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return computers, nil
}

func (r *ComputerRepository) GetByID(id int64) (*entities.Computer, error) {
	res := r.db.QueryRow(`
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
			mousepad,
			created_at, 
			updated_at
		FROM computers
		WHERE computer_id = $1
	`, id)
	if res.Err() != nil {
		return nil, res.Err()
	}

	var c entities.Computer
	err := res.Scan(
		&c.ID,
		&c.Status,
		&c.CPU,
		&c.GPU,
		&c.RAM,
		&c.SSD,
		&c.HDD,
		&c.Monitor,
		&c.Keyboard,
		&c.Headset,
		&c.Mouse,
		&c.Mousepad,
		&c.CreatedAt,
		&c.UpdatedAt,
	)
	if err != nil {
		return nil, err
	}

	return &c, nil
}

func (r *ComputerRepository) Create(computer *entities.Computer) (int64, error) {
	query := `
		INSERT INTO computers 
		(status, cpu, gpu, ram, ssd, hdd, monitor, keyboard, headset, mouse, mousepad, created_at, updated_at)
		VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, NOW(), NOW())
		RETURNING computer_id
	`

	var id int64
	err := r.db.QueryRow(
		query,
		computer.Status,
		computer.CPU,
		computer.GPU,
		computer.RAM,
		computer.SSD,
		computer.HDD,
		computer.Monitor,
		computer.Keyboard,
		computer.Headset,
		computer.Mouse,
		computer.Mousepad,
	).Scan(&id)
	if err != nil {
		return 0, err
	}

	return id, nil
}

func (r *ComputerRepository) Update(computer *entities.Computer) (int64, error) {
	return 0, nil
}

func (r *ComputerRepository) Delete(id int64) (int64, error) {
	query := `DELETE FROM computers WHERE computer_id = $1`

	result, err := r.db.Exec(query, id)
	if err != nil {
		return 0, err
	}

	RowsAffected, err := result.RowsAffected()
	if err != nil {
		return 0, err
	}

	return RowsAffected, nil
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
