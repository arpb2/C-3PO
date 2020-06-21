package main

import (
	"fmt"
	"log"
	"os"

	"github.com/arpb2/C-3PO/pkg/data/usecase/user/validation"
	"github.com/arpb2/C-3PO/pkg/infrastructure/gin"
	"github.com/arpb2/C-3PO/pkg/infrastructure/hystrix"
	"github.com/arpb2/C-3PO/pkg/infrastructure/pipeline"
	"github.com/arpb2/C-3PO/pkg/presentation/health"
	"github.com/arpb2/C-3PO/pkg/presentation/level"
	"github.com/arpb2/C-3PO/pkg/presentation/session"
	"github.com/arpb2/C-3PO/pkg/presentation/user"

	"github.com/arpb2/C-3PO/pkg/infrastructure/jwt"
	"github.com/arpb2/C-3PO/pkg/infrastructure/mysql"

	"github.com/arpb2/C-3PO/pkg/infrastructure/hystrix/decorator"
	credentialrepository "github.com/arpb2/C-3PO/pkg/infrastructure/mysql/credential"
	levelrepository "github.com/arpb2/C-3PO/pkg/infrastructure/mysql/level"
	userrepository "github.com/arpb2/C-3PO/pkg/infrastructure/mysql/user"
)

const (
	envPort             = "PORT"
	envMysqlDSN         = "MYSQL_DSN"
	envSecretJWT        = "SECRET_JWT"
	envSecretAdminToken = "SECRET_TOKEN_ADMIN"

	ParamLevelId = "level_id"
	ParamUserId  = "user_id"

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
	engine := gin.CreateEngine()

	traceDecorator := decorator.CreateTraceDecorator(os.Stdout)

	dbClient, driver := mysql.CreateMysqlClient(assertEnv(envMysqlDSN))
	defer driver.Close()

	httpExecutor := hystrix.CreateHttpExecutor(traceDecorator)
	httpPipeline := pipeline.CreateHttpPipeline(httpExecutor)

	tokenHandler := jwt.CreateTokenRepository([]byte(assertEnv(envSecretJWT)))

	userRepository := userrepository.CreateUserRepository(dbClient)
	levelRepository := levelrepository.CreateRepository(dbClient)
	userLevelRepository := userrepository.CreateLevelRepository(dbClient)
	credentialRepository := credentialrepository.CreateRepository(dbClient)

	debugAuthMiddleware := session.CreateAuthenticateDebugMiddleware()
	adminAuthMiddleware := session.CreateAuthenticateAdminMiddleware([]byte(assertEnv(envSecretAdminToken)))
	singleAuthMiddleware := session.CreateAuthenticateUserMiddleware(ParamUserId, tokenHandler)

	emptyUserValidation := validation.EmptyUser
	emptyEmailValidation := validation.EmptyEmail
	emptyNameValidation := validation.EmptyName
	emptyPasswordValidation := validation.EmptyPassword
	emptySurnameValidation := validation.EmptySurname
	idProvidedValidation := validation.IdProvided
	malformedTypeValidation := validation.MalformedType
	typeProvidedValidation := validation.TypeProvided
	securePasswordValidation := validation.SecurePassword

	healthGetHandler := health.CreateGetHandler()
	sessionPostHandler := session.CreatePostHandler(httpPipeline, tokenHandler, credentialRepository, []validation.Validation{
		emptyUserValidation,
		emptyEmailValidation,
		emptyPasswordValidation,
	})
	userGetHandler := user.CreateGetUserHandler(ParamUserId, httpPipeline, userRepository)
	userPostHandler := user.CreatePostUserHandler(httpPipeline, userRepository, []validation.Validation{
		emptyEmailValidation,
		emptyNameValidation,
		emptySurnameValidation,
		emptyPasswordValidation,
		securePasswordValidation,
		idProvidedValidation,
		malformedTypeValidation,
	})
	userPutHandler := user.CreatePutUserHandler(ParamUserId, httpPipeline, userRepository, []validation.Validation{
		idProvidedValidation,
		typeProvidedValidation,
		securePasswordValidation,
	})
	userDeleteHandler := user.CreateDeleteUserHandler(ParamUserId, httpPipeline, userRepository)
	userLevelGetHandler := user.CreateGetLevelHandler(ParamUserId, ParamLevelId, httpPipeline, userLevelRepository)
	userLevelPutHandler := user.CreatePutLevelHandler(
		ParamUserId, ParamLevelId,
		httpPipeline,
		userLevelRepository,
	)
	levelGetHandler := level.CreateGetHandler(ParamLevelId, httpPipeline, levelRepository)
	levelPutHandler := level.CreatePutHandler(ParamLevelId, httpPipeline, levelRepository)

	/****** Global middle-wares ******/
	engine.Use(debugAuthMiddleware)

	/****** Health routes ******/
	engine.GET("/ping", healthGetHandler)

	/****** Session routes ******/
	engine.POST("/session", sessionPostHandler)

	/****** User routes ******/
	engine.GET(
		fmt.Sprintf("/users/:%s", ParamUserId),
		singleAuthMiddleware, userGetHandler,
	)
	engine.POST("/users", userPostHandler)
	engine.PUT(
		fmt.Sprintf("/users/:%s", ParamUserId),
		singleAuthMiddleware, userPutHandler)
	engine.DELETE(
		fmt.Sprintf("/users/:%s", ParamUserId),
		singleAuthMiddleware, userDeleteHandler,
	)

	/****** User level routes ******/
	engine.GET(
		fmt.Sprintf("/users/:%s/levels/:%s", ParamUserId, ParamLevelId),
		singleAuthMiddleware, userLevelGetHandler,
	)
	engine.PUT(
		fmt.Sprintf("/users/:%s/levels/:%s", ParamUserId, ParamLevelId),
		singleAuthMiddleware, userLevelPutHandler,
	)

	/****** Level routes ******/
	engine.GET(
		fmt.Sprintf("/levels/:%s", ParamLevelId),
		levelGetHandler,
	)
	engine.PUT(
		fmt.Sprintf("/levels/:%s", ParamLevelId),
		adminAuthMiddleware, levelPutHandler,
	)

	if err := engine.Run(port); err != nil {
		log.Fatalf("error running server %s", err.Error())
	}
}
