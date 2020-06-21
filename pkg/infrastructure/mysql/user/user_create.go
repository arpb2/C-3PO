package user

import (
	"context"
	"crypto/rand"
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
		Save(ctx)

	if err != nil {
		if ent.IsConstraintError(err) {
			return nil, http.CreateBadRequestError("constraint error")
		}
		return nil, mysql.Rollback(tx, err)
	}

	if result == nil {
		return nil, mysql.Rollback(tx, err)
	}

	return result, nil
}

func createCredential(tx *ent.Tx, ctx context.Context, holder *ent.User, user user.AuthenticatedUser) error {
	salt := make([]byte, saltBytes)
	_, err := io.ReadFull(rand.Reader, salt)
	if err != nil {
		return mysql.Rollback(tx, err)
	}

	hash, err := mysql.SaltHash([]byte(user.Password), salt)
	if err != nil {
		return mysql.Rollback(tx, err)
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
		return mysql.Rollback(tx, err)
	}

	if credential == nil {
		return mysql.Rollback(tx, err)
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
		return mysql.Rollback(tx, err)
	}

	if cr == nil {
		return mysql.Rollback(tx, err)
	}
	return nil
}

func create(dbClient *ent.Client, authUser user.AuthenticatedUser) (user.User, error) {
	var userModel user.User
	ctx := context.Background()
	tx, err := dbClient.Tx(ctx)

	if err != nil {
		return userModel, err
	}

	result, err := createUser(tx, ctx, authUser.User)
	if err != nil {
		return userModel, err
	}

	err = createCredential(tx, ctx, result, authUser)
	if err != nil {
		return userModel, err
	}

	err = createClassroom(tx, ctx, result, authUser)
	if err != nil {
		return userModel, err
	}

	err = tx.Commit()

	if err != nil {
		return userModel, err
	}

	mapToDTO(result.ID, result, &userModel)

	return userModel, nil
}
