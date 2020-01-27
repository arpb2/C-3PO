package credential

import (
	"github.com/arpb2/C-3PO/api/model"
)

type Service interface {
	Store(user *model.AuthenticatedUser) error

	Retrieve(email, password string) (uint, error)
}
