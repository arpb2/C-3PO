package pipeline_test

import (
	"testing"

	api "github.com/arpb2/C-3PO/pkg/domain/infrastructure/pipeline"

	"github.com/arpb2/C-3PO/cmd/c3po/infrastructure/pipeline"
	http2 "github.com/arpb2/C-3PO/pkg/domain/infrastructure/http"
	"github.com/arpb2/C-3PO/test/mock/http"
	pipeline2 "github.com/arpb2/C-3PO/test/mock/pipeline"
	gopipeline "github.com/saantiaguilera/go-pipeline"
	"github.com/stretchr/testify/mock"
)

func TestHttpPipeline_GivenOne_WhenRunning_ThenContextHasReader(t *testing.T) {
	pipe := pipeline.CreateHttpPipeline(nil)
	ctx := &http2.Context{
		Reader:     &http.MockReader{},
		Writer:     &http.MockWriter{},
		Middleware: &http.MockMiddleware{},
	}
	stage := new(pipeline2.MockStage)
	stage.On("Run", nil, mock.MatchedBy(func(obj gopipeline.Context) bool {
		value, exists := obj.Get(api.TagHttpReader)
		return exists && value != nil
	})).Return(nil)

	pipe.Run(ctx, stage)

	stage.AssertExpectations(t)
}

func TestHttpPipeline_GivenOne_WhenRunning_ThenContextHasWriter(t *testing.T) {
	pipe := pipeline.CreateHttpPipeline(nil)
	ctx := &http2.Context{
		Reader:     &http.MockReader{},
		Writer:     &http.MockWriter{},
		Middleware: &http.MockMiddleware{},
	}
	stage := new(pipeline2.MockStage)
	stage.On("Run", nil, mock.MatchedBy(func(obj gopipeline.Context) bool {
		value, exists := obj.Get(api.TagHttpWriter)
		return exists && value != nil
	})).Return(nil)

	pipe.Run(ctx, stage)

	stage.AssertExpectations(t)
}

func TestHttpPipeline_GivenOne_WhenRunning_ThenContextHasMiddleware(t *testing.T) {
	pipe := pipeline.CreateHttpPipeline(nil)
	ctx := &http2.Context{
		Reader:     &http.MockReader{},
		Writer:     &http.MockWriter{},
		Middleware: &http.MockMiddleware{},
	}
	stage := new(pipeline2.MockStage)
	stage.On("Run", nil, mock.MatchedBy(func(obj gopipeline.Context) bool {
		value, exists := obj.Get(api.TagHttpMiddleware)
		return exists && value != nil
	})).Return(nil)

	pipe.Run(ctx, stage)

	stage.AssertExpectations(t)
}

func TestHttpPipeline_GivenOne_WhenRunningAndFailing_ThenContextAborts(t *testing.T) {
	pipe := pipeline.CreateHttpPipeline(nil)
	err := http2.CreateInternalError()
	middleware := new(http.MockMiddleware)
	middleware.On("AbortTransactionWithError", err)
	stage := new(pipeline2.MockStage)
	stage.On("Run", nil, mock.Anything).Return(err)
	ctx := &http2.Context{
		Reader:     &http.MockReader{},
		Writer:     &http.MockWriter{},
		Middleware: middleware,
	}

	pipe.Run(ctx, stage)

	stage.AssertExpectations(t)
	middleware.AssertExpectations(t)
}
