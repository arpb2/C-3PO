package code

import (
	"github.com/arpb2/C-3PO/api/http"
	codeservice "github.com/arpb2/C-3PO/api/service/code"
	usercommand "github.com/arpb2/C-3PO/pkg/command/user"
	"github.com/saantiaguilera/go-pipeline"
)

type createCodeCommand struct {
	service codeservice.Service
}

func (c *createCodeCommand) Name() string {
	return "create_code_command"
}

func (c *createCodeCommand) Run(ctx pipeline.Context) error {
	codeRaw, existsCode := ctx.GetString(TagCodeRaw)
	userId, existsUserId := ctx.GetUInt(usercommand.TagUserId)

	if !existsCode || !existsUserId {
		return http.CreateInternalError()
	}

	code, err := c.service.CreateCode(userId, codeRaw)

	if err != nil {
		return err
	}

	ctx.Set(TagCode, *code)
	return nil
}

func CreateCreateCodeCommand(service codeservice.Service) pipeline.Step {
	return &createCodeCommand{
		service: service,
	}
}
