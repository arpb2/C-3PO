package command

import (
	"github.com/arpb2/C-3PO/pkg/domain/infrastructure/http"
	service2 "github.com/arpb2/C-3PO/pkg/domain/user/service"
	"github.com/saantiaguilera/go-pipeline"
)

type deleteUserCommand struct {
	service service2.Service
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

func CreateDeleteUserCommand(service service2.Service) pipeline.Step {
	return &deleteUserCommand{
		service: service,
	}
}
