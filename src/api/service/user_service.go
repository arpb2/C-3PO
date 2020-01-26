package service

import "github.com/arpb2/C-3PO/src/api/model"

type UserService interface {

	GetUser(userId uint) (user *model.User, err error)

	CreateUser(authenticatedUser *model.AuthenticatedUser) (user *model.User, err error)

	UpdateUser(authenticatedUser *model.AuthenticatedUser) (user *model.User, err error)

	DeleteUser(userId uint) error

}