package pipeline

import (
	"io"

	"github.com/arpb2/C-3PO/pkg/domain/http"
	api "github.com/arpb2/C-3PO/pkg/domain/pipeline"
	"github.com/saantiaguilera/go-pipeline"
)

func CreateDrawablePipeline(writer io.WriteCloser, renderer pipeline.DiagramRenderer) api.HttpPipeline {
	return &drawablePipeline{
		Writer:   writer,
		Renderer: renderer,
	}
}

type drawablePipeline struct {
	Writer   io.WriteCloser
	Renderer pipeline.DiagramRenderer
}

func (h *drawablePipeline) Run(context *http.Context, stage pipeline.Stage) {
	graphDiagram := pipeline.CreateUMLActivityGraphDiagram()

	stage.Draw(graphDiagram)

	_ = h.Renderer.Render(graphDiagram, h.Writer)
}
