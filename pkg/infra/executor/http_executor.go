package executor

import (
	"net/http"

	"github.com/arpb2/C-3PO/pkg/domain/executor/decorator"

	"github.com/afex/hystrix-go/hystrix"
	httpwrapper "github.com/arpb2/C-3PO/pkg/domain/http"
	gopipeline "github.com/saantiaguilera/go-pipeline"
	"golang.org/x/xerrors"
)

func CreateHttpExecutor(decorators ...decorator.RunnableDecorator) gopipeline.Executor {
	return CreateCircuitBreakerHttpExecutor(func(name string, run func() error) error {
		return hystrix.Do(name, run, nil)
	}, decorators...)
}

type CircuitBreaker func(name string, run func() error) error

func CreateCircuitBreakerHttpExecutor(cb CircuitBreaker, decorators ...decorator.RunnableDecorator) gopipeline.Executor {
	return &httpExecutor{
		Decorators:     decorators,
		CircuitBreaker: cb,
	}
}

type httpExecutor struct {
	Decorators     []decorator.RunnableDecorator
	CircuitBreaker CircuitBreaker
}

func (e *httpExecutor) Run(runnable gopipeline.Runnable, ctx gopipeline.Context) error {
	for _, d := range e.Decorators {
		runnable = d(runnable)
	}

	var err error
	_ = e.CircuitBreaker(runnable.Name(), func() error {
		err = runnable.Run(ctx)

		var httpError httpwrapper.Error
		if xerrors.As(err, &httpError) && httpError.Code < http.StatusInternalServerError {
			return nil
		}
		return err
	})

	return err
}

func CreateDebugHttpExecutor() gopipeline.Executor {
	return &httpDebugExecutor{}
}

type httpDebugExecutor struct{}

func (e *httpDebugExecutor) Run(runnable gopipeline.Runnable, ctx gopipeline.Context) error {
	return runnable.Run(ctx)
}
