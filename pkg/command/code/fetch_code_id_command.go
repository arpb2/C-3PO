package code

import (
	"fmt"
	"strconv"

	"github.com/arpb2/C-3PO/pkg/controller"

	httppipeline "github.com/arpb2/C-3PO/pkg/pipeline"

	"github.com/saantiaguilera/go-pipeline"

	"github.com/arpb2/C-3PO/api/http"
)

type fetchCodeIdCommand struct {
	context *http.Context
}

func (c *fetchCodeIdCommand) Name() string {
	return fmt.Sprintf("fetch_%s_command", controller.ParamCodeId)
}

func (c *fetchCodeIdCommand) Run(ctx pipeline.Context) error {
	ctxAware := httppipeline.CreateContextAware(ctx)

	httpReader, err := ctxAware.GetReader()

	if err != nil {
		return err
	}

	codeId := httpReader.GetParameter(controller.ParamCodeId)

	if codeId == "" {
		return http.CreateBadRequestError(fmt.Sprintf("'%s' empty", controller.ParamCodeId))
	}

	codeIdUint, err := strconv.ParseUint(codeId, 10, 64)

	if err != nil {
		return http.CreateBadRequestError(fmt.Sprintf("'%s' malformed, expecting a positive number", controller.ParamCodeId))
	}

	ctx.Set(TagCodeId, uint(codeIdUint))
	return nil
}

func CreateFetchCodeIdCommand() pipeline.Step {
	return &fetchCodeIdCommand{}
}
