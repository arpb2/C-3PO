package user

import (
	"context"

	"github.com/arpb2/C-3PO/pkg/domain/http"
	"github.com/arpb2/C-3PO/pkg/domain/model"
	"github.com/arpb2/C-3PO/third_party/ent"
)

func get(dbClient *ent.Client, userId uint) (model.User, error) {
	var userModel model.User
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
