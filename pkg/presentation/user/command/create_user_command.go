package command

import (
	httppipeline "github.com/arpb2/C-3PO/pkg/domain/architecture/pipeline"
	repository2 "github.com/arpb2/C-3PO/pkg/domain/user/repository"
	"github.com/saantiaguilera/go-pipeline"
)

type createUserCommand struct {
	repository repository2.UserRepository
}

func (c *createUserCommand) Name() string {
	return "create_user_command"
}

func (c *createUserCommand) Run(ctx pipeline.Context) error {
	ctxAware := httppipeline.CreateContextAware(ctx)

	authenticatedUser, err := ctxAware.GetAuthenticatedUser(TagAuthenticatedUser)

	if err != nil {
		return err
	}

	user, err := c.repository.CreateUser(authenticatedUser)

	if err != nil {
		return err
	}

	ctx.Set(TagUser, user)
	return nil
}

func CreateCreateUserCommand(repository repository2.UserRepository) pipeline.Step {
	return &createUserCommand{
		repository: repository,
	}
}
