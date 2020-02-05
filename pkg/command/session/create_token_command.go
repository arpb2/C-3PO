package session

import (
	"github.com/arpb2/C-3PO/api/auth"
	"github.com/arpb2/C-3PO/api/http"
	"github.com/arpb2/C-3PO/api/model"
	"github.com/arpb2/C-3PO/pkg/command/user"
	"github.com/saantiaguilera/go-pipeline"
)

type createSessionCommand struct {
	tokenHandler auth.TokenHandler
}

func (c *createSessionCommand) Name() string {
	return "create_token_command"
}

func (c *createSessionCommand) Run(ctx pipeline.Context) error {
	userId, exists := ctx.GetUInt(user.TagUserId)

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
