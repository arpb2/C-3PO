package code

import (
	"github.com/arpb2/C-3PO/api/engine"
	"github.com/arpb2/C-3PO/api/http"
	"github.com/arpb2/C-3PO/api/pipeline"
	codeservice "github.com/arpb2/C-3PO/api/service/code"
	codecontroller "github.com/arpb2/C-3PO/pkg/controller/code"
)

type binder struct {
	Executor       pipeline.HttpPipeline
	AuthMiddleware http.Handler
	CodeService    codeservice.Service
}

func (b binder) BindControllers(controllerRegistrable engine.ControllerRegistrable) {
	controllerRegistrable.Register(codecontroller.CreateGetController(b.Executor, b.AuthMiddleware, b.CodeService))
	controllerRegistrable.Register(codecontroller.CreatePostController(b.Executor, b.AuthMiddleware, b.CodeService))
	controllerRegistrable.Register(codecontroller.CreatePutController(b.Executor, b.AuthMiddleware, b.CodeService))
}

func CreateBinder(exec pipeline.HttpPipeline, authMiddleware http.Handler, codeService codeservice.Service) engine.ControllerBinder {
	return &binder{
		Executor:       exec,
		AuthMiddleware: authMiddleware,
		CodeService:    codeService,
	}
}
