package health

import (
	"github.com/arpb2/C-3PO/src/api/controller"
	"github.com/arpb2/C-3PO/src/api/http_wrapper"
	"net/http"
)

var GetController = controller.Controller{
	Method: "GET",
	Path:   "/ping",
	Body:   healthGet,
}

func healthGet(ctx *http_wrapper.Context) {
	ctx.WriteString(http.StatusOK, "pong")
}
