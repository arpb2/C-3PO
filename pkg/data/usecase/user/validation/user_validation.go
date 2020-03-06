package validation

import (
	"errors"

	"github.com/arpb2/C-3PO/pkg/domain/model/user"
)

func EmptyUser(user *user.AuthenticatedUser) error {
	if user == nil {
		return errors.New("no user found")
	}
	return nil
}
