package engine

import (
	controller2 "github.com/arpb2/C-3PO/api/controller"
	"net/http"
)

type ServerEngine interface {
	ControllerRegistrable
	http.Handler

	Run() error

	Shutdown() error
}

type ControllerRegistrable interface {
	Register(controller controller2.Controller)
}

type ControllerBinder interface {
	BindControllers(controllerRegistrable ControllerRegistrable)
}
