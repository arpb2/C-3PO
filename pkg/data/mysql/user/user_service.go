package user

import (
	model2 "github.com/arpb2/C-3PO/pkg/domain/user/model"
	"github.com/arpb2/C-3PO/pkg/domain/user/service"
	"github.com/arpb2/C-3PO/third_party/ent"
)

func CreateService(dbClient *ent.Client) service.Service {
	return &userService{
		dbClient: dbClient,
	}
}

type userService struct {
	dbClient *ent.Client
}

func mapToDTO(userId uint, input *ent.User, output *model2.User) {
	if input == nil {
		return
	}

	output.Id = userId
	output.Email = input.Email
	output.Name = input.Name
	output.Surname = input.Surname
}

func (u *userService) GetUser(userId uint) (user model2.User, err error) {
	return get(u.dbClient, userId)
}

func (u *userService) CreateUser(user model2.AuthenticatedUser) (model2.User, error) {
	return create(u.dbClient, user)
}

func (u *userService) UpdateUser(user model2.AuthenticatedUser) (result model2.User, err error) {
	return update(u.dbClient, user)
}

func (u *userService) DeleteUser(userId uint) error {
	return del(u.dbClient, userId)
}
