package session_validation

import (
	"errors"
	"github.com/arpb2/C-3PO/src/api/model"
)

func EmptyNameValidation(user *model.AuthenticatedUser) error {
	if len(user.Name) == 0 {
		return errors.New("no 'name' provided")
	}

	return nil
}
