package executor

import (
	"github.com/arpb2/C-3PO/api/http_wrapper"
	api "github.com/arpb2/C-3PO/api/pipeline"
	go_pipeline "github.com/saantiaguilera/go-pipeline/pkg/api"
	"github.com/saantiaguilera/go-pipeline/pkg/pipeline"
)

func CreatePipeline(executor go_pipeline.Executor) api.HttpPipeline {
	pipe := pipeline.CreatePipeline(executor)

	return &httpPipeline{
		Pipeline: pipe,
	}
}

type httpPipeline struct {
	Pipeline go_pipeline.Pipeline
}

func (h *httpPipeline) Run(context *http_wrapper.Context, stage go_pipeline.Stage) {
	err := h.Pipeline.Run(stage)

	if err != nil {
		context.AbortTransactionWithError(err)
	}
}
