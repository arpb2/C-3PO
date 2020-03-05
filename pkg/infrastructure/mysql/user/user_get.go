package user

import (
	"context"

	"github.com/arpb2/C-3PO/pkg/domain/model/user"

	"github.com/arpb2/C-3PO/pkg/domain/http"
	"github.com/arpb2/C-3PO/third_party/ent"
)

func get(dbClient *ent.Client, userId uint) (user.User, error) {
	var userModel user.User
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
