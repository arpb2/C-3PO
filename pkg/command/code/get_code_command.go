package code

import (
	"github.com/arpb2/C-3PO/api/http"
	codeservice "github.com/arpb2/C-3PO/api/service/code"
	usercommand "github.com/arpb2/C-3PO/pkg/command/user"
	"github.com/saantiaguilera/go-pipeline"
)

type getCodeCommand struct {
	service codeservice.Service
}

func (c *getCodeCommand) Name() string {
	return "get_code_command"
}

func (c *getCodeCommand) Run(ctx pipeline.Context) error {
	codeId, existsCodeId := ctx.GetUInt(TagCodeId)
	userId, existsUserId := ctx.GetUInt(usercommand.TagUserId)

	if !existsCodeId || !existsUserId {
		return http.CreateInternalError()
	}

	code, err := c.service.GetCode(userId, codeId)

	if err != nil {
		return err
	}

	if code == nil {
		return http.CreateNotFoundError()
	}

	ctx.Set(TagCode, *code)
	return nil
}

func CreateGetCodeCommand(service codeservice.Service) pipeline.Step {
	return &getCodeCommand{
		service: service,
	}
}
