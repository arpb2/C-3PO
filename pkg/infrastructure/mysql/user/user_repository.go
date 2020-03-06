package user

import (
	user3 "github.com/arpb2/C-3PO/pkg/data/repository/user"
	"github.com/arpb2/C-3PO/pkg/domain/model/user"
	"github.com/arpb2/C-3PO/third_party/ent"
)

func CreateUserRepository(dbClient *ent.Client) user3.Repository {
	return &userRepository{
		dbClient: dbClient,
	}
}

type userRepository struct {
	dbClient *ent.Client
}

func mapToDTO(userId uint, input *ent.User, output *user.User) {
	if input == nil {
		return
	}

	output.Id = userId
	output.Email = input.Email
	output.Name = input.Name
	output.Surname = input.Surname
}

func (u *userRepository) GetUser(userId uint) (user user.User, err error) {
	return get(u.dbClient, userId)
}

func (u *userRepository) CreateUser(user user.AuthenticatedUser) (user.User, error) {
	return create(u.dbClient, user)
}

func (u *userRepository) UpdateUser(user user.AuthenticatedUser) (result user.User, err error) {
	return update(u.dbClient, user)
}

func (u *userRepository) DeleteUser(userId uint) error {
	return del(u.dbClient, userId)
}
