package health

import (
	"github.com/arpb2/C-3PO/api/engine"
	healthcontroller "github.com/arpb2/C-3PO/pkg/controller/health"
)

type binder struct{}

func (b binder) BindControllers(controllerRegistrable engine.ControllerRegistrable) {
	controllerRegistrable.Register(healthcontroller.CreateGetController())
}

func CreateBinder() engine.ControllerBinder {
	return &binder{}
}
