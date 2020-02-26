package validation

import (
	"errors"

	model2 "github.com/arpb2/C-3PO/pkg/domain/user/model"
)

func SecurePassword(user *model2.AuthenticatedUser) error {
	if len(user.Password) > 0 && len(user.Password) < 8 {
		return errors.New("password must have at least 8 characters")
	}

	return nil
}
