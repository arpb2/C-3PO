package user

import (
	"github.com/arpb2/C-3PO/src/api/http_wrapper"
	"github.com/arpb2/C-3PO/src/api/controller"
	"net/http"
)

func CreateGetController(authMiddleware http_wrapper.Handler) controller.Controller {
	return controller.Controller{
		Method: "GET",
		Path:   "/users/:user_id",
		Middleware: []http_wrapper.Handler{
			authMiddleware,
		},
		Body:   userGet,
	}
}

func userGet(ctx *http_wrapper.Context) {
	ctx.JSON(http.StatusOK, http_wrapper.Json{})
}