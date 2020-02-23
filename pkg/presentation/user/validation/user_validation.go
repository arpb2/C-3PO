package validation

import (
	"errors"

	"github.com/arpb2/C-3PO/pkg/domain/model"
)

func EmptyUser(user *model.AuthenticatedUser) error {
	if user == nil {
		return errors.New("no user found")
	}
	return nil
}
