package server

import (
	"fmt"
	"github.com/arpb2/C-3PO/src/api/auth/jwt"
	"github.com/arpb2/C-3PO/src/api/controller/code"
	"github.com/arpb2/C-3PO/src/api/controller/health"
	"github.com/arpb2/C-3PO/src/api/controller/session"
	"github.com/arpb2/C-3PO/src/api/controller/user"
	"github.com/arpb2/C-3PO/src/api/engine"
	"github.com/arpb2/C-3PO/src/api/middleware/auth/single_auth"
	"github.com/arpb2/C-3PO/src/api/middleware/auth/teacher_auth"
	"github.com/arpb2/C-3PO/src/api/service/code_service"
	"github.com/arpb2/C-3PO/src/api/service/teacher_service"
	"github.com/arpb2/C-3PO/src/api/service/user_service"
)

func StartApplication(engine engine.ServerEngine) error {
	RegisterRoutes(engine, CreateBinders())

	if err := engine.Run(); err != nil {
		_ = fmt.Errorf("error running server %s", err.Error())
		return err
	}
	return nil
}

func CreateBinders() []engine.ControllerBinder {
	tokenHandler := jwt.CreateTokenHandler()

	userService := user_service.CreateService()
	teacherService := teacher_service.CreateService(userService)
	codeService := code_service.CreateService()

	singleAuthMiddleware := single_auth.CreateMiddleware(tokenHandler)
	teacherAuthMiddleware := teacher_auth.CreateMiddleware(tokenHandler, teacherService)

	return []engine.ControllerBinder{
		health.CreateBinder(),
		code.CreateBinder(teacherAuthMiddleware, codeService),
		user.CreateBinder(singleAuthMiddleware, userService),
		session.CreateBinder(),
	}
}

func RegisterRoutes(engine engine.ServerEngine, binders []engine.ControllerBinder) {
	for _, binder := range binders {
		binder.BindControllers(engine)
	}
}
