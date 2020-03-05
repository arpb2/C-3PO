package validation

import (
	"errors"

	"github.com/arpb2/C-3PO/pkg/domain/model/user"
)

func EmptyName(user *user.AuthenticatedUser) error {
	if len(user.Name) == 0 {
		return errors.New("no 'name' provided")
	}

	return nil
}
