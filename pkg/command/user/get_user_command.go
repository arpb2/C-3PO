package user

import (
	"github.com/arpb2/C-3PO/api/http"
	userservice "github.com/arpb2/C-3PO/api/service/user"
	"github.com/saantiaguilera/go-pipeline"
)

type getUserCommand struct {
	service userservice.Service
}

func (c *getUserCommand) Name() string {
	return "get_user_command"
}

func (c *getUserCommand) Run(ctx pipeline.Context) error {
	userId, existsUserId := ctx.GetUInt(TagUserId)

	if !existsUserId {
		return http.CreateInternalError()
	}

	user, err := c.service.GetUser(userId)

	if err != nil {
		return err
	}

	if user == nil {
		return http.CreateNotFoundError()
	}

	ctx.Set(TagUser, *user)
	return nil
}

func CreateGetUserCommand(service userservice.Service) pipeline.Step {
	return &getUserCommand{
		service: service,
	}
}
