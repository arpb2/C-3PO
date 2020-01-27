package server

import (
	"fmt"

	"github.com/arpb2/C-3PO/api/engine"
	"github.com/arpb2/C-3PO/pkg/auth/jwt"
	codebinder "github.com/arpb2/C-3PO/pkg/binder/code"
	healthbinder "github.com/arpb2/C-3PO/pkg/binder/health"
	sessionbinder "github.com/arpb2/C-3PO/pkg/binder/session"
	userbinder "github.com/arpb2/C-3PO/pkg/binder/user"
	"github.com/arpb2/C-3PO/pkg/executor"
	"github.com/arpb2/C-3PO/pkg/executor/decorator"
	"github.com/arpb2/C-3PO/pkg/middleware/auth/single"
	"github.com/arpb2/C-3PO/pkg/middleware/auth/teacher"
	"github.com/arpb2/C-3PO/pkg/pipeline"
	codeservice "github.com/arpb2/C-3PO/pkg/service/code"
	credentialservice "github.com/arpb2/C-3PO/pkg/service/credential"
	teacherservice "github.com/arpb2/C-3PO/pkg/service/teacher"
	userservice "github.com/arpb2/C-3PO/pkg/service/user"
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
	traceDecorator := decorator.TraceRunnableDecorator

	httpExecutor := executor.CreateHttpExecutor(traceDecorator)
	httpPipeline := pipeline.CreatePipeline(httpExecutor)

	tokenHandler := jwt.CreateTokenHandler()

	userService := userservice.CreateService()
	teacherService := teacherservice.CreateService(userService)
	codeService := codeservice.CreateService()
	credentialService := credentialservice.CreateService()

	singleAuthMiddleware := single.CreateMiddleware(tokenHandler)
	teacherAuthMiddleware := teacher.CreateMiddleware(tokenHandler, teacherService)

	return []engine.ControllerBinder{
		healthbinder.CreateBinder(),
		codebinder.CreateBinder(httpPipeline, teacherAuthMiddleware, codeService),
		userbinder.CreateBinder(httpPipeline, singleAuthMiddleware, userService),
		sessionbinder.CreateBinder(httpPipeline, tokenHandler, credentialService),
	}
}

func RegisterRoutes(engine engine.ServerEngine, binders []engine.ControllerBinder) {
	for _, binder := range binders {
		binder.BindControllers(engine)
	}
}
