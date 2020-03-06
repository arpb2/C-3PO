package validation

import (
	"github.com/arpb2/C-3PO/pkg/domain/model/user"
)

type Validation func(user *user.AuthenticatedUser) error
