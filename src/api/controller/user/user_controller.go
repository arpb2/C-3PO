package user

import (
	"github.com/arpb2/C-3PO/src/api/engine"
)

func Binder(handler engine.ControllerRegistrable) {
	handler.Register(CreateGetController())
	handler.Register(CreatePostController())
	handler.Register(CreatePutController())
	handler.Register(CreateDeleteController())
}