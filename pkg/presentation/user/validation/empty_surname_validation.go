package validation

import (
	"errors"

	model2 "github.com/arpb2/C-3PO/pkg/domain/user/model"
)

func EmptySurname(user *model2.AuthenticatedUser) error {
	if len(user.Surname) == 0 {
		return errors.New("no 'surname' provided")
	}

	return nil
}
