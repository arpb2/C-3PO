package engine

import (
	"net/http"

	"github.com/arpb2/C-3PO/api/controller"
)

type ServerEngine interface {
	ControllerRegistrable
	http.Handler

	Run() error

	Shutdown() error
}

type ControllerRegistrable interface {
	Register(controller controller.Controller)
}

type ControllerBinder interface {
	BindControllers(controllerRegistrable ControllerRegistrable)
}
