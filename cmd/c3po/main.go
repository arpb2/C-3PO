package main

import (
	"fmt"
	"log"
	"os"

	"github.com/arpb2/C-3PO/pkg/data/jwt"
	"github.com/arpb2/C-3PO/pkg/data/mysql"

	"github.com/arpb2/C-3PO/cmd/c3po/infrastructure/engine/gin"
	"github.com/arpb2/C-3PO/cmd/c3po/infrastructure/executor"
	"github.com/arpb2/C-3PO/cmd/c3po/infrastructure/executor/decorator"
	"github.com/arpb2/C-3PO/cmd/c3po/infrastructure/pipeline"
	"github.com/arpb2/C-3PO/cmd/c3po/infrastructure/server"
	"github.com/arpb2/C-3PO/pkg/domain/architecture/controller"
	health "github.com/arpb2/C-3PO/pkg/presentation/health/controller"
	level "github.com/arpb2/C-3PO/pkg/presentation/level/controller"
	"github.com/arpb2/C-3PO/pkg/presentation/middleware/auth/admin"
	"github.com/arpb2/C-3PO/pkg/presentation/middleware/auth/user/single"
	"github.com/arpb2/C-3PO/pkg/presentation/middleware/auth/user/teacher"
	session "github.com/arpb2/C-3PO/pkg/presentation/session/controller"
	user "github.com/arpb2/C-3PO/pkg/presentation/user/controller"
	"github.com/arpb2/C-3PO/pkg/presentation/user/validation"
	userlevel "github.com/arpb2/C-3PO/pkg/presentation/user_level/controller"

	credentialrepository "github.com/arpb2/C-3PO/pkg/data/mysql/credential"
	levelrepository "github.com/arpb2/C-3PO/pkg/data/mysql/level"
	teacherrepository "github.com/arpb2/C-3PO/pkg/data/mysql/teacher"
	userrepository "github.com/arpb2/C-3PO/pkg/data/mysql/user"
	userlevelrepository "github.com/arpb2/C-3PO/pkg/data/mysql/user_level"
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

	dbClient, driver := mysql.CreateMysqlClient(assertEnv(envMysqlDSN))
	defer driver.Close()

	httpExecutor := executor.CreateHttpExecutor(traceDecorator)
	httpPipeline := pipeline.CreateHttpPipeline(httpExecutor)

	tokenHandler := jwt.CreateTokenRepository([]byte(assertEnv(envSecretJWT)))

	userRepository := userrepository.CreateRepository(dbClient)
	teacherRepository := teacherrepository.CreateRepository(userRepository, dbClient)
	levelRepository := levelrepository.CreateRepository(dbClient)
	userLevelRepository := userlevelrepository.CreateRepository(dbClient)
	credentialRepository := credentialrepository.CreateRepository(dbClient)

	adminAuthMiddleware := admin.CreateMiddleware([]byte(assertEnv(envSecretAdminToken)))
	singleAuthMiddleware := single.CreateMiddleware(tokenHandler)
	teacherAuthMiddleware := teacher.CreateMiddleware(tokenHandler, teacherRepository)

	emptyUserValidation := validation.EmptyUser
	emptyEmailValidation := validation.EmptyEmail
	emptyNameValidation := validation.EmptyName
	emptyPasswordValidation := validation.EmptyPassword
	emptySurnameValidation := validation.EmptySurname
	idProvidedValidation := validation.IdProvided
	securePasswordValidation := validation.SecurePassword

	controllers := []controller.Controller{
		health.CreateGetController(),

		session.CreatePostController(httpPipeline, tokenHandler, credentialRepository, []validation.Validation{
			emptyUserValidation,
			emptyEmailValidation,
			emptyPasswordValidation,
		}),

		user.CreatePostController(httpPipeline, userRepository, []validation.Validation{
			emptyEmailValidation,
			emptyNameValidation,
			emptySurnameValidation,
			emptyPasswordValidation,
			securePasswordValidation,
			idProvidedValidation,
		}),
		user.CreateGetController(httpPipeline, singleAuthMiddleware, userRepository),
		user.CreatePutController(httpPipeline, singleAuthMiddleware, userRepository, []validation.Validation{
			idProvidedValidation,
			securePasswordValidation,
		}),
		user.CreateDeleteController(httpPipeline, singleAuthMiddleware, userRepository),

		userlevel.CreateGetController(httpPipeline, teacherAuthMiddleware, userLevelRepository),
		userlevel.CreatePutController(httpPipeline, teacherAuthMiddleware, userLevelRepository),

		level.CreateGetController(httpPipeline, levelRepository),
		level.CreatePutController(httpPipeline, adminAuthMiddleware, levelRepository),
	}

	if err := server.StartApplication(engine, controllers); err != nil {
		log.Fatal(err)
	}
}
