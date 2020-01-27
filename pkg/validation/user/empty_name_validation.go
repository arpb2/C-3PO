package user_validation

import (
	"errors"

	"github.com/arpb2/C-3PO/api/model"
)

func EmptyName(user *model.AuthenticatedUser) error {
	if len(user.Name) == 0 {
		return errors.New("no 'name' provided")
	}

	return nil
}
