package health

import (
	"github.com/arpb2/C-3PO/src/api/controller"
	"github.com/arpb2/C-3PO/src/api/engine"
	"github.com/arpb2/C-3PO/src/api/http_wrapper"
	"net/http"
)

func Binder(handler engine.ControllerRegistrable) {
	handler.Register(GetController)
}

var GetController = controller.Controller{
	Method: "GET",
	Path:   "/ping",
	Body:   healthGet,
}

func healthGet(ctx *http_wrapper.Context) {
	ctx.String(http.StatusOK, "pong")
}
