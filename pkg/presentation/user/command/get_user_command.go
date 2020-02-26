package command

import (
	"github.com/arpb2/C-3PO/pkg/domain/architecture/http"
	service2 "github.com/arpb2/C-3PO/pkg/domain/user/service"
	"github.com/saantiaguilera/go-pipeline"
)

type getUserCommand struct {
	service service2.Service
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

	ctx.Set(TagUser, user)
	return nil
}

func CreateGetUserCommand(service service2.Service) pipeline.Step {
	return &getUserCommand{
		service: service,
	}
}
