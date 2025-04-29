package entities

import "errors"

var (
	ErrNegativeID = errors.New("negative id provided")
	ErrZeroID     = errors.New("zero id provided")
)

func ValidateID(id int64) error {
	if id < 0 {
		return ErrNegativeID
	} else if id == 0 {
		return ErrZeroID
	}

	return nil
}
