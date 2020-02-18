package user

import (
	"context"
	"time"

	"github.com/arpb2/C-3PO/api/http"
	"github.com/arpb2/C-3PO/api/model"
	"github.com/arpb2/C-3PO/pkg/ent"
	"github.com/arpb2/C-3PO/pkg/ent/credential"
	"github.com/arpb2/C-3PO/pkg/ent/service"
	"github.com/arpb2/C-3PO/pkg/ent/user"
)

func updateUser(tx *ent.Tx, ctx context.Context, user model.User) (*ent.User, error) {
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
		return nil, service.Rollback(tx, err)
	}
	return result, nil
}

func hasToUpdateCredential(dbClient *ent.Client, ctx context.Context, authUser model.AuthenticatedUser) (*[]byte, error) {
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

		newPwHash, err := service.SaltHash([]byte(authUser.Password), []byte(pwHash[0]))

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
		return service.Rollback(tx, err)
	}
	return nil
}

func update(dbClient *ent.Client, authUser model.AuthenticatedUser) (model.User, error) {
	var userModel model.User
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
