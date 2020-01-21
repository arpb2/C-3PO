package user_test

import (
	"github.com/arpb2/C-3PO/src/api/controller/user"
	"github.com/arpb2/C-3PO/src/api/http_wrapper"
	"github.com/arpb2/C-3PO/src/api/http_wrapper/gin_wrapper"
	"github.com/arpb2/C-3PO/src/api/model"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"net/http"
	"testing"
)

type MockUserService struct {
	mock.Mock
}

func (m MockUserService) GetUser(userId uint) (user *model.User, err error) {
	args := m.Called(userId)

	firstParam := args.Get(0)
	if firstParam != nil {
		user = firstParam.(*model.User)
	}

	err = args.Error(1)
	return
}

func (m MockUserService) CreateUser(authenticatedUser *model.AuthenticatedUser) (user *model.User, err error) {
	args := m.Called(authenticatedUser)

	firstParam := args.Get(0)
	if firstParam != nil {
		user = firstParam.(*model.AuthenticatedUser).User
	}

	err = args.Error(1)
	return
}

func (m MockUserService) UpdateUser(authenticatedUser *model.AuthenticatedUser) (user *model.User, err error) {
	args := m.Called(authenticatedUser)

	firstParam := args.Get(0)
	if firstParam != nil {
		user = firstParam.(*model.User)
	}

	err = args.Error(1)
	return
}

func (m MockUserService) DeleteUser(userId uint) error {
	args := m.Called(userId)
	return args.Error(0)
}

func TestFetchUserId_RetrievesFromParam(t *testing.T) {
	reader := new(http_wrapper.TestReader)
	reader.On("Param", "user_id").Return("1234").Once()

	c, _ := gin_wrapper.CreateTestContext()
	c.Reader = reader

	userId, halt := user.FetchUserId(c)

	assert.False(t, halt)
	assert.Equal(t, uint(1234), userId)
	reader.AssertExpectations(t)
}

func TestFetchUserId_RetrievesFromParam_400IfMalformed(t *testing.T) {
	reader := new(http_wrapper.TestReader)
	reader.On("Param", "user_id").Return("not a number").Once()

	c, _ := gin_wrapper.CreateTestContext()
	c.Reader = reader

	userId, halt := user.FetchUserId(c)

	assert.True(t, halt)
	assert.Zero(t, userId)
	reader.AssertExpectations(t)
}

func TestFetchUserId_HaltsWith400OnError(t *testing.T) {
	c, recorder := gin_wrapper.CreateTestContext()

	userId, halt := user.FetchUserId(c)

	assert.True(t, halt)
	assert.Zero(t, userId)
	assert.Equal(t, http.StatusBadRequest, recorder.Code)
}

