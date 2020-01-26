package health_binder

import (
	"github.com/arpb2/C-3PO/api/engine"
	"github.com/arpb2/C-3PO/internal/controller/health"
)

type binder struct{}

func (b binder) BindControllers(controllerRegistrable engine.ControllerRegistrable) {
	controllerRegistrable.Register(health_controller.CreateGetController())
}

func CreateBinder() engine.ControllerBinder {
	return &binder{}
}