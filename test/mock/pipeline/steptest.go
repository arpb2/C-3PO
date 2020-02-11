package pipeline

import (
	gopipeline "github.com/saantiaguilera/go-pipeline"
	"github.com/stretchr/testify/mock"
)

type MockStep struct {
	mock.Mock
}

func (m MockStep) Name() string {
	args := m.Called()
	return args.String(0)
}

func (m MockStep) Run(ctx gopipeline.Context) error {
	args := m.Called()
	return args.Error(0)
}
