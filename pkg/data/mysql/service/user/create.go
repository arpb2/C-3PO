package user

import (
	"context"
	"crypto/rand"
	"io"

	"github.com/arpb2/C-3PO/pkg/data/mysql/service"
	"github.com/arpb2/C-3PO/pkg/domain/model"
	"github.com/arpb2/C-3PO/third_party/ent"
)

const (
	saltBytes = 32
)

func createUser(tx *ent.Tx, ctx context.Context, user model.User) (*ent.User, error) {
	result, err := tx.User.Create().
		SetName(user.Name).
		SetSurname(user.Surname).
		SetEmail(user.Email).
		Save(ctx)

	if err != nil || result == nil {
		return nil, service.Rollback(tx, err)
	}
	return result, nil
}

func createCredential(tx *ent.Tx, ctx context.Context, holder *ent.User, user model.AuthenticatedUser) error {
	salt := make([]byte, saltBytes)
	_, err := io.ReadFull(rand.Reader, salt)
	if err != nil {
		return service.Rollback(tx, err)
	}

	hash, err := service.SaltHash([]byte(user.Password), salt)
	if err != nil {
		return service.Rollback(tx, err)
	}

	credential, err := tx.Credential.Create().
		SetHolder(holder).
		SetSalt(salt).
		SetPasswordHash(hash).
		Save(ctx)

	if err != nil || credential == nil {
		return service.Rollback(tx, err)
	}
	return nil
}

func create(dbClient *ent.Client, user model.AuthenticatedUser) (model.User, error) {
	var userModel model.User
	ctx := context.Background()
	tx, err := dbClient.Tx(ctx)

	if err != nil {
		return userModel, err
	}

	result, err := createUser(tx, ctx, user.User)
	if err != nil {
		return userModel, err
	}

	err = createCredential(tx, ctx, result, user)
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
