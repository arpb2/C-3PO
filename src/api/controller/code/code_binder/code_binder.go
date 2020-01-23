package code_binder

import (
	"github.com/arpb2/C-3PO/src/api/controller/code"
	"github.com/arpb2/C-3PO/src/api/engine"
	"github.com/arpb2/C-3PO/src/api/executor"
	"github.com/arpb2/C-3PO/src/api/http_wrapper"
	"github.com/arpb2/C-3PO/src/api/service"
)

type binder struct {
	Executor executor.HttpExecutor
	AuthMiddleware http_wrapper.Handler
	CodeService service.CodeService
}

func (b binder) BindControllers(controllerRegistrable engine.ControllerRegistrable) {
	controllerRegistrable.Register(code.CreateGetController(b.Executor, b.AuthMiddleware, b.CodeService))
	controllerRegistrable.Register(code.CreatePostController(b.Executor, b.AuthMiddleware, b.CodeService))
	controllerRegistrable.Register(code.CreatePutController(b.Executor, b.AuthMiddleware, b.CodeService))
}

func CreateBinder(exec executor.HttpExecutor, authMiddleware http_wrapper.Handler, codeService service.CodeService) engine.ControllerBinder {
	return &binder{
		Executor:       exec,
		AuthMiddleware: authMiddleware,
		CodeService:    codeService,
	}
}