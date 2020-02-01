package pipeline

import (
	"github.com/arpb2/C-3PO/api/http"
	api "github.com/arpb2/C-3PO/api/pipeline"
	"github.com/saantiaguilera/go-pipeline"
)

func CreateHttpPipeline(executor pipeline.Executor) api.HttpPipeline {
	pipe := pipeline.CreatePipeline(executor)

	return &httpPipeline{
		Pipeline: pipe,
	}
}

type httpPipeline struct {
	Pipeline pipeline.Pipeline
}

func (h *httpPipeline) Run(context *http.Context, stage pipeline.Stage) {
	err := h.Pipeline.Run(stage)

	if err != nil {
		context.AbortTransactionWithError(err)
	}
}
