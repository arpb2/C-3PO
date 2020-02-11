package pipeline

import (
	"io"

	gopipeline "github.com/saantiaguilera/go-pipeline"
	"github.com/stretchr/testify/mock"
)

type MockRenderer struct {
	mock.Mock
}

func (m MockRenderer) Render(graphDiagram gopipeline.GraphDiagram, output io.WriteCloser) error {
	args := m.Called(graphDiagram, output)
	return args.Error(0)
}
