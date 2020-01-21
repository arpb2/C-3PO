package user

import (
	"github.com/arpb2/C-3PO/src/api/http_wrapper"
	"github.com/arpb2/C-3PO/src/api/controller"
	"net/http"
)

func CreatePutController(authMiddleware http_wrapper.Handler) controller.Controller {
	return controller.Controller{
		Method: "PUT",
		Path:   "/users/:user_id",
		Middleware: []http_wrapper.Handler{
			authMiddleware,
		},
		Body:   userPut,
	}
}

func userPut(ctx *http_wrapper.Context) {
	ctx.JSON(http.StatusOK, http_wrapper.Json{})
}
