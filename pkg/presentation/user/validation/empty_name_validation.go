package validation

import (
	"errors"

	model2 "github.com/arpb2/C-3PO/pkg/domain/user/model"
)

func EmptyName(user *model2.AuthenticatedUser) error {
	if len(user.Name) == 0 {
		return errors.New("no 'name' provided")
	}

	return nil
}
