package code

import (
	"github.com/arpb2/C-3PO/api/http"
	"github.com/arpb2/C-3PO/api/model"
	codeservice "github.com/arpb2/C-3PO/api/service/code"
	usercommand "github.com/arpb2/C-3PO/pkg/command/user"
	"github.com/saantiaguilera/go-pipeline"
)

type replaceCodeCommand struct {
	service codeservice.Service
}

func (c *replaceCodeCommand) Name() string {
	return "replace_code_command"
}

func (c *replaceCodeCommand) Run(ctx pipeline.Context) error {
	codeId, existsCodeId := ctx.GetUInt(TagCodeId)
	userId, existsUserId := ctx.GetUInt(usercommand.TagUserId)
	codeRaw, existsCode := ctx.GetString(TagCodeRaw)

	if !existsCodeId || !existsUserId || !existsCode {
		return http.CreateInternalError()
	}

	code := &model.Code{
		Id:     codeId,
		UserId: userId,
		Code:   codeRaw,
	}

	err := c.service.ReplaceCode(code)

	if err != nil {
		return err
	}

	ctx.Set(TagCode, *code)
	return nil
}

func CreateReplaceCodeCommand(service codeservice.Service) pipeline.Step {
	return &replaceCodeCommand{
		service: service,
	}
}
