package validation

import (
	"errors"

	"github.com/arpb2/C-3PO/pkg/domain/model/user"
)

func EmptyEmail(user *user.AuthenticatedUser) error {
	if len(user.Email) == 0 {
		return errors.New("no 'email' provided")
	}

	return nil
}
