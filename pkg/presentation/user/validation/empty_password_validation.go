package validation

import (
	"errors"

	model2 "github.com/arpb2/C-3PO/pkg/domain/user/model"
)

func EmptyPassword(user *model2.AuthenticatedUser) error {
	if len(user.Password) == 0 {
		return errors.New("no 'password' provided")
	}

	return nil
}
