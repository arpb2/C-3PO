package user

import (
	"github.com/arpb2/C-3PO/api/model"
	userservice "github.com/arpb2/C-3PO/api/service/user"
)

func CreateService() userservice.Service {
	return &userService{}
}

type userService struct{}

func (u userService) GetUser(userId uint) (user model.User, err error) {
	panic("implement me")
}

func (u userService) CreateUser(user model.User) (result model.User, err error) {
	panic("implement me")
}

func (u userService) UpdateUser(user model.User) (result model.User, err error) {
	panic("implement me")
}

func (u userService) DeleteUser(userId uint) error {
	panic("implement me")
}
