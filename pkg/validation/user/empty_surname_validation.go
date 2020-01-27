package user

import (
	"errors"

	"github.com/arpb2/C-3PO/api/model"
)

func EmptySurname(user *model.AuthenticatedUser) error {
	if len(user.Surname) == 0 {
		return errors.New("no 'surname' provided")
	}

	return nil
}
