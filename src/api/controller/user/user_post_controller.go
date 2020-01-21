package user

import (
	"github.com/arpb2/C-3PO/src/api/controller"
	"github.com/arpb2/C-3PO/src/api/http_wrapper"
	"github.com/arpb2/C-3PO/src/api/service"
	"net/http"
)

func CreatePostController(userService service.UserService) controller.Controller {
	return controller.Controller{
		Method: "POST",
		Path:   "/users",
		Body: CreatePostBody(userService),
	}
}

func CreatePostBody(userService service.UserService) http_wrapper.Handler {
	return func(ctx *http_wrapper.Context) {
		authenticatedUser, halt := FetchAuthenticatedUser(ctx)
		if halt {
			return
		}

		user, err := userService.CreateUser(authenticatedUser)

		if err != nil || user == nil {
			controller.Halt(ctx, http.StatusInternalServerError, "internal error")
			return
		}

		ctx.WriteJson(http.StatusOK, user)
	}
}