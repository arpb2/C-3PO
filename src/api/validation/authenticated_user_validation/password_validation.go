package authenticated_user_validation

import (
	"errors"
	"github.com/arpb2/C-3PO/src/api/model"
)

func PasswordValidation(user *model.AuthenticatedUser) error {
	if user.Password == "" {
		return errors.New("no 'password' provided")
	}
	return nil
}
