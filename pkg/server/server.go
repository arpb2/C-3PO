package server

import (
	"fmt"
	"os"

	"github.com/arpb2/C-3PO/pkg/ent/client"

	"github.com/arpb2/C-3PO/api/engine"
	"github.com/arpb2/C-3PO/pkg/auth/jwt"
	healthbinder "github.com/arpb2/C-3PO/pkg/binder/health"
	sessionbinder "github.com/arpb2/C-3PO/pkg/binder/session"
	userbinder "github.com/arpb2/C-3PO/pkg/binder/user"
	userlevelbinder "github.com/arpb2/C-3PO/pkg/binder/user_level"
	credentialservice "github.com/arpb2/C-3PO/pkg/ent/service/credential"
	teacherservice "github.com/arpb2/C-3PO/pkg/ent/service/teacher"
	userservice "github.com/arpb2/C-3PO/pkg/ent/service/user"
	userlevelservice "github.com/arpb2/C-3PO/pkg/ent/service/user_level"
	"github.com/arpb2/C-3PO/pkg/executor"
	"github.com/arpb2/C-3PO/pkg/executor/decorator"
	"github.com/arpb2/C-3PO/pkg/middleware/auth/single"
	"github.com/arpb2/C-3PO/pkg/middleware/auth/teacher"
	"github.com/arpb2/C-3PO/pkg/pipeline"
)

func StartApplication(engine engine.ServerEngine) error {
	binders, deferredClose := CreateBinders()
	defer deferredClose()

	RegisterRoutes(engine, binders)

	if err := engine.Run(); err != nil {
		_ = fmt.Errorf("error running server %s", err.Error())
		return err
	}
	return nil
}

func CreateBinders() ([]engine.ControllerBinder, func()) {
	traceDecorator := decorator.CreateTraceDecorator(os.Stdout)

	dbClient, driver := client.CreateMysqlClient(os.Getenv("MYSQL_DSN"))

	httpExecutor := executor.CreateHttpExecutor(traceDecorator)
	httpPipeline := pipeline.CreateHttpPipeline(httpExecutor)

	tokenHandler := jwt.CreateTokenHandler()

	userService := userservice.CreateService(dbClient)
	teacherService := teacherservice.CreateService(userService, dbClient)
	userLevelService := userlevelservice.CreateService(dbClient)
	credentialService := credentialservice.CreateService(dbClient)

	singleAuthMiddleware := single.CreateMiddleware(tokenHandler)
	teacherAuthMiddleware := teacher.CreateMiddleware(tokenHandler, teacherService)

	return []engine.ControllerBinder{
			healthbinder.CreateBinder(),
			userlevelbinder.CreateBinder(httpPipeline, teacherAuthMiddleware, userLevelService),
			userbinder.CreateBinder(httpPipeline, singleAuthMiddleware, userService),
			sessionbinder.CreateBinder(httpPipeline, tokenHandler, credentialService),
		}, func() {
			_ = driver.Close()
		}
}

func RegisterRoutes(engine engine.ServerEngine, binders []engine.ControllerBinder) {
	for _, binder := range binders {
		binder.BindControllers(engine)
	}
}
