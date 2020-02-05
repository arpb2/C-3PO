package user

import (
	"github.com/arpb2/C-3PO/api/http"
	"github.com/arpb2/C-3PO/api/model"
	userservice "github.com/arpb2/C-3PO/api/service/user"
	"github.com/saantiaguilera/go-pipeline"
)

type updateUserCommand struct {
	service userservice.Service
}

func (c *updateUserCommand) Name() string {
	return "update_user_command"
}

func (c *updateUserCommand) Run(ctx pipeline.Context) error {
	value, existsUser := ctx.Get(TagAuthenticatedUser)
	userId, existsUserId := ctx.GetUInt(TagUserId)

	if !existsUser || !existsUserId {
		return http.CreateInternalError()
	}

	authenticatedUser := value.(model.AuthenticatedUser)
	authenticatedUser.Id = userId

	user, err := c.service.UpdateUser(&authenticatedUser)

	if err != nil {
		return err
	}

	if user == nil {
		return http.CreateInternalError()
	}

	ctx.Set(TagUser, *user)
	return nil
}

func CreateUpdateUserCommand(service userservice.Service) pipeline.Step {
	return &updateUserCommand{
		service: service,
	}
}
