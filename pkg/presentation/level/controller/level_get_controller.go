package controller

import (
	"github.com/arpb2/C-3PO/pkg/domain/controller"
	"github.com/arpb2/C-3PO/pkg/domain/http"
)

func CreateGetController() controller.Controller {
	return controller.Controller{
		Method: "GET",
		Path:   "/levels/:id",
		Body:   get,
	}
}

func get(ctx *http.Context) {
	ctx.WriteString(200, "stub")
}
