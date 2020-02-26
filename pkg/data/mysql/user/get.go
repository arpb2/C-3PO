package user

import (
	"context"

	model2 "github.com/arpb2/C-3PO/pkg/domain/user/model"

	"github.com/arpb2/C-3PO/pkg/domain/infrastructure/http"
	"github.com/arpb2/C-3PO/third_party/ent"
)

func get(dbClient *ent.Client, userId uint) (model2.User, error) {
	var userModel model2.User
	ctx := context.Background()
	result, err := dbClient.User.Get(ctx, userId)

	if err != nil {
		if ent.IsNotFound(err) {
			return userModel, http.CreateNotFoundError()
		}
		return userModel, err
	}

	if result == nil {
		return userModel, http.CreateNotFoundError()
	}

	mapToDTO(userId, result, &userModel)

	return userModel, nil
}
