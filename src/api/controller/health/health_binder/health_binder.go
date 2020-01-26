package health_binder

import (
	"github.com/arpb2/C-3PO/src/api/controller/health"
	"github.com/arpb2/C-3PO/src/api/engine"
)

type binder struct{}

func (b binder) BindControllers(controllerRegistrable engine.ControllerRegistrable) {
	controllerRegistrable.Register(health.CreateGetController())
}

func CreateBinder() engine.ControllerBinder {
	return &binder{}
}
