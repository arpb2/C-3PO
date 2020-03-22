package validation

import (
	"errors"

	"github.com/arpb2/C-3PO/pkg/domain/model/user"
)

func TypeProvided(user *user.AuthenticatedUser) error {
	if len(user.Type) > 0 {
		return errors.New("client can't define user 'type'")
	}

	return nil
}
