package user

import (
	"context"

	user3 "github.com/arpb2/C-3PO/pkg/data/repository/user"
	"github.com/arpb2/C-3PO/pkg/domain/model/user"
	"github.com/arpb2/C-3PO/third_party/ent"
	"github.com/arpb2/C-3PO/third_party/ent/classroom"
	"github.com/arpb2/C-3PO/third_party/ent/level"
	user2 "github.com/arpb2/C-3PO/third_party/ent/user"
	"github.com/arpb2/C-3PO/third_party/ent/userlevel"
)

func CreateUserRepository(dbClient *ent.Client) user3.Repository {
	return &userRepository{
		dbClient: dbClient,
	}
}

type userRepository struct {
	dbClient *ent.Client
}

func mapToDTO(dbClient *ent.Client, userId uint, input *ent.User, output *user.User) error {
	if input == nil {
		return nil
	}

	output.Id = userId
	output.Email = input.Email
	output.Name = input.Name
	output.Surname = input.Surname
	output.Type = user.Type(input.Type)

	ctx := context.Background()
	ul, err := dbClient.UserLevel.
		Query().
		WithLevel().
		Where(
			userlevel.HasDeveloperWith(user2.ID(userId)),
		).
		Order(ent.Desc(level.FieldID)).
		First(ctx)

	if err != nil && !ent.IsNotFound(err) {
		return err
	}
	if ul != nil {
		output.CurrentLevel = ul.Edges.Level.ID
	}

	cr, err := dbClient.Classroom.
		Query().
		Where(
			classroom.Or(
				classroom.HasTeacherWith(user2.ID(userId)),
				classroom.HasStudentsWith(user2.ID(userId)),
			),
		).
		First(ctx)

	if err != nil && !ent.IsNotFound(err) {
		return err
	}
	if cr != nil {
		output.ClassroomID = cr.ID
	}
	return nil
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
