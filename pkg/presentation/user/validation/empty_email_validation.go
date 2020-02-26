package validation

import (
	"errors"

	model2 "github.com/arpb2/C-3PO/pkg/domain/user/model"
)

func EmptyEmail(user *model2.AuthenticatedUser) error {
	if len(user.Email) == 0 {
		return errors.New("no 'email' provided")
	}

	return nil
}
