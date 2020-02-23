package validation

import "github.com/arpb2/C-3PO/pkg/domain/model"

type Validation func(user *model.AuthenticatedUser) error
