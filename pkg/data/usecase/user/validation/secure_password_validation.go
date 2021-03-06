package validation

import (
	"errors"

	"github.com/arpb2/C-3PO/pkg/domain/model/user"
)

func SecurePassword(user *user.AuthenticatedUser) error {
	if len(user.Password) > 0 && len(user.Password) < 8 {
		return errors.New("password must have at least 8 characters")
	}

	return nil
}
