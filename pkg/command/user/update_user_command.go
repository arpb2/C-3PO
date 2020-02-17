package user

import (
	"github.com/arpb2/C-3PO/api/http"
	userservice "github.com/arpb2/C-3PO/api/service/user"
	httppipeline "github.com/arpb2/C-3PO/pkg/pipeline"
	"github.com/saantiaguilera/go-pipeline"
)

type updateUserCommand struct {
	service userservice.Service
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

	user, err := c.service.UpdateUser(authenticatedUser.User)

	if err != nil {
		return err
	}

	ctx.Set(TagUser, user)
	return nil
}

func CreateUpdateUserCommand(service userservice.Service) pipeline.Step {
	return &updateUserCommand{
		service: service,
	}
}
