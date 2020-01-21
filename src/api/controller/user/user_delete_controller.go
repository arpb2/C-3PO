package user

import (
	"github.com/arpb2/C-3PO/src/api/controller"
	"github.com/arpb2/C-3PO/src/api/http_wrapper"
	"github.com/arpb2/C-3PO/src/api/service"
	"net/http"
)

func CreateDeleteController(authMiddleware http_wrapper.Handler, userService service.UserService) controller.Controller {
	return controller.Controller{
		Method: "DELETE",
		Path:   "/users/:user_id",
		Middleware: []http_wrapper.Handler{
			authMiddleware,
		},
		Body:   CreateDeleteBody(userService),
	}
}

func CreateDeleteBody(userService service.UserService) http_wrapper.Handler {
	return func(ctx *http_wrapper.Context) {
		userId, halt := FetchUserId(ctx)
		if halt {
			return
		}

		err := userService.DeleteUser(userId)

		if err != nil {
			controller.Halt(ctx, http.StatusInternalServerError, "internal error")
			return
		}

		ctx.WriteStatus(http.StatusOK)
	}
}