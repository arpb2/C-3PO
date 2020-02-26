package service

import (
	model2 "github.com/arpb2/C-3PO/pkg/domain/user/model"
)

type Service interface {
	GetUser(userId uint) (user model2.User, err error)

	CreateUser(user model2.AuthenticatedUser) (result model2.User, err error)

	UpdateUser(user model2.AuthenticatedUser) (result model2.User, err error)

	DeleteUser(userId uint) error
}
