package user_level

import (
	"github.com/arpb2/C-3PO/api/http"
	userlevelservice "github.com/arpb2/C-3PO/api/service/user_level"
	usercommand "github.com/arpb2/C-3PO/pkg/command/user"
	"github.com/saantiaguilera/go-pipeline"
)

type createUserLevelCommand struct {
	service userlevelservice.Service
}

func (c *createUserLevelCommand) Name() string {
	return "create_user_level_command"
}

func (c *createUserLevelCommand) Run(ctx pipeline.Context) error {
	codeRaw, existsCode := ctx.GetString(TagCodeRaw)
	userId, existsUserId := ctx.GetUInt(usercommand.TagUserId)

	if !existsCode || !existsUserId {
		return http.CreateInternalError()
	}

	userLevel, err := c.service.CreateUserLevel(userId, codeRaw)

	if err != nil {
		return err
	}

	ctx.Set(TagUserLevel, *userLevel)
	return nil
}

func CreateCreateUserLevelCommand(service userlevelservice.Service) pipeline.Step {
	return &createUserLevelCommand{
		service: service,
	}
}
