package validation

import (
	"errors"

	"github.com/arpb2/C-3PO/pkg/domain/model"
)

func EmptyPassword(user *model.AuthenticatedUser) error {
	if len(user.Password) == 0 {
		return errors.New("no 'password' provided")
	}

	return nil
}