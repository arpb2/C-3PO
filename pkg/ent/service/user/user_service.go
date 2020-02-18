package user

import (
	"github.com/arpb2/C-3PO/api/model"
	userservice "github.com/arpb2/C-3PO/api/service/user"
	"github.com/arpb2/C-3PO/pkg/ent"
)

func CreateService(dbClient *ent.Client) userservice.Service {
	return &userService{
		dbClient: dbClient,
	}
}

type userService struct {
	dbClient *ent.Client
}

func mapToDTO(userId uint, input *ent.User, output *model.User) {
	if input == nil {
		return
	}

	output.Id = userId
	output.Email = input.Email
	output.Name = input.Name
	output.Surname = input.Surname
}

func (u *userService) GetUser(userId uint) (user model.User, err error) {
	return get(u.dbClient, userId)
}

func (u *userService) CreateUser(user model.AuthenticatedUser) (model.User, error) {
	return create(u.dbClient, user)
}

func (u *userService) UpdateUser(user model.AuthenticatedUser) (result model.User, err error) {
	return update(u.dbClient, user)
}

func (u *userService) DeleteUser(userId uint) error {
	return del(u.dbClient, userId)
}
