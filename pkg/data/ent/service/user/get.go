package user

import (
	"context"

	"github.com/arpb2/C-3PO/pkg/data/ent"
	"github.com/arpb2/C-3PO/pkg/domain/http"
	"github.com/arpb2/C-3PO/pkg/domain/model"
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
