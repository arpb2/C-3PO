package binder

import (
	"github.com/arpb2/C-3PO/pkg/domain/engine"
	"github.com/arpb2/C-3PO/pkg/presentation/health/controller"
)

type binder struct{}

func (b binder) BindControllers(controllerRegistrable engine.ControllerRegistrable) {
	controllerRegistrable.Register(controller.CreateGetController())
}

func CreateBinder() engine.ControllerBinder {
	return &binder{}
}
