package server

import (
	"fmt"
	"github.com/arpb2/C-3PO/src/api/auth/jwt"
	"github.com/arpb2/C-3PO/src/api/controller/code/code_binder"
	"github.com/arpb2/C-3PO/src/api/controller/health/health_binder"
	"github.com/arpb2/C-3PO/src/api/controller/session/session_binder"
	"github.com/arpb2/C-3PO/src/api/controller/user/user_binder"
	"github.com/arpb2/C-3PO/src/api/engine"
	"github.com/arpb2/C-3PO/src/api/executor"
	"github.com/arpb2/C-3PO/src/api/middleware/auth/single_auth"
	"github.com/arpb2/C-3PO/src/api/middleware/auth/teacher_auth"
	"github.com/arpb2/C-3PO/src/api/service/code_service"
	"github.com/arpb2/C-3PO/src/api/service/credential_service"
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
	httpExecutor := executor.CreateHttpExecutor()
	httpPipeline := executor.CreatePipeline(httpExecutor)

	tokenHandler := jwt.CreateTokenHandler()

	userService := user_service.CreateService()
	teacherService := teacher_service.CreateService(userService)
	codeService := code_service.CreateService()
	credentialService := credential_service.CreateService()

	singleAuthMiddleware := single_auth.CreateMiddleware(tokenHandler)
	teacherAuthMiddleware := teacher_auth.CreateMiddleware(tokenHandler, teacherService)

	return []engine.ControllerBinder{
		health_binder.CreateBinder(),
		code_binder.CreateBinder(httpPipeline, teacherAuthMiddleware, codeService),
		user_binder.CreateBinder(httpPipeline, singleAuthMiddleware, userService),
		session_binder.CreateBinder(httpPipeline, tokenHandler, credentialService),
	}
}

func RegisterRoutes(engine engine.ServerEngine, binders []engine.ControllerBinder) {
	for _, binder := range binders {
		binder.BindControllers(engine)
	}
}
