package authenticated_user_validation

import "github.com/arpb2/C-3PO/src/api/model"

type Validation func(user *model.AuthenticatedUser) error
