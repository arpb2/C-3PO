package decorator

import (
	"io"

	"github.com/arpb2/C-3PO/pkg/domain/executor/decorator"
	"github.com/saantiaguilera/go-pipeline"
)

func CreateTraceDecorator(writer io.Writer) decorator.RunnableDecorator {
	return func(runnable pipeline.Runnable) pipeline.Runnable {
		return pipeline.CreateTracedStepWithWriter(runnable, writer)
	}
}
