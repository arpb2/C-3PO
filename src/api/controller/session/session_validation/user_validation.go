package session_validation

import (
	"errors"
	"github.com/arpb2/C-3PO/src/api/model"
)

func UserValidation(user *model.AuthenticatedUser) error {
	if user == nil {
		return errors.New("no user found")
	}
	return nil
}
