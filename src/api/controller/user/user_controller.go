package user

import (
	"github.com/arpb2/C-3PO/src/api/engine"
)

func Binder(handler engine.ControllerRegistrable) {
	handler.Register(GetController)
	handler.Register(PostController)
	handler.Register(PutController)
	handler.Register(DeleteController)
}