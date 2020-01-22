package user_validation

import (
	"errors"
	"github.com/arpb2/C-3PO/src/api/model"
)

func SecurePasswordValidation(user *model.AuthenticatedUser) error {
	if len(user.Password) > 0 && len(user.Password) < 8 {
		return errors.New("password must have at least 8 characters")
	}

	return nil
}
