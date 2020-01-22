package session_validation

import (
	"errors"
	"github.com/arpb2/C-3PO/src/api/model"
)

func EmptyEmailValidation(user *model.AuthenticatedUser) error {
	if len(user.Email) == 0 {
		return errors.New("no 'email' provided")
	}

	return nil
}
