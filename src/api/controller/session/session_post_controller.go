package session

import (
	"github.com/arpb2/C-3PO/src/api/auth"
	"github.com/arpb2/C-3PO/src/api/controller"
	"github.com/arpb2/C-3PO/src/api/controller/session/session_task"
	"github.com/arpb2/C-3PO/src/api/controller/session/session_validation"
	"github.com/arpb2/C-3PO/src/api/http_wrapper"
	"github.com/arpb2/C-3PO/src/api/model"
	"github.com/arpb2/C-3PO/src/api/service"
	"net/http"
)

func CreatePostController(tokenHandler auth.TokenHandler,
	                      service service.CredentialService,
	                      validations []session_validation.Validation,
	                      fetchUserTask session_task.FetchUserTask) controller.Controller {
	return controller.Controller{
		Method: "POST",
		Path:   "/session",
		Body:   PostBody{
			TokenHandler:    tokenHandler,
			Service:         service,
			Validations:     validations,
			FetchUserTask:   fetchUserTask,
		}.Post,
	}
}

type PostBody struct {
	TokenHandler auth.TokenHandler
	Service      service.CredentialService

	Validations  []session_validation.Validation

	FetchUserTask func(ctx *http_wrapper.Context) (user *model.AuthenticatedUser, err error)
}

func (b PostBody) getUserData(ctx *http_wrapper.Context) (email, password string, ok bool){
	user, err := b.FetchUserTask(ctx)

	if err != nil {
		controller.Halt(ctx, http.StatusBadRequest, err.Error())
		ok = false
		return
	}

	for _, requirement := range b.Validations {
		if err := requirement(user); err != nil {
			controller.Halt(ctx, http.StatusBadRequest, err.Error())
			ok = false
			return
		}
	}

	email = user.Email
	password = user.Password
	ok = true
	return
}

func (b PostBody) authenticate(ctx *http_wrapper.Context, email, password string) (token *string, userId uint, ok bool) {
	userId, err := b.Service.Retrieve(email, password)

	if err != nil {
		controller.Halt(ctx, http.StatusInternalServerError, "internal error")
		ok = false
		return
	}

	token, tokenErr := b.TokenHandler.Create(auth.Token{
		UserId: userId,
	})

	if tokenErr != nil {
		controller.Halt(ctx, tokenErr.Status, tokenErr.Error.Error())
		ok = false
		return
	}

	ok = true
	return
}

func (b PostBody) Post(ctx *http_wrapper.Context) {
	if email, password, ok := b.getUserData(ctx); ok {
		if token, userId, ok := b.authenticate(ctx, email, password); ok {
			ctx.WriteJson(http.StatusOK, http_wrapper.Json{
				"user_id": userId,
				"token": *token,
			})
		}
	}
}
