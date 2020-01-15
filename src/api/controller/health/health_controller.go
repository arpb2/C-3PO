package health

import (
	"github.com/arpb2/C-3PO/src/api/controller"
	"github.com/arpb2/C-3PO/src/api/engine"
	"github.com/gin-gonic/gin"
	"net/http"
)

func Binder(handler engine.ControllerHandler) {
	handler.Register(GetController)
}

var GetController = controller.Controller{
	Method: "GET",
	Path:   "/ping",
	Body:   healthGet,
}

func healthGet(ctx *gin.Context) {
	ctx.String(http.StatusOK, "pong")
}
