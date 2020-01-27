package health

import (
	"net/http"

	"github.com/arpb2/C-3PO/api/controller"
	httpwrapper "github.com/arpb2/C-3PO/api/http"
)

func CreateGetController() controller.Controller {
	return controller.Controller{
		Method: "GET",
		Path:   "/ping",
		Body:   healthGet,
	}
}

func healthGet(ctx *httpwrapper.Context) {
	ctx.WriteString(http.StatusOK, "pong")
}
