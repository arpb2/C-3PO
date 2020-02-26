package controller

import (
	"fmt"

	controller2 "github.com/arpb2/C-3PO/pkg/domain/user/controller"
	"github.com/arpb2/C-3PO/pkg/domain/user/service"

	"github.com/arpb2/C-3PO/pkg/presentation/user/command"
	"github.com/arpb2/C-3PO/pkg/presentation/user/validation"

	"github.com/arpb2/C-3PO/pkg/domain/infrastructure/controller"
	"github.com/arpb2/C-3PO/pkg/domain/infrastructure/http"
	"github.com/arpb2/C-3PO/pkg/domain/infrastructure/pipeline"
	gopipeline "github.com/saantiaguilera/go-pipeline"
)

func CreatePutController(exec pipeline.HttpPipeline, authMiddleware http.Handler, userService service.Service, validations []validation.Validation) controller.Controller {
	return controller.Controller{
		Method: "PUT",
		Path:   fmt.Sprintf("/users/:%s", controller2.ParamUserId),
		Middleware: []http.Handler{
			authMiddleware,
		},
		Body: CreatePutBody(exec, userService, validations),
	}
}

func CreatePutBody(exec pipeline.HttpPipeline, userService service.Service, validations []validation.Validation) http.Handler {
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
