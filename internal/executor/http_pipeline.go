package executor

import (
	"github.com/arpb2/C-3PO/api/http_wrapper"
	api "github.com/arpb2/C-3PO/api/pipeline"
	"github.com/saantiaguilera/go-pipeline/pkg"
	"github.com/saantiaguilera/go-pipeline/pkg/pipeline"
)

func CreatePipeline(executor pkg.Executor) api.HttpPipeline {
	pipe := pipeline.CreatePipeline(executor)

	return &httpPipeline{
		Pipeline: pipe,
	}
}

type httpPipeline struct {
	Pipeline pkg.Pipeline
}

func (h *httpPipeline) Run(context *http_wrapper.Context, stage pkg.Stage) {
	err := h.Pipeline.Run(stage)

	if err != nil {
		context.AbortTransactionWithError(err)
	}
}
