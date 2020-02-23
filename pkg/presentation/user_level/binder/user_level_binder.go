package binder

import (
	"github.com/arpb2/C-3PO/pkg/domain/engine"
	"github.com/arpb2/C-3PO/pkg/domain/http"
	"github.com/arpb2/C-3PO/pkg/domain/pipeline"
	userlevelservice "github.com/arpb2/C-3PO/pkg/domain/service/user_level"
	"github.com/arpb2/C-3PO/pkg/presentation/user_level/controller"
)

type binder struct {
	Executor         pipeline.HttpPipeline
	AuthMiddleware   http.Handler
	UserLevelService userlevelservice.Service
}

func (b binder) BindControllers(controllerRegistrable engine.ControllerRegistrable) {
	controllerRegistrable.Register(controller.CreateGetController(b.Executor, b.AuthMiddleware, b.UserLevelService))
	controllerRegistrable.Register(controller.CreatePutController(b.Executor, b.AuthMiddleware, b.UserLevelService))
}

func CreateBinder(exec pipeline.HttpPipeline, authMiddleware http.Handler, userLevelService userlevelservice.Service) engine.ControllerBinder {
	return &binder{
		Executor:         exec,
		AuthMiddleware:   authMiddleware,
		UserLevelService: userLevelService,
	}
}
