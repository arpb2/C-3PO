package authenticated_user_task

import (
	"errors"
	"github.com/arpb2/C-3PO/src/api/http_wrapper"
	"github.com/arpb2/C-3PO/src/api/model"
)

type FetchUserTask func(ctx *http_wrapper.Context) (user *model.AuthenticatedUser, err error)

func FetchUserTaskImpl(ctx *http_wrapper.Context) (user *model.AuthenticatedUser, err error) {
	var authenticatedUser model.AuthenticatedUser

	if err := ctx.ReadBody(&authenticatedUser); err != nil {
		return nil, errors.New("malformed body")
	}

	return &authenticatedUser, nil
}
