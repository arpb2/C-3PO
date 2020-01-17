package engine

import (
	"github.com/arpb2/C-3PO/src/api/controller"
	"net/http"
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