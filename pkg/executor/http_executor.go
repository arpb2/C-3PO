package executor

import (
	"net/http"

	"github.com/afex/hystrix-go/hystrix"
	"github.com/arpb2/C-3PO/api/http_wrapper"
	go_pipeline "github.com/saantiaguilera/go-pipeline/pkg/api"
	"github.com/saantiaguilera/go-pipeline/pkg/step/trace"
	"golang.org/x/xerrors"
)

func CreateHttpExecutor() go_pipeline.Executor {
	return &httpExecutor{}
}

type httpExecutor struct{}

func (e *httpExecutor) Run(runnable go_pipeline.Runnable) error {
	runnable = trace.CreateTracedStep(runnable)

	var err error
	_ = hystrix.Do(runnable.Name(), func() error {
		err = runnable.Run()

		var httpError http_wrapper.HttpError
		if xerrors.As(err, &httpError) && httpError.Code < http.StatusInternalServerError {
			return nil
		}
		return err
	}, nil)

	return err
}

func CreateDebugHttpExecutor() go_pipeline.Executor {
	return &httpDebugExecutor{}
}

type httpDebugExecutor struct{}

func (e *httpDebugExecutor) Run(runnable go_pipeline.Runnable) error {
	return runnable.Run()
}
