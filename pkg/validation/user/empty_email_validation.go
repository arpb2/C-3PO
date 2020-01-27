package user

import (
	"errors"

	"github.com/arpb2/C-3PO/api/model"
)

func EmptyEmail(user *model.AuthenticatedUser) error {
	if len(user.Email) == 0 {
		return errors.New("no 'email' provided")
	}

	return nil
}
