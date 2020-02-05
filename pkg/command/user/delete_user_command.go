package user

import (
	"github.com/arpb2/C-3PO/api/http"
	userservice "github.com/arpb2/C-3PO/api/service/user"
	"github.com/saantiaguilera/go-pipeline"
)

type deleteUserCommand struct {
	service userservice.Service
}

func (c *deleteUserCommand) Name() string {
	return "delete_user_command"
}

func (c *deleteUserCommand) Run(ctx pipeline.Context) error {
	userId, exists := ctx.GetUInt(TagUserId)

	if !exists {
		return http.CreateInternalError()
	}

	return c.service.DeleteUser(userId)
}

func CreateDeleteUserCommand(service userservice.Service) pipeline.Step {
	return &deleteUserCommand{
		service: service,
	}
}
