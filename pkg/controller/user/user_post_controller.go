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

func CreatePostController(exec pipeline.HttpPipeline, validations []uservalidation.Validation, userService userservice.Service) controller.Controller {
	return controller.Controller{
		Method: "POST",
		Path:   "/users",
		Body:   CreatePostBody(exec, validations, userService),
	}
}

func CreatePostBody(exec pipeline.HttpPipeline, validations []uservalidation.Validation, userService userservice.Service) http.Handler {
	return func(ctx *http.Context) {
		fetchUserCommand := usercommand.CreateFetchAuthenticatedUserCommand(ctx)
		validateCommand := usercommand.CreateValidateParametersCommand(ctx, fetchUserCommand.OutputStream, validations)
		serviceCommand := usercommand.CreateCreateUserCommand(ctx, userService, validateCommand.OutputStream)
		renderCommand := usercommand.CreateRenderUserCommand(ctx, serviceCommand.OutputStream)

		graph := gopipeline.CreateSequentialStage(
			fetchUserCommand,
			validateCommand,
			serviceCommand,
			renderCommand,
		)

		exec.Run(ctx, graph)
	}
}
