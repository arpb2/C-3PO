package command

import (
	"github.com/arpb2/C-3PO/pkg/domain/architecture/http"
	model2 "github.com/arpb2/C-3PO/pkg/domain/session/model"
	"github.com/arpb2/C-3PO/pkg/domain/session/repository"
	"github.com/arpb2/C-3PO/pkg/presentation/user/command"
	"github.com/saantiaguilera/go-pipeline"
)

type createSessionCommand struct {
	tokenHandler repository.TokenRepository
}

func (c *createSessionCommand) Name() string {
	return "create_session_command"
}

func (c *createSessionCommand) Run(ctx pipeline.Context) error {
	userId, exists := ctx.GetUInt(command.TagUserId)

	if !exists {
		return http.CreateInternalError()
	}

	token, err := c.tokenHandler.Create(&repository.Token{
		UserId: userId,
	})

	if err != nil {
		return err
	}

	ctx.Set(TagSession, model2.Session{
		UserId: userId,
		Token:  token,
	})
	return nil
}

func CreateCreateSessionCommand(tokenHandler repository.TokenRepository) pipeline.Step {
	return &createSessionCommand{
		tokenHandler: tokenHandler,
	}
}
