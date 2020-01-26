package session_task_test

import (
	"errors"
	"github.com/arpb2/C-3PO/src/api/controller/session/session_task"
	"github.com/arpb2/C-3PO/src/api/http_wrapper"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"testing"
)

func TestFetchUserTaskImpl_FailsOnReadError(t *testing.T) {
	reader := new(http_wrapper.TestReader)
	reader.On("ReadBody", mock.MatchedBy(func(obj interface{}) bool {
		return true
	})).Return(errors.New("error")).Once()

	user, err := session_task.CreateFetchUserTask()(
		&http_wrapper.Context{
			Reader:     reader,
			Writer:     nil,
			Middleware: nil,
		},
	)

	assert.Nil(t, user)
	assert.NotNil(t, err)
	reader.AssertExpectations(t)
}

func TestFetchUserIdTaskImpl_ReturnsUser(t *testing.T) {
	reader := new(http_wrapper.TestReader)
	reader.On("ReadBody", mock.MatchedBy(func(obj interface{}) bool {
		return true
	})).Return(nil).Once()

	user, err := session_task.CreateFetchUserTask()(
		&http_wrapper.Context{
			Reader:     reader,
			Writer:     &http_wrapper.TestWriter{},
			Middleware: &http_wrapper.TestMiddleware{},
		},
	)

	assert.NotNil(t, user)
	assert.Nil(t, err)
	reader.AssertExpectations(t)
}
