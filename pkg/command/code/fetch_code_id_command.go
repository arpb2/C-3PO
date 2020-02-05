package code

import (
	"strconv"

	"github.com/arpb2/C-3PO/pkg/command"
	"github.com/saantiaguilera/go-pipeline"

	"github.com/arpb2/C-3PO/api/http"
)

type fetchCodeIdCommand struct {
	context *http.Context
}

func (c *fetchCodeIdCommand) Name() string {
	return "fetch_code_id_command"
}

func (c *fetchCodeIdCommand) Run(ctx pipeline.Context) error {
	httpReader, exists := ctx.Get(command.TagHttpReader)

	if !exists {
		return http.CreateInternalError()
	}

	codeId := httpReader.(http.Reader).GetParameter("code_id")

	if codeId == "" {
		return http.CreateBadRequestError("'code_id' empty")
	}

	codeIdUint, err := strconv.ParseUint(codeId, 10, 64)

	if err != nil {
		return http.CreateBadRequestError("'code_id' malformed, expecting a positive number")
	}

	ctx.Set(TagCodeId, uint(codeIdUint))
	return nil
}

func CreateFetchCodeIdCommand() pipeline.Step {
	return &fetchCodeIdCommand{}
}
