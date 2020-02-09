package code

import (
	"github.com/arpb2/C-3PO/api/http"
	httppipeline "github.com/arpb2/C-3PO/pkg/pipeline"
	"github.com/saantiaguilera/go-pipeline"
)

type fetchCodeCommand struct{}

func (c *fetchCodeCommand) Name() string {
	return "fetch_code_command"
}

func (c *fetchCodeCommand) Run(ctx pipeline.Context) error {
	ctxAware := httppipeline.CreateContextAware(ctx)

	httpReader, err := ctxAware.GetReader()

	if err != nil {
		return err
	}

	code, exists := httpReader.GetFormData("code")

	if !exists {
		return http.CreateBadRequestError("'code' part not found")
	}

	ctx.Set(TagCodeRaw, code)
	return nil
}

func CreateFetchCodeCommand() pipeline.Step {
	return &fetchCodeCommand{}
}
