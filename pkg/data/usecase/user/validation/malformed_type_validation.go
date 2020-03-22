package validation

import (
	"errors"

	"github.com/arpb2/C-3PO/pkg/domain/model/user"
)

func MalformedType(u *user.AuthenticatedUser) error {
	if u.Type != user.TypeTeacher && u.Type != user.TypeStudent {
		return errors.New("'type' is malformed")
	}

	return nil
}
