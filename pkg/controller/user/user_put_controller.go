package user

import (
	"github.com/arpb2/C-3PO/api/controller"
	"github.com/arpb2/C-3PO/api/http"
	"github.com/arpb2/C-3PO/api/pipeline"
	userservice "github.com/arpb2/C-3PO/api/service/user"
	usercommand "github.com/arpb2/C-3PO/pkg/command/user"
	uservalidation "github.com/arpb2/C-3PO/pkg/validation/user"
	"github.com/saantiaguilera/go-pipeline/pkg/stage/concurrent"
	"github.com/saantiaguilera/go-pipeline/pkg/stage/sequential"
)

func CreatePutController(exec pipeline.HttpPipeline, validations []uservalidation.Validation, authMiddleware http.Handler, userService userservice.Service) controller.Controller {
	return controller.Controller{
		Method: "PUT",
		Path:   "/users/:user_id",
		Middleware: []http.Handler{
			authMiddleware,
		},
		Body: CreatePutBody(exec, validations, userService),
	}
}

func CreatePutBody(exec pipeline.HttpPipeline, validations []uservalidation.Validation, userService userservice.Service) http.Handler {
	return func(ctx *http.Context) {
		fetchUserIdCommand := usercommand.CreateFetchUserIdCommand(ctx)
		fetchUserCommand := usercommand.CreateFetchAuthenticatedUserCommand(ctx)
		validateCommand := usercommand.CreateValidateParametersCommand(ctx, fetchUserCommand.OutputStream, validations)
		serviceCommand := usercommand.CreateUpdateUserCommand(ctx, userService, fetchUserIdCommand.OutputStream, validateCommand.OutputStream)
		renderCommand := usercommand.CreateRenderUserCommand(ctx, serviceCommand.OutputStream)

		graph := sequential.CreateSequentialGroup(
			concurrent.CreateConcurrentGroup(
				sequential.CreateSequentialStage(
					fetchUserIdCommand,
				),
				sequential.CreateSequentialStage(
					fetchUserCommand,
					validateCommand,
				),
			),
			sequential.CreateSequentialStage(
				serviceCommand,
				renderCommand,
			),
		)

		exec.Run(ctx, graph)
	}
}
