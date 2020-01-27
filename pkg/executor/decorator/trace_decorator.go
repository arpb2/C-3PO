package decorator

import (
	"github.com/saantiaguilera/go-pipeline/pkg/api"
	"github.com/saantiaguilera/go-pipeline/pkg/step/trace"
)

func TraceRunnableDecorator(runnable api.Runnable) api.Runnable {
	return trace.CreateTracedStep(runnable)
}
