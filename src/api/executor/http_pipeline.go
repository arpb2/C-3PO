package executor

import (
	"github.com/arpb2/C-3PO/src/api/http_wrapper"
	"github.com/saantiaguilera/go-pipeline/pkg"
	"github.com/saantiaguilera/go-pipeline/pkg/pipeline"
)

type HttpPipeline interface {

	Run(context *http_wrapper.Context, stage pkg.Stage)

}

func CreatePipeline(executor pkg.Executor) HttpPipeline {
	pipe := pipeline.CreatePipeline(executor)

	return &httpPipeline{
		Pipeline: pipe,
	}
}

type httpPipeline struct{
	Pipeline pkg.Pipeline
}

func (h *httpPipeline) Run(context *http_wrapper.Context, stage pkg.Stage) {
	err := h.Pipeline.Run(stage)

	if err != nil {
		context.AbortTransactionWithError(err)
	}
}
