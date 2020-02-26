package pipeline

import (
	"github.com/arpb2/C-3PO/pkg/domain/infrastructure/http"
	api "github.com/arpb2/C-3PO/pkg/domain/infrastructure/pipeline"
	gopipeline "github.com/saantiaguilera/go-pipeline"
)

func CreateDebugHttpExecutor() gopipeline.Executor {
	return &httpDebugExecutor{}
}

type httpDebugExecutor struct{}

func (e *httpDebugExecutor) Run(runnable gopipeline.Runnable, ctx gopipeline.Context) error {
	return runnable.Run(ctx)
}

func CreateDebugHttpPipeline() api.HttpPipeline {
	pipe := gopipeline.CreatePipeline(CreateDebugHttpExecutor())

	return &httpPipeline{
		Pipeline: pipe,
	}
}

type httpPipeline struct {
	Pipeline gopipeline.Pipeline
}

func (h *httpPipeline) Run(context *http.Context, stage gopipeline.Stage) {
	ctx := gopipeline.CreateContext()
	ctx.Set(api.TagHttpReader, context)
	ctx.Set(api.TagHttpWriter, context)
	ctx.Set(api.TagHttpMiddleware, context)

	err := h.Pipeline.Run(stage, ctx)

	if err != nil {
		context.AbortTransactionWithError(err)
	}
}
