package user

import (
	"context"
	"crypto/rand"
	user2 "github.com/arpb2/C-3PO/third_party/ent/user"
	"io"

	"github.com/arpb2/C-3PO/pkg/domain/http"

	"github.com/arpb2/C-3PO/pkg/domain/model/user"

	"github.com/arpb2/C-3PO/pkg/infrastructure/mysql"

	"github.com/arpb2/C-3PO/third_party/ent"
)

const (
	saltBytes = 32
)

func createUser(tx *ent.Tx, ctx context.Context, user user.User) (*ent.User, error) {
	result, err := tx.User.Create().
		SetName(user.Name).
		SetSurname(user.Surname).
		SetEmail(user.Email).
		SetType(user2.Type(user.Type)).
		Save(ctx)

	if err != nil {
		if ent.IsConstraintError(err) {
			return nil, http.CreateBadRequestError("constraint error")
		}
		return nil, err
	}

	if result == nil {
		return nil, http.CreateInternalError()
	}

	return result, nil
}

func createCredential(tx *ent.Tx, ctx context.Context, holder *ent.User, user user.AuthenticatedUser) error {
	salt := make([]byte, saltBytes)
	_, err := io.ReadFull(rand.Reader, salt)
	if err != nil {
		return err
	}

	hash, err := mysql.SaltHash([]byte(user.Password), salt)
	if err != nil {
		return err
	}

	credential, err := tx.Credential.Create().
		SetHolder(holder).
		SetSalt(salt).
		SetPasswordHash(hash).
		Save(ctx)

	if err != nil {
		if ent.IsConstraintError(err) {
			return http.CreateBadRequestError("constraint error")
		}
		return err
	}

	if credential == nil {
		return http.CreateInternalError()
	}
	return nil
}

func createClassroom(tx *ent.Tx, ctx context.Context, holder *ent.User, user user.AuthenticatedUser) error {
	cr, err := tx.Classroom.Create().
		SetTeacher(holder).
		Save(ctx)

	if err != nil {
		if ent.IsConstraintError(err) {
			return http.CreateBadRequestError("constraint error")
		}
		return err
	}

	if cr == nil {
		return http.CreateInternalError()
	}
	return nil
}

func create(dbClient *ent.Client, authUser user.AuthenticatedUser) (user.User, error) {
	var userModel user.User
	ctx := context.Background()
	tx, err := dbClient.Tx(ctx)

	if err != nil {
		return userModel, mysql.Rollback(tx, err)
	}

	result, err := createUser(tx, ctx, authUser.User)
	if err != nil {
		return userModel, mysql.Rollback(tx, err)
	}

	err = createCredential(tx, ctx, result, authUser)
	if err != nil {
		return userModel, mysql.Rollback(tx, err)
	}

	if authUser.Type == user.TypeTeacher {
		err = createClassroom(tx, ctx, result, authUser)
		if err != nil {
			return userModel, mysql.Rollback(tx, err)
		}
	}

	err = tx.Commit()

	if err != nil {
		return userModel, err
	}

	if err = mapToDTO(dbClient, result.ID, result, &userModel); err != nil {
		return userModel, err
	}

	return userModel, nil
}
