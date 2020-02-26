package command

import (
	httppipeline "github.com/arpb2/C-3PO/pkg/domain/architecture/pipeline"
	service2 "github.com/arpb2/C-3PO/pkg/domain/user/service"
	"github.com/saantiaguilera/go-pipeline"
)

type createUserCommand struct {
	service service2.Service
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

	user, err := c.service.CreateUser(authenticatedUser)

	if err != nil {
		return err
	}

	ctx.Set(TagUser, user)
	return nil
}

func CreateCreateUserCommand(service service2.Service) pipeline.Step {
	return &createUserCommand{
		service: service,
	}
}
