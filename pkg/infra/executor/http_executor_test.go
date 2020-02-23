package executor_test

import (
	"errors"
	"testing"

	"github.com/arpb2/C-3PO/pkg/domain/http"
	"github.com/arpb2/C-3PO/pkg/infra/executor"
	pipeline2 "github.com/arpb2/C-3PO/test/mock/pipeline"
	"github.com/saantiaguilera/go-pipeline"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestHttpExecutor_GivenOneDecorated_WhenRunning_ItAppliesDecorators(t *testing.T) {
	mockStep := new(pipeline2.MockStep)
	mockStep.On("Name").Return("name")
	mockStep.On("Run", mock.Anything).Return(nil)
	decorator := func(runnable pipeline.Runnable) pipeline.Runnable {
		return mockStep
	}
	exec := executor.CreateCircuitBreakerHttpExecutor(func(name string, run func() error) error {
		return run()
	}, decorator)

	err := exec.Run(pipeline2.MockStep{}, nil)

	assert.Nil(t, err)
	mockStep.AssertExpectations(t)
}

func TestHttpExecutor_GivenOne_WhenRunningAndErroringAnonymously_ThenCircuitBreakerGetsIt(t *testing.T) {
	expectedErr := errors.New("anonymous error")
	mockStep := new(pipeline2.MockStep)
	mockStep.On("Name").Return("name")
	mockStep.On("Run", mock.Anything).Return(expectedErr)
	decorator := func(runnable pipeline.Runnable) pipeline.Runnable {
		return mockStep
	}
	var cbErr error
	exec := executor.CreateCircuitBreakerHttpExecutor(func(name string, run func() error) error {
		cbErr = run()
		return cbErr
	}, decorator)

	err := exec.Run(pipeline2.MockStep{}, nil)

	assert.Equal(t, expectedErr, err)
	assert.Equal(t, expectedErr, cbErr)
	mockStep.AssertExpectations(t)
}

func TestHttpExecutor_GivenOne_WhenRunningAndErroring5xx_ThenCircuitBreakerGetsIt(t *testing.T) {
	expectedErr := http.CreateInternalError()
	mockStep := new(pipeline2.MockStep)
	mockStep.On("Name").Return("name")
	mockStep.On("Run", mock.Anything).Return(expectedErr)
	decorator := func(runnable pipeline.Runnable) pipeline.Runnable {
		return mockStep
	}
	var cbErr error
	exec := executor.CreateCircuitBreakerHttpExecutor(func(name string, run func() error) error {
		cbErr = run()
		return cbErr
	}, decorator)

	err := exec.Run(pipeline2.MockStep{}, nil)

	assert.Equal(t, expectedErr, err)
	assert.Equal(t, expectedErr, cbErr)
	mockStep.AssertExpectations(t)
}

func TestHttpExecutor_GivenOne_WhenRunningAndErroringWith4xx_ThenCircuitBreakerGetsNil(t *testing.T) {
	expectedErr := http.CreateBadRequestError("client error")
	mockStep := new(pipeline2.MockStep)
	mockStep.On("Name").Return("name")
	mockStep.On("Run", mock.Anything).Return(expectedErr)
	decorator := func(runnable pipeline.Runnable) pipeline.Runnable {
		return mockStep
	}
	var cbErr error
	exec := executor.CreateCircuitBreakerHttpExecutor(func(name string, run func() error) error {
		cbErr = run()
		return cbErr
	}, decorator)

	err := exec.Run(pipeline2.MockStep{}, nil)

	assert.Equal(t, expectedErr, err)
	assert.Nil(t, cbErr)
	mockStep.AssertExpectations(t)
}
