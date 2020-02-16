package server

import (
	"fmt"
	"os"

	"github.com/arpb2/C-3PO/api/engine"
	"github.com/arpb2/C-3PO/pkg/auth/jwt"
	healthbinder "github.com/arpb2/C-3PO/pkg/binder/health"
	sessionbinder "github.com/arpb2/C-3PO/pkg/binder/session"
	userbinder "github.com/arpb2/C-3PO/pkg/binder/user"
	userlevelbinder "github.com/arpb2/C-3PO/pkg/binder/user_level"
	"github.com/arpb2/C-3PO/pkg/executor"
	"github.com/arpb2/C-3PO/pkg/executor/decorator"
	"github.com/arpb2/C-3PO/pkg/middleware/auth/single"
	"github.com/arpb2/C-3PO/pkg/middleware/auth/teacher"
	"github.com/arpb2/C-3PO/pkg/pipeline"
	credentialservice "github.com/arpb2/C-3PO/pkg/service/credential"
	teacherservice "github.com/arpb2/C-3PO/pkg/service/teacher"
	userservice "github.com/arpb2/C-3PO/pkg/service/user"
	userlevelservice "github.com/arpb2/C-3PO/pkg/service/user_level"
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
	traceDecorator := decorator.CreateTraceDecorator(os.Stdout)

	httpExecutor := executor.CreateHttpExecutor(traceDecorator)
	httpPipeline := pipeline.CreateHttpPipeline(httpExecutor)

	tokenHandler := jwt.CreateTokenHandler()

	userService := userservice.CreateService()
	teacherService := teacherservice.CreateService(userService)
	userLevelService := userlevelservice.CreateService()
	credentialService := credentialservice.CreateService()

	singleAuthMiddleware := single.CreateMiddleware(tokenHandler)
	teacherAuthMiddleware := teacher.CreateMiddleware(tokenHandler, teacherService)

	return []engine.ControllerBinder{
		healthbinder.CreateBinder(),
		userlevelbinder.CreateBinder(httpPipeline, teacherAuthMiddleware, userLevelService),
		userbinder.CreateBinder(httpPipeline, singleAuthMiddleware, userService),
		sessionbinder.CreateBinder(httpPipeline, tokenHandler, credentialService),
	}
}

func RegisterRoutes(engine engine.ServerEngine, binders []engine.ControllerBinder) {
	for _, binder := range binders {
		binder.BindControllers(engine)
	}
}
