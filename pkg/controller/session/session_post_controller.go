package session

import (
	"github.com/arpb2/C-3PO/api/auth"
	"github.com/arpb2/C-3PO/api/controller"
	"github.com/arpb2/C-3PO/api/http"
	"github.com/arpb2/C-3PO/api/pipeline"
	credentialservice "github.com/arpb2/C-3PO/api/service/credential"
	sessioncommand "github.com/arpb2/C-3PO/pkg/command/session"
	usercommand "github.com/arpb2/C-3PO/pkg/command/user"
	uservalidation "github.com/arpb2/C-3PO/pkg/validation/user"
	gopipeline "github.com/saantiaguilera/go-pipeline"
)

func CreatePostController(executor pipeline.HttpPipeline,
	tokenHandler auth.TokenHandler,
	service credentialservice.Service,
	validations []uservalidation.Validation) controller.Controller {
	return controller.Controller{
		Method: "POST",
		Path:   "/session",
		Body:   CreatePostBody(executor, tokenHandler, service, validations),
	}
}

func CreatePostBody(executor pipeline.HttpPipeline, tokenHandler auth.TokenHandler, service credentialservice.Service, validations []uservalidation.Validation) http.Handler {
	fetchUserCommand := usercommand.CreateFetchAuthenticatedUserCommand()
	validateParamsCommand := usercommand.CreateValidateParametersCommand(validations)
	authenticateCommand := sessioncommand.CreateAuthenticateCommand(service)
	createSessionCommand := sessioncommand.CreateCreateSessionCommand(tokenHandler)
	renderCommand := sessioncommand.CreateRenderSessionCommand()

	graph := gopipeline.CreateSequentialStage(
		fetchUserCommand,
		validateParamsCommand,
		authenticateCommand,
		createSessionCommand,
		renderCommand,
	)

	return func(ctx *http.Context) {
		executor.Run(ctx, graph)
	}
}
