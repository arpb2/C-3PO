package code_binder

import (
	"github.com/arpb2/C-3PO/src/api/controller/code"
	"github.com/arpb2/C-3PO/src/api/engine"
	"github.com/arpb2/C-3PO/src/api/http_wrapper"
	"github.com/arpb2/C-3PO/src/api/service"
)

type binder struct {
	AuthMiddleware http_wrapper.Handler
	CodeService service.CodeService
}

func (b binder) BindControllers(controllerRegistrable engine.ControllerRegistrable) {
	controllerRegistrable.Register(code.CreateGetController(b.AuthMiddleware, b.CodeService))
	controllerRegistrable.Register(code.CreatePostController(b.AuthMiddleware, b.CodeService))
	controllerRegistrable.Register(code.CreatePutController(b.AuthMiddleware, b.CodeService))
}

func CreateBinder(authMiddleware http_wrapper.Handler, codeService service.CodeService) engine.ControllerBinder {
	return &binder{
		AuthMiddleware: authMiddleware,
		CodeService:    codeService,
	}
}