package code_binder

import (
	"github.com/arpb2/C-3PO/api/engine"
	"github.com/arpb2/C-3PO/api/http_wrapper"
	"github.com/arpb2/C-3PO/api/pipeline"
	code_service "github.com/arpb2/C-3PO/api/service/code"
	code_controller "github.com/arpb2/C-3PO/internal/controller/code"
)

type binder struct {
	Executor       pipeline.HttpPipeline
	AuthMiddleware http_wrapper.Handler
	CodeService    code_service.Service
}

func (b binder) BindControllers(controllerRegistrable engine.ControllerRegistrable) {
	controllerRegistrable.Register(code_controller.CreateGetController(b.Executor, b.AuthMiddleware, b.CodeService))
	controllerRegistrable.Register(code_controller.CreatePostController(b.Executor, b.AuthMiddleware, b.CodeService))
	controllerRegistrable.Register(code_controller.CreatePutController(b.Executor, b.AuthMiddleware, b.CodeService))
}

func CreateBinder(exec pipeline.HttpPipeline, authMiddleware http_wrapper.Handler, codeService code_service.Service) engine.ControllerBinder {
	return &binder{
		Executor:       exec,
		AuthMiddleware: authMiddleware,
		CodeService:    codeService,
	}
}
