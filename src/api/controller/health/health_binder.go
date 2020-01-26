package health

import "github.com/arpb2/C-3PO/src/api/engine"

type binder struct{}

func (b binder) BindControllers(controllerRegistrable engine.ControllerRegistrable) {
	controllerRegistrable.Register(CreateGetController())
}

func CreateBinder() engine.ControllerBinder {
	return &binder{}
}
