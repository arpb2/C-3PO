package pipeline

import (
	gopipeline "github.com/saantiaguilera/go-pipeline"
	"github.com/stretchr/testify/mock"
)

type MockStage struct {
	mock.Mock
}

func (s *MockStage) Draw(graph gopipeline.GraphDiagram) {
	_ = s.Called(graph)
}

func (s *MockStage) Run(executor gopipeline.Executor, ctx gopipeline.Context) error {
	args := s.Called(executor, ctx)
	return args.Error(0)
}
