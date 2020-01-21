package user

import (
	"github.com/arpb2/C-3PO/src/api/controller"
	"github.com/arpb2/C-3PO/src/api/http_wrapper"
	"github.com/arpb2/C-3PO/src/api/service"
	"net/http"
)

func CreatePutController(authMiddleware http_wrapper.Handler, userService service.UserService) controller.Controller {
	return controller.Controller{
		Method: "PUT",
		Path:   "/users/:user_id",
		Middleware: []http_wrapper.Handler{
			authMiddleware,
		},
		Body:   CreatePutBody(userService),
	}
}

func CreatePutBody(userService service.UserService) http_wrapper.Handler {
	return func(ctx *http_wrapper.Context) {
		userId, halt := FetchUserId(ctx)
		if halt {
			return
		}

		authenticatedUser, halt := FetchAuthenticatedUser(ctx)
		if halt {
			return
		}

		if authenticatedUser.Id != 0 {
			controller.Halt(ctx, http.StatusBadRequest, "'id' in user is immutable")
			return
		}

		authenticatedUser.Id = userId

		user, err := userService.UpdateUser(authenticatedUser)

		if err != nil || user == nil {
			controller.Halt(ctx, http.StatusInternalServerError, "internal error")
			return
		}

		ctx.JSON(http.StatusOK, user)
	}
}