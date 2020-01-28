package decorator

import (
	"github.com/saantiaguilera/go-pipeline"
)

func TraceRunnableDecorator(runnable pipeline.Runnable) pipeline.Runnable {
	return pipeline.CreateTracedStep(runnable)
}
