package code

import (
	"github.com/arpb2/C-3PO/api/http"
	"github.com/arpb2/C-3PO/pkg/command"
	"github.com/saantiaguilera/go-pipeline"
)

type fetchCodeCommand struct{}

func (c *fetchCodeCommand) Name() string {
	return "fetch_code_command"
}

func (c *fetchCodeCommand) Run(ctx pipeline.Context) error {
	httpReader, exists := ctx.Get(command.TagHttpReader)

	if !exists {
		return http.CreateInternalError()
	}

	code, exists := httpReader.(http.Reader).GetFormData("code")

	if !exists {
		return http.CreateBadRequestError("'code' part not found")
	}

	ctx.Set(TagCodeRaw, code)
	return nil
}

func CreateFetchCodeCommand() pipeline.Step {
	return &fetchCodeCommand{}
}
