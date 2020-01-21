package session

import (
	"github.com/arpb2/C-3PO/src/api/auth"
	"github.com/arpb2/C-3PO/src/api/controller"
	"github.com/arpb2/C-3PO/src/api/http_wrapper"
	"github.com/arpb2/C-3PO/src/api/service"
	"github.com/arpb2/C-3PO/src/api/task/authenticated_user_task"
	"github.com/arpb2/C-3PO/src/api/task/token_task"
	"github.com/arpb2/C-3PO/src/api/validation/authenticated_user_validation"
	"net/http"
)

func CreatePostController(tokenHandler auth.TokenHandler,
	                      service service.CredentialService,
	                      validations []authenticated_user_validation.Validation,
	                      fetchUserTask authenticated_user_task.FetchUserTask,
	                      fetchUserIdTask authenticated_user_task.FetchUserIdTask,
	                      createTokenTask token_task.CreateTokenTask) controller.Controller {
	return controller.Controller{
		Method: "POST",
		Path:   "/session",
		Body:   PostBody{
			TokenHandler:    tokenHandler,
			Service:         service,
			Validations:     validations,
			FetchUserTask:   fetchUserTask,
			FetchUserIdTask: fetchUserIdTask,
			CreateTokenTask: createTokenTask,
		}.Handle,
	}
}

type PostBody struct {
	TokenHandler auth.TokenHandler
	Service      service.CredentialService
	Validations  []authenticated_user_validation.Validation

	FetchUserTask authenticated_user_task.FetchUserTask
	FetchUserIdTask authenticated_user_task.FetchUserIdTask
	CreateTokenTask token_task.CreateTokenTask
}

func (b PostBody) Handle(ctx *http_wrapper.Context) {
	user, err := b.FetchUserTask(ctx)

	if err != nil {
		controller.Halt(ctx, http.StatusBadRequest, err.Error())
		return
	}

	for _, requirement := range b.Validations {
		if err := requirement(user); err != nil {
			controller.Halt(ctx, http.StatusBadRequest, err.Error())
			return
		}
	}

	userId, err := b.FetchUserIdTask(b.Service, user)

	if err != nil {
		controller.Halt(ctx, http.StatusInternalServerError, "internal error")
		return
	}

	token, tokenErr := b.CreateTokenTask(userId, b.TokenHandler)

	if tokenErr != nil {
		controller.Halt(ctx, tokenErr.Status, tokenErr.Error.Error())
		return
	}

	ctx.WriteJson(http.StatusOK, http_wrapper.Json{
		"user_id": userId,
		"token": *token,
	})
}
