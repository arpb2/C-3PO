package executor

import (
	"net/http"

	"github.com/arpb2/C-3PO/api/executor/decorator"

	"github.com/afex/hystrix-go/hystrix"
	httpwrapper "github.com/arpb2/C-3PO/api/http"
	gopipeline "github.com/saantiaguilera/go-pipeline"
	"golang.org/x/xerrors"
)

func CreateHttpExecutor(decorators ...decorator.RunnableDecorator) gopipeline.Executor {
	return &httpExecutor{
		Decorators: decorators,
	}
}

type httpExecutor struct {
	Decorators []decorator.RunnableDecorator
}

func (e *httpExecutor) Run(runnable gopipeline.Runnable, ctx gopipeline.Context) error {
	for _, decorator := range e.Decorators {
		runnable = decorator(runnable)
	}

	var err error
	_ = hystrix.Do(runnable.Name(), func() error {
		err = runnable.Run(ctx)

		var httpError httpwrapper.Error
		if xerrors.As(err, &httpError) && httpError.Code < http.StatusInternalServerError {
			return nil
		}
		return err
	}, nil)

	return err
}

func CreateDebugHttpExecutor() gopipeline.Executor {
	return &httpDebugExecutor{}
}

type httpDebugExecutor struct{}

func (e *httpDebugExecutor) Run(runnable gopipeline.Runnable, ctx gopipeline.Context) error {
	return runnable.Run(ctx)
}
