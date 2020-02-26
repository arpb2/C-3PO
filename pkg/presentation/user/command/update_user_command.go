package command

import (
	"github.com/arpb2/C-3PO/pkg/domain/architecture/http"
	httppipeline "github.com/arpb2/C-3PO/pkg/domain/architecture/pipeline"
	service2 "github.com/arpb2/C-3PO/pkg/domain/user/service"
	"github.com/saantiaguilera/go-pipeline"
)

type updateUserCommand struct {
	service service2.Service
}

func (c *updateUserCommand) Name() string {
	return "update_user_command"
}

func (c *updateUserCommand) Run(ctx pipeline.Context) error {
	ctxAware := httppipeline.CreateContextAware(ctx)

	authenticatedUser, errAuthenticatedUser := ctxAware.GetAuthenticatedUser(TagAuthenticatedUser)
	userId, existsUserId := ctx.GetUInt(TagUserId)

	if errAuthenticatedUser != nil || !existsUserId {
		return http.CreateInternalError()
	}

	authenticatedUser.Id = userId

	user, err := c.service.UpdateUser(authenticatedUser)

	if err != nil {
		return err
	}

	ctx.Set(TagUser, user)
	return nil
}

func CreateUpdateUserCommand(service service2.Service) pipeline.Step {
	return &updateUserCommand{
		service: service,
	}
}
