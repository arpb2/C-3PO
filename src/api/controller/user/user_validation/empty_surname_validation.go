package session_validation

import (
	"errors"
	"github.com/arpb2/C-3PO/src/api/model"
)

func EmptySurnameValidation(user *model.AuthenticatedUser) error {
	if len(user.Surname) == 0 {
		return errors.New("no 'surname' provided")
	}

	return nil
}
