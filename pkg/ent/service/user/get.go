package user

import (
	"context"

	"github.com/arpb2/C-3PO/api/http"
	"github.com/arpb2/C-3PO/api/model"
	"github.com/arpb2/C-3PO/pkg/ent"
)

func get(dbClient *ent.Client, userId uint) (model.User, error) {
	var userModel model.User
	ctx := context.Background()
	result, err := dbClient.User.Get(ctx, userId)

	if err != nil {
		return userModel, err
	}

	if result == nil {
		return userModel, http.CreateNotFoundError()
	}

	mapToDTO(userId, result, &userModel)

	return userModel, nil
}
