package user

import "github.com/arpb2/C-3PO/api/model"

type Service interface {
	GetUser(userId uint) (user model.User, err error)

	CreateUser(user model.User) (result model.User, err error)

	UpdateUser(user model.User) (result model.User, err error)

	DeleteUser(userId uint) error
}
