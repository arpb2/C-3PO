package session

import "github.com/arpb2/C-3PO/src/api/engine"

type binder struct{}

func (b binder) BindControllers(controllerRegistrable engine.ControllerRegistrable) {
	controllerRegistrable.Register(CreatePostController())
}

func CreateBinder() engine.ControllerBinder {
	return &binder{}
}
