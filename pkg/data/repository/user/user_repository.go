package user

import (
	"github.com/arpb2/C-3PO/pkg/domain/model/user"
)

type Repository interface {
	GetUser(userId uint) (user user.User, err error)

	CreateUser(user user.AuthenticatedUser) (result user.User, err error)

	UpdateUser(user user.AuthenticatedUser) (result user.User, err error)

	DeleteUser(userId uint) error
}
