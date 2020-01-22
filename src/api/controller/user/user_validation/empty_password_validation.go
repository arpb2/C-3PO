package user_validation

import (
	"errors"
	"github.com/arpb2/C-3PO/src/api/model"
)

func EmptyPasswordValidation(user *model.AuthenticatedUser) error {
	if len(user.Password) == 0 {
		return errors.New("no 'password' provided")
	}

	return nil
}
