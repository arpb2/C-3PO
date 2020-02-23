package validation

import (
	"errors"

	"github.com/arpb2/C-3PO/pkg/domain/model"
)

func IdProvided(user *model.AuthenticatedUser) error {
	if user.Id > 0 {
		return errors.New("client can't define user 'id'")
	}

	return nil
}
