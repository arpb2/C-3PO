package command

import (
	httppipeline "github.com/arpb2/C-3PO/pkg/domain/infrastructure/pipeline"
	credentialservice "github.com/arpb2/C-3PO/pkg/domain/session/service"
	"github.com/arpb2/C-3PO/pkg/presentation/user/command"
	"github.com/saantiaguilera/go-pipeline"
)

type authenticateCommand struct {
	service credentialservice.Service
}

func (c *authenticateCommand) Name() string {
	return "authenticate_command"
}

func (c *authenticateCommand) Run(ctx pipeline.Context) error {
	ctxAware := httppipeline.CreateContextAware(ctx)

	authenticatedUser, err := ctxAware.GetAuthenticatedUser(command.TagAuthenticatedUser)

	if err != nil {
		return err
	}

	userId, err := c.service.GetUserId(authenticatedUser.Email, authenticatedUser.Password)

	if err != nil {
		return err
	}

	ctx.Set(command.TagUserId, userId)
	return nil
}

func CreateAuthenticateCommand(service credentialservice.Service) pipeline.Step {
	return &authenticateCommand{
		service: service,
	}
}
