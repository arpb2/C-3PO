package command

import (
	"github.com/arpb2/C-3PO/pkg/domain/auth"
	"github.com/arpb2/C-3PO/pkg/domain/http"
	"github.com/arpb2/C-3PO/pkg/domain/model"
	"github.com/arpb2/C-3PO/pkg/presentation/user/command"
	"github.com/saantiaguilera/go-pipeline"
)

type createSessionCommand struct {
	tokenHandler auth.TokenHandler
}

func (c *createSessionCommand) Name() string {
	return "create_session_command"
}

func (c *createSessionCommand) Run(ctx pipeline.Context) error {
	userId, exists := ctx.GetUInt(command.TagUserId)

	if !exists {
		return http.CreateInternalError()
	}

	token, err := c.tokenHandler.Create(&auth.Token{
		UserId: userId,
	})

	if err != nil {
		return err
	}

	ctx.Set(TagSession, model.Session{
		UserId: userId,
		Token:  token,
	})
	return nil
}

func CreateCreateSessionCommand(tokenHandler auth.TokenHandler) pipeline.Step {
	return &createSessionCommand{
		tokenHandler: tokenHandler,
	}
}
