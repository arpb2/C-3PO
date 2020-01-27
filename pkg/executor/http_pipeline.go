package executor

import (
	"github.com/arpb2/C-3PO/api/http"
	api "github.com/arpb2/C-3PO/api/pipeline"
	gopipeline "github.com/saantiaguilera/go-pipeline/pkg/api"
	"github.com/saantiaguilera/go-pipeline/pkg/pipeline"
)

func CreatePipeline(executor gopipeline.Executor) api.HttpPipeline {
	pipe := pipeline.CreatePipeline(executor)

	return &httpPipeline{
		Pipeline: pipe,
	}
}

type httpPipeline struct {
	Pipeline gopipeline.Pipeline
}

func (h *httpPipeline) Run(context *http.Context, stage gopipeline.Stage) {
	err := h.Pipeline.Run(stage)

	if err != nil {
		context.AbortTransactionWithError(err)
	}
}
