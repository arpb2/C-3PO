package user

import (
	"github.com/arpb2/C-3PO/src/api/controller"
	"github.com/arpb2/C-3PO/src/api/http_wrapper"
	"github.com/arpb2/C-3PO/src/api/service"
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

	}
}