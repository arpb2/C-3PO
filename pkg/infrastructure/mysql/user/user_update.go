package user

import (
	"context"
	"time"

	user2 "github.com/arpb2/C-3PO/pkg/domain/model/user"

	"github.com/arpb2/C-3PO/pkg/infrastructure/mysql"

	"github.com/arpb2/C-3PO/pkg/domain/http"
	"github.com/arpb2/C-3PO/third_party/ent"
	"github.com/arpb2/C-3PO/third_party/ent/credential"
	"github.com/arpb2/C-3PO/third_party/ent/user"
)

func updateUser(tx *ent.Tx, ctx context.Context, user user2.User) (*ent.User, error) {
	updateOp := tx.User.UpdateOneID(user.Id)
	if len(user.Name) > 0 {
		updateOp = updateOp.SetName(user.Name)
	}
	if len(user.Surname) > 0 {
		updateOp = updateOp.SetSurname(user.Surname)
	}
	if len(user.Email) > 0 {
		updateOp = updateOp.SetEmail(user.Email)
	}

	result, err := updateOp.
		SetUpdatedAt(time.Now()).
		Save(ctx)

	if err != nil {
		if ent.IsNotFound(err) {
			return nil, http.CreateNotFoundError()
		}
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

func hasToUpdateCredential(dbClient *ent.Client, ctx context.Context, authUser user2.AuthenticatedUser) (*[]byte, error) {
	if len(authUser.Password) > 0 {
		pwHash, err := dbClient.Credential.
			Query().
			Where(credential.HasHolderWith(user.ID(authUser.Id))).
			Select(credential.FieldSalt).
			Strings(ctx)

		if err != nil {
			if ent.IsNotFound(err) {
				return nil, http.CreateNotFoundError()
			}
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

	if err != nil {
		if ent.IsNotFound(err) {
			return http.CreateNotFoundError()
		}
		if ent.IsConstraintError(err) {
			return http.CreateBadRequestError("constraint error")
		}
		return err
	}

	if matches != 1 {
		return http.CreateInternalError()
	}

	return nil
}

func update(dbClient *ent.Client, authUser user2.AuthenticatedUser) (user2.User, error) {
	var userModel user2.User
	ctx := context.Background()
	tx, err := dbClient.Tx(ctx)

	if err != nil {
		return userModel, mysql.Rollback(tx, err)
	}

	hashPw, err := hasToUpdateCredential(dbClient, ctx, authUser)
	if err != nil {
		return userModel, mysql.Rollback(tx, err)
	}

	if hashPw != nil {
		err = updateCredential(tx, ctx, authUser.Id, *hashPw)

		if err != nil {
			return userModel, mysql.Rollback(tx, err)
		}
	}

	result, err := updateUser(tx, ctx, authUser.User)
	if err != nil {
		return userModel, mysql.Rollback(tx, err)
	}

	if err = tx.Commit(); err != nil {
		return userModel, err
	}

	if err = mapToDTO(dbClient, result.ID, result, &userModel); err != nil {
		return userModel, err
	}

	return userModel, nil
}
