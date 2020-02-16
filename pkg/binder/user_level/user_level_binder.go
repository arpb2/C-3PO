package user_level

import (
	"github.com/arpb2/C-3PO/api/engine"
	"github.com/arpb2/C-3PO/api/http"
	"github.com/arpb2/C-3PO/api/pipeline"
	userlevelservice "github.com/arpb2/C-3PO/api/service/user_level"
	userlevelcontroller "github.com/arpb2/C-3PO/pkg/controller/user_level"
)

type binder struct {
	Executor         pipeline.HttpPipeline
	AuthMiddleware   http.Handler
	UserLevelService userlevelservice.Service
}

func (b binder) BindControllers(controllerRegistrable engine.ControllerRegistrable) {
	controllerRegistrable.Register(userlevelcontroller.CreateGetController(b.Executor, b.AuthMiddleware, b.UserLevelService))
	controllerRegistrable.Register(userlevelcontroller.CreatePutController(b.Executor, b.AuthMiddleware, b.UserLevelService))
}

func CreateBinder(exec pipeline.HttpPipeline, authMiddleware http.Handler, userLevelService userlevelservice.Service) engine.ControllerBinder {
	return &binder{
		Executor:         exec,
		AuthMiddleware:   authMiddleware,
		UserLevelService: userLevelService,
	}
}
