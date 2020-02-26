package validation

import (
	"errors"

	model2 "github.com/arpb2/C-3PO/pkg/domain/user/model"
)

func IdProvided(user *model2.AuthenticatedUser) error {
	if user.Id > 0 {
		return errors.New("client can't define user 'id'")
	}

	return nil
}
