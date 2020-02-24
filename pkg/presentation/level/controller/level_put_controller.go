package controller

import (
	"github.com/arpb2/C-3PO/pkg/domain/controller"
	"github.com/arpb2/C-3PO/pkg/domain/http"
)

func CreatePutController(authMiddleware http.Handler) controller.Controller {
	return controller.Controller{
		Method: "PUT",
		Path:   "/levels/:id",
		Middleware: []http.Handler{
			authMiddleware,
		},
		Body: put,
	}
}

func put(ctx *http.Context) {
	ctx.WriteString(200, "stub")
}
