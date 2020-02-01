package pipeline

import (
	"fmt"
	"io"

	"github.com/arpb2/C-3PO/api/http"
	api "github.com/arpb2/C-3PO/api/pipeline"
	"github.com/saantiaguilera/go-pipeline"
)

func CreateDrawablePipeline(writer io.WriteCloser) api.HttpPipeline {
	return &drawablePipeline{
		Writer: writer,
	}
}

type drawablePipeline struct {
	Writer io.WriteCloser
}

func (h *drawablePipeline) Run(context *http.Context, stage pipeline.Stage) {
	graphDiagram := pipeline.CreateUMLActivityGraphDiagram()
	graphRenderer := pipeline.CreateUMLActivityRenderer(pipeline.UMLOptions{
		Type: pipeline.UMLFormatSVG,
	})

	stage.Draw(graphDiagram)

	err := graphRenderer.Render(graphDiagram, h.Writer)

	if err != nil {
		fmt.Print(err.Error())
	}
}
