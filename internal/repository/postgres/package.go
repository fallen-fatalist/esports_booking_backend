package postgres

import (
	"esports_club_booking/internal/entities"
	"database/sql"
)

type PackageRepository struct {
	db *sql.DB
}

func NewPackageRepository(db *sql.DB) (*PackageRepository, error) {
	if db == nil {
		panic("Empty databaase provided to package repository")
	}
	return &PackageRepository{db}, nil
}
func (r *PackageRepository) GetAll() ([]*entities.Package, error) {
	return nil, nil
}

func (r *PackageRepository) GetByID(id int64) (*entities.Package, error) {
	return nil, nil
}

func (r *PackageRepository) Create(booking *entities.Package) (int64, error) {
	return 0, nil
}

func (r *PackageRepository) Update(booking *entities.Package) (int64, error) {
	return 0, nil
}

func (r *PackageRepository) Delete(id int64) error {
	return nil
}

func (r *ComputerRepository) GetAllPackages() ([]entities.User, error) {
	query := `
		SELECT user_id, login, email, status, role, balance, created_at 
		FROM users 
	`

	rows, err := r.db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var users []entities.User

	for rows.Next() {
		var user entities.User

		err = rows.Scan(
			&user.ID,
			&user.Login,
			&user.Email,
			&user.Status,
			&user.Role,
			&user.Balance,
			&user.CreatedAt,
		)

		if err != nil {
			return nil, err
		}

		users = append(users, user)
	}
	if err = rows.Err(); err != nil {
		return nil, err
	}

	return users, nil
}

func (r *ComputerRepository) GetPackages() ([]entities.Package, error) {
	query := `
		SELECT package_name, price, startTime, endTime, created_at 
		FROM packages 
	`

	rows, err := r.db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var packages []entities.Package

	for rows.Next() {
		var package_ entities.Package

		err = rows.Scan(
			&package_.Name,
			&package_.Price,
			&package_.StartTime,
			&package_.EndTime,
			&package_.CreatedAt,
		)

		if err != nil {
			return nil, err
		}

		packages = append(packages, package_)

	}
	if err = rows.Err(); err != nil {
		return nil, err
	}

	return packages, nil
}
