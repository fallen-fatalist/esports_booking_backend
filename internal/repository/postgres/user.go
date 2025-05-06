package postgres

import (
	"esports_club_booking/internal/entities"
	"database/sql"
)

type UserRepository struct {
	db *sql.DB
}

func NewUserRepository(db *sql.DB) (*UserRepository, error) {
	if db != nil {
		return &UserRepository{db}, nil
	}

	panic("Empty database provided to User repository constructor")
}
func (r *UserRepository) GetAll() ([]*entities.User, error) {
	query := `
		SELECT user_id, login, email, status, role, balance, created_at 
		FROM users 
	`

	rows, err := r.db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var users []*entities.User

	for rows.Next() {
		var user *entities.User

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

func (r *UserRepository) GetByID(int64) (*entities.User, error) {
	return nil, nil
}

func (r *UserRepository) Create(*entities.User) (int64, error) {
	return 0, nil
}

func (r *UserRepository) Update(booking *entities.Booking) (int64, error) {
	return 0, nil
}

func (r *UserRepository) Delete(id int64) error {
	return nil
}
