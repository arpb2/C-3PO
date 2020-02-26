package command

import (
	"github.com/arpb2/C-3PO/pkg/domain/architecture/http"
	repository2 "github.com/arpb2/C-3PO/pkg/domain/user/repository"
	"github.com/saantiaguilera/go-pipeline"
)

type deleteUserCommand struct {
	repository repository2.UserRepository
}

func (c *deleteUserCommand) Name() string {
	return "delete_user_command"
}

func (c *deleteUserCommand) Run(ctx pipeline.Context) error {
	userId, exists := ctx.GetUInt(TagUserId)

	if !exists {
		return http.CreateInternalError()
	}

	return c.repository.DeleteUser(userId)
}

func CreateDeleteUserCommand(repository repository2.UserRepository) pipeline.Step {
	return &deleteUserCommand{
		repository: repository,
	}
}
