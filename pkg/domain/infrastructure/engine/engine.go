package engine

import (
	"net/http"

	"github.com/arpb2/C-3PO/pkg/domain/infrastructure/controller"
)

type ServerEngine interface {
	ControllerRegistrable
	http.Handler

	Run() error
}

type ControllerRegistrable interface {
	Register(controller controller.Controller)
}

type ControllerBinder interface {
	BindControllers(controllerRegistrable ControllerRegistrable)
}
