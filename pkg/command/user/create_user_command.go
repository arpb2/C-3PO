package user

import (
	userservice "github.com/arpb2/C-3PO/api/service/user"
	httppipeline "github.com/arpb2/C-3PO/pkg/pipeline"
	"github.com/saantiaguilera/go-pipeline"
)

type createUserCommand struct {
	service userservice.Service
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

	user, err := c.service.CreateUser(authenticatedUser.User)

	if err != nil {
		return err
	}

	ctx.Set(TagUser, user)
	return nil
}

func CreateCreateUserCommand(service userservice.Service) pipeline.Step {
	return &createUserCommand{
		service: service,
	}
}
