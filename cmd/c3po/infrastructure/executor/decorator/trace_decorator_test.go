package decorator_test

import (
	"bytes"
	"testing"

	"github.com/arpb2/C-3PO/cmd/c3po/infrastructure/executor/decorator"
	pipeline2 "github.com/arpb2/C-3PO/test/mock/pipeline"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestTraceRunnable_GivenADecorator_WhenRunning_ThenOutputsTrace(t *testing.T) {
	mockStep := new(pipeline2.MockStep)
	mockStep.On("Name").Return("test name")
	mockStep.On("Run", mock.Anything).Return(nil)
	writer := bytes.NewBufferString("")
	step := decorator.CreateTraceDecorator(writer)

	_ = step(mockStep).Run(nil)

	assert.True(t, writer.Len() > 0)
	mockStep.AssertExpectations(t)
}
