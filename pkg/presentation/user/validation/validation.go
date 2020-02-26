package validation

import (
	model2 "github.com/arpb2/C-3PO/pkg/domain/user/model"
)

type Validation func(user *model2.AuthenticatedUser) error
