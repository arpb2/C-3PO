package controller

import (
	"fmt"
	"github.com/arpb2/C-3PO/pkg/presentation/user/command"
	"github.com/arpb2/C-3PO/pkg/presentation/user/validation"

	"github.com/arpb2/C-3PO/pkg/domain/controller"
	"github.com/arpb2/C-3PO/pkg/domain/http"
	"github.com/arpb2/C-3PO/pkg/domain/pipeline"
	userservice "github.com/arpb2/C-3PO/pkg/domain/service/user"
	gopipeline "github.com/saantiaguilera/go-pipeline"
)

func CreatePutController(exec pipeline.HttpPipeline, validations []validation.Validation, authMiddleware http.Handler, userService userservice.Service) controller.Controller {
	return controller.Controller{
		Method: "PUT",
		Path:   fmt.Sprintf("/users/:%s", controller.ParamUserId),
		Middleware: []http.Handler{
			authMiddleware,
		},
		Body: CreatePutBody(exec, validations, userService),
	}
}

func CreatePutBody(exec pipeline.HttpPipeline, validations []validation.Validation, userService userservice.Service) http.Handler {
	fetchUserIdCommand := command.CreateFetchUserIdCommand()
	fetchUserCommand := command.CreateFetchAuthenticatedUserCommand()
	validateCommand := command.CreateValidateUserParametersCommand(validations)
	serviceCommand := command.CreateUpdateUserCommand(userService)
	renderCommand := command.CreateRenderUserCommand()

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
