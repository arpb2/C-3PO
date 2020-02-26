package command

import (
	httppipeline "github.com/arpb2/C-3PO/pkg/domain/architecture/pipeline"
	credentialrepository "github.com/arpb2/C-3PO/pkg/domain/session/repository"
	"github.com/arpb2/C-3PO/pkg/presentation/user/command"
	"github.com/saantiaguilera/go-pipeline"
)

type authenticateCommand struct {
	repository credentialrepository.CredentialRepository
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

	userId, err := c.repository.GetUserId(authenticatedUser.Email, authenticatedUser.Password)

	if err != nil {
		return err
	}

	ctx.Set(command.TagUserId, userId)
	return nil
}

func CreateAuthenticateCommand(repository credentialrepository.CredentialRepository) pipeline.Step {
	return &authenticateCommand{
		repository: repository,
	}
}
