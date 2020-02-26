package command

import (
	"github.com/arpb2/C-3PO/pkg/domain/architecture/http"
	repository2 "github.com/arpb2/C-3PO/pkg/domain/user/repository"
	"github.com/saantiaguilera/go-pipeline"
)

type getUserCommand struct {
	repository repository2.UserRepository
}

func (c *getUserCommand) Name() string {
	return "get_user_command"
}

func (c *getUserCommand) Run(ctx pipeline.Context) error {
	userId, existsUserId := ctx.GetUInt(TagUserId)

	if !existsUserId {
		return http.CreateInternalError()
	}

	user, err := c.repository.GetUser(userId)

	if err != nil {
		return err
	}

	ctx.Set(TagUser, user)
	return nil
}

func CreateGetUserCommand(repository repository2.UserRepository) pipeline.Step {
	return &getUserCommand{
		repository: repository,
	}
}
