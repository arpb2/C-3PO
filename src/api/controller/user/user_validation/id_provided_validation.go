package session_validation

import (
	"errors"
	"github.com/arpb2/C-3PO/src/api/model"
)

func IdProvidedValidation(user *model.AuthenticatedUser) error {
	if user.Id > 0 {
		return errors.New("client can't define user 'id'")
	}

	return nil
}
