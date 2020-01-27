package user

import "github.com/arpb2/C-3PO/api/model"

type Validation func(user *model.AuthenticatedUser) error
