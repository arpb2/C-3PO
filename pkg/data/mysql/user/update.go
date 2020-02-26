package user

import (
	"context"
	"time"

	"github.com/arpb2/C-3PO/pkg/data/mysql"
	model2 "github.com/arpb2/C-3PO/pkg/domain/user/model"

	"github.com/arpb2/C-3PO/pkg/domain/infrastructure/http"
	"github.com/arpb2/C-3PO/third_party/ent"
	"github.com/arpb2/C-3PO/third_party/ent/credential"
	"github.com/arpb2/C-3PO/third_party/ent/user"
)

func updateUser(tx *ent.Tx, ctx context.Context, user model2.User) (*ent.User, error) {
	createOp := tx.User.Create()
	if len(user.Name) > 0 {
		createOp = createOp.SetName(user.Name)
	}
	if len(user.Surname) > 0 {
		createOp = createOp.SetSurname(user.Surname)
	}
	if len(user.Email) > 0 {
		createOp = createOp.SetEmail(user.Email)
	}

	result, err := createOp.
		SetUpdatedAt(time.Now()).
		Save(ctx)
	if err != nil || result == nil {
		return nil, mysql.Rollback(tx, err)
	}
	return result, nil
}

func hasToUpdateCredential(dbClient *ent.Client, ctx context.Context, authUser model2.AuthenticatedUser) (*[]byte, error) {
	if len(authUser.Password) > 0 {
		pwHash, err := dbClient.Credential.
			Query().
			Where(credential.HasHolderWith(user.ID(authUser.Id))).
			Select(credential.FieldSalt).
			Strings(ctx)

		if err != nil {
			return nil, err
		}

		if len(pwHash) != 1 {
			return nil, http.CreateInternalError()
		}

		newPwHash, err := mysql.SaltHash([]byte(authUser.Password), []byte(pwHash[0]))

		if err != nil {
			return nil, err
		}

		return &newPwHash, nil
	}
	return nil, nil
}

func updateCredential(tx *ent.Tx, ctx context.Context, userId uint, hash []byte) error {
	matches, err := tx.Credential.Update().
		Where(credential.HasHolderWith(user.ID(userId))).
		SetPasswordHash(hash).
		Save(ctx)

	if err != nil || matches != 1 {
		return mysql.Rollback(tx, err)
	}
	return nil
}

func update(dbClient *ent.Client, authUser model2.AuthenticatedUser) (model2.User, error) {
	var userModel model2.User
	ctx := context.Background()
	tx, err := dbClient.Tx(ctx)

	if err != nil {
		return userModel, err
	}

	hashPw, err := hasToUpdateCredential(dbClient, ctx, authUser)
	if err != nil {
		return userModel, http.CreateInternalError()
	}

	if hashPw != nil {
		err = updateCredential(tx, ctx, authUser.Id, *hashPw)

		if err != nil {
			return userModel, err
		}
	}

	result, err := updateUser(tx, ctx, authUser.User)

	if err = tx.Commit(); err != nil {
		return userModel, err
	}

	mapToDTO(result.ID, result, &userModel)

	return userModel, nil
}
