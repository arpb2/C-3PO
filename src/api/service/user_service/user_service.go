package user_service

import (
	"github.com/arpb2/C-3PO/src/api/model"
	"github.com/arpb2/C-3PO/src/api/service"
)

func GetService() service.UserService {
	return globalRef
}

var globalRef = &userService{}

type userService struct {}
func (u userService) GetUser(userId uint) (user *model.User, err error) {
	panic("implement me")
}

func (u userService) CreateUser(authenticatedUser *model.AuthenticatedUser) (user *model.User, err error) {
	panic("implement me")
}

func (u userService) UpdateUser(authenticatedUser *model.AuthenticatedUser) (user *model.User, err error) {
	panic("implement me")
}

func (u userService) DeleteUser(userId uint) error {
	panic("implement me")
}

