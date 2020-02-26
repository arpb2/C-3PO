package pipeline

import (
	"github.com/arpb2/C-3PO/pkg/domain/architecture/http"
	api "github.com/arpb2/C-3PO/pkg/domain/architecture/pipeline"
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
	ctx := pipeline.CreateContext()
	ctx.Set(api.TagHttpReader, context)
	ctx.Set(api.TagHttpWriter, context)
	ctx.Set(api.TagHttpMiddleware, context)

	err := h.Pipeline.Run(stage, ctx)

	if err != nil {
		context.AbortTransactionWithError(err)
	}
}
