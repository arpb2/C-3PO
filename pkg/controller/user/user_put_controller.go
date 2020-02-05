package user

import (
	"github.com/arpb2/C-3PO/api/controller"
	"github.com/arpb2/C-3PO/api/http"
	"github.com/arpb2/C-3PO/api/pipeline"
	userservice "github.com/arpb2/C-3PO/api/service/user"
	usercommand "github.com/arpb2/C-3PO/pkg/command/user"
	uservalidation "github.com/arpb2/C-3PO/pkg/validation/user"
	gopipeline "github.com/saantiaguilera/go-pipeline"
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
	fetchUserIdCommand := usercommand.CreateFetchUserIdCommand()
	fetchUserCommand := usercommand.CreateFetchAuthenticatedUserCommand()
	validateCommand := usercommand.CreateValidateParametersCommand(validations)
	serviceCommand := usercommand.CreateUpdateUserCommand(userService)
	renderCommand := usercommand.CreateRenderUserCommand()

	graph := gopipeline.CreateSequentialGroup(
		gopipeline.CreateConcurrentGroup(
			gopipeline.CreateSequentialStage(
				fetchUserIdCommand,
			),
			gopipeline.CreateSequentialStage(
				fetchUserCommand,
				validateCommand,
			),
		),
		gopipeline.CreateSequentialStage(
			serviceCommand,
			renderCommand,
		),
	)

	return func(ctx *http.Context) {
		exec.Run(ctx, graph)
	}
}
