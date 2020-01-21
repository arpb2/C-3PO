package user

import (
	"fmt"
	"github.com/arpb2/C-3PO/src/api/controller"
	"github.com/arpb2/C-3PO/src/api/http_wrapper"
	"github.com/arpb2/C-3PO/src/api/service"
	"net/http"
)

func CreateGetController(authMiddleware http_wrapper.Handler, userService service.UserService) controller.Controller {
	return controller.Controller{
		Method: "GET",
		Path:   "/users/:user_id",
		Middleware: []http_wrapper.Handler{
			authMiddleware,
		},
		Body:   CreateGetBody(userService),
	}
}

func CreateGetBody(userService service.UserService) http_wrapper.Handler {
	return func(ctx *http_wrapper.Context) {
		userId, halt := FetchUserId(ctx)
		if halt {
			return
		}

		user, err := userService.GetUser(userId)

		if err != nil {
			controller.Halt(ctx, http.StatusInternalServerError, "internal error")
			return
		}

		if user == nil {
			controller.Halt(ctx, http.StatusNotFound, fmt.Sprintf("no user associated to the user_id '%d' found", userId))
			return
		}

		ctx.JSON(http.StatusOK, user)
	}
}