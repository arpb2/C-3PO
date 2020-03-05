package validation

import (
	"errors"

	"github.com/arpb2/C-3PO/pkg/domain/model/user"
)

func EmptySurname(user *user.AuthenticatedUser) error {
	if len(user.Surname) == 0 {
		return errors.New("no 'surname' provided")
	}

	return nil
}
