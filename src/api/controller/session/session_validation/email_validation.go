package session_validation

import (
	"errors"
	"github.com/arpb2/C-3PO/src/api/model"
)

func EmailValidation(user *model.AuthenticatedUser) error {
	if user.Email == "" {
		return errors.New("no 'email' provided")
	}
	return nil
}
