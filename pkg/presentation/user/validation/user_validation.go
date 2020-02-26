package validation

import (
	"errors"

	model2 "github.com/arpb2/C-3PO/pkg/domain/user/model"
)

func EmptyUser(user *model2.AuthenticatedUser) error {
	if user == nil {
		return errors.New("no user found")
	}
	return nil
}
