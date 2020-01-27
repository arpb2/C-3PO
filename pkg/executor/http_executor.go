package executor

import (
	"net/http"

	"github.com/afex/hystrix-go/hystrix"
	httpwrapper "github.com/arpb2/C-3PO/api/http"
	gopipeline "github.com/saantiaguilera/go-pipeline/pkg/api"
	"github.com/saantiaguilera/go-pipeline/pkg/step/trace"
	"golang.org/x/xerrors"
)

func CreateHttpExecutor() gopipeline.Executor {
	return &httpExecutor{}
}

type httpExecutor struct{}

func (e *httpExecutor) Run(runnable gopipeline.Runnable) error {
	runnable = trace.CreateTracedStep(runnable)

	var err error
	_ = hystrix.Do(runnable.Name(), func() error {
		err = runnable.Run()

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

func (e *httpDebugExecutor) Run(runnable gopipeline.Runnable) error {
	return runnable.Run()
}
