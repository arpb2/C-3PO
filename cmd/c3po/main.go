package main

import (
	"fmt"
	"log"
	"os"

	"github.com/arpb2/C-3PO/pkg/data/jwt"
	"github.com/arpb2/C-3PO/pkg/data/mysql/client"
	"github.com/arpb2/C-3PO/pkg/domain/controller"
	"github.com/arpb2/C-3PO/pkg/infra/engine/gin"
	"github.com/arpb2/C-3PO/pkg/infra/executor"
	"github.com/arpb2/C-3PO/pkg/infra/executor/decorator"
	"github.com/arpb2/C-3PO/pkg/infra/pipeline"
	"github.com/arpb2/C-3PO/pkg/infra/server"
	health "github.com/arpb2/C-3PO/pkg/presentation/health/controller"
	level "github.com/arpb2/C-3PO/pkg/presentation/level/controller"
	"github.com/arpb2/C-3PO/pkg/presentation/middleware/admin"
	"github.com/arpb2/C-3PO/pkg/presentation/middleware/user/single"
	"github.com/arpb2/C-3PO/pkg/presentation/middleware/user/teacher"
	session "github.com/arpb2/C-3PO/pkg/presentation/session/controller"
	user "github.com/arpb2/C-3PO/pkg/presentation/user/controller"
	"github.com/arpb2/C-3PO/pkg/presentation/user/validation"
	userlevel "github.com/arpb2/C-3PO/pkg/presentation/user_level/controller"

	credentialservice "github.com/arpb2/C-3PO/pkg/data/mysql/service/credential"
	levelservice "github.com/arpb2/C-3PO/pkg/data/mysql/service/level"
	teacherservice "github.com/arpb2/C-3PO/pkg/data/mysql/service/teacher"
	userservice "github.com/arpb2/C-3PO/pkg/data/mysql/service/user"
	userlevelservice "github.com/arpb2/C-3PO/pkg/data/mysql/service/user_level"
)

const (
	envPort             = "PORT"
	envMysqlDSN         = "MYSQL_DSN"
	envSecretJWT        = "SECRET_JWT"
	envSecretAdminToken = "SECRET_TOKEN_ADMIN"

	defaultPort = "8080"
)

func assertEnv(env string) string {
	v, exists := os.LookupEnv(env)
	if !exists {
		panic(fmt.Sprintf("No %s envvar found. Maybe you forgot to add it?", env))
	}
	return v
}

func main() {
	var port = os.Getenv(envPort)
	if len(port) == 0 {
		port = defaultPort
	}
	engine := gin.CreateEngine(port)

	traceDecorator := decorator.CreateTraceDecorator(os.Stdout)

	dbClient, driver := client.CreateMysqlClient(assertEnv(envMysqlDSN))
	defer driver.Close()

	httpExecutor := executor.CreateHttpExecutor(traceDecorator)
	httpPipeline := pipeline.CreateHttpPipeline(httpExecutor)

	tokenHandler := jwt.CreateTokenHandler([]byte(assertEnv(envSecretJWT)))

	userService := userservice.CreateService(dbClient)
	teacherService := teacherservice.CreateService(userService, dbClient)
	levelService := levelservice.CreateService(dbClient)
	userLevelService := userlevelservice.CreateService(dbClient)
	credentialService := credentialservice.CreateService(dbClient)

	adminAuthMiddleware := admin.CreateMiddleware([]byte(assertEnv(envSecretAdminToken)))
	singleAuthMiddleware := single.CreateMiddleware(tokenHandler)
	teacherAuthMiddleware := teacher.CreateMiddleware(tokenHandler, teacherService)

	emptyEmailValidation := validation.EmptyEmail
	emptyNameValidation := validation.EmptyName
	emptyPasswordValidation := validation.EmptyPassword
	emptySurnameValidation := validation.EmptySurname
	idProvidedValidation := validation.IdProvided
	securePasswordValidation := validation.SecurePassword

	controllers := []controller.Controller{
		health.CreateGetController(),

		session.CreatePostController(httpPipeline, tokenHandler, credentialService, []validation.Validation{
			validation.EmptyUser,
			validation.EmptyEmail,
			validation.EmptyPassword,
		}),

		user.CreatePostController(httpPipeline, userService, []validation.Validation{
			emptyEmailValidation,
			emptyNameValidation,
			emptySurnameValidation,
			emptyPasswordValidation,
			securePasswordValidation,
			idProvidedValidation,
		}),
		user.CreateGetController(httpPipeline, singleAuthMiddleware, userService),
		user.CreatePutController(httpPipeline, singleAuthMiddleware, userService, []validation.Validation{
			idProvidedValidation,
			securePasswordValidation,
		}),
		user.CreateDeleteController(httpPipeline, singleAuthMiddleware, userService),

		userlevel.CreateGetController(httpPipeline, teacherAuthMiddleware, userLevelService),
		userlevel.CreatePutController(httpPipeline, teacherAuthMiddleware, userLevelService),

		level.CreateGetController(httpPipeline, levelService),
		level.CreatePutController(httpPipeline, adminAuthMiddleware, levelService),
	}

	if err := server.StartApplication(engine, controllers); err != nil {
		log.Fatal(err)
	}
}
