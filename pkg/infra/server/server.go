package server

import (
	"fmt"
	"github.com/arpb2/C-3PO/pkg/presentation/health/binder"
	binder2 "github.com/arpb2/C-3PO/pkg/presentation/session/binder"
	binder3 "github.com/arpb2/C-3PO/pkg/presentation/user/binder"
	binder4 "github.com/arpb2/C-3PO/pkg/presentation/user_level/binder"
	"os"

	"github.com/arpb2/C-3PO/pkg/data/ent/client"

	credentialservice "github.com/arpb2/C-3PO/pkg/data/ent/service/credential"
	teacherservice "github.com/arpb2/C-3PO/pkg/data/ent/service/teacher"
	userservice "github.com/arpb2/C-3PO/pkg/data/ent/service/user"
	userlevelservice "github.com/arpb2/C-3PO/pkg/data/ent/service/user_level"
	"github.com/arpb2/C-3PO/pkg/data/jwt"
	"github.com/arpb2/C-3PO/pkg/domain/engine"
	"github.com/arpb2/C-3PO/pkg/infra/executor"
	"github.com/arpb2/C-3PO/pkg/infra/executor/decorator"
	"github.com/arpb2/C-3PO/pkg/infra/pipeline"
	"github.com/arpb2/C-3PO/pkg/presentation/auth/middleware/single"
	"github.com/arpb2/C-3PO/pkg/presentation/auth/middleware/teacher"
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
			binder.CreateBinder(),
			binder4.CreateBinder(httpPipeline, teacherAuthMiddleware, userLevelService),
			binder3.CreateBinder(httpPipeline, singleAuthMiddleware, userService),
			binder2.CreateBinder(httpPipeline, tokenHandler, credentialService),
		}, func() {
			_ = driver.Close()
		}
}

func RegisterRoutes(engine engine.ServerEngine, binders []engine.ControllerBinder) {
	for _, binder := range binders {
		binder.BindControllers(engine)
	}
}
