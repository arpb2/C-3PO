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
	fetchUserCommand := usercommand.CreateFetchAuthenticatedUserCommand()
	validateCommand := usercommand.CreateValidateUserParametersCommand(validations)
	serviceCommand := usercommand.CreateCreateUserCommand(userService)
	renderCommand := usercommand.CreateRenderUserCommand()

	graph := gopipeline.CreateSequentialStage(
		fetchUserCommand,
		validateCommand,
		serviceCommand,
		renderCommand,
	)

	return func(ctx *http.Context) {
		exec.Run(ctx, graph)
	}
}
