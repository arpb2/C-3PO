package health_controller

import (
	controller2 "github.com/arpb2/C-3PO/api/controller"
	"github.com/arpb2/C-3PO/api/http_wrapper"
	"net/http"
)

func CreateGetController() controller2.Controller {
	return controller2.Controller{
		Method: "GET",
		Path:   "/ping",
		Body:   healthGet,
	}
}

func healthGet(ctx *http_wrapper.Context) {
	ctx.WriteString(http.StatusOK, "pong")
}
