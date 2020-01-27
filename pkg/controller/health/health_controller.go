package health_controller

import (
	"net/http"

	"github.com/arpb2/C-3PO/api/controller"
	"github.com/arpb2/C-3PO/api/http_wrapper"
)

func CreateGetController() controller.Controller {
	return controller.Controller{
		Method: "GET",
		Path:   "/ping",
		Body:   healthGet,
	}
}

func healthGet(ctx *http_wrapper.Context) {
	ctx.WriteString(http.StatusOK, "pong")
}
