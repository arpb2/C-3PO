package user_test

import (
	"bytes"
	"errors"
	"net/http"
	"testing"

	user2 "github.com/arpb2/C-3PO/pkg/domain/model/user"
	"github.com/arpb2/C-3PO/pkg/presentation/user"

	pipeline2 "github.com/arpb2/C-3PO/test/mock/pipeline"

	http2 "github.com/arpb2/C-3PO/pkg/domain/http"

	"github.com/arpb2/C-3PO/test/mock/golden"
	testhttpwrapper "github.com/arpb2/C-3PO/test/mock/http"
	"github.com/arpb2/C-3PO/test/mock/repository"
	"github.com/stretchr/testify/assert"
)

func createGetUserHandler() http2.Handler {
	return user.CreateGetUserHandler(
		"user_id",
		pipeline2.CreateDebugHttpPipeline(),
		nil,
	)
}

func TestUserGetControllerBody_400OnEmptyUserId(t *testing.T) {
	reader := new(testhttpwrapper.MockReader)
	reader.On("GetParameter", "user_id").Return("").Once()

	c, w := testhttpwrapper.CreateTestContext()
	c.Reader = reader

	createGetUserHandler()(c)
	actual := bytes.TrimSpace([]byte(w.Body.String()))
	expected := golden.Get(t, actual, "bad_request.empty.user_id.golden.json")

	assert.Equal(t, http.StatusBadRequest, w.Code)
	assert.Equal(t, expected, actual)
	reader.AssertExpectations(t)
}

func TestUserGetControllerBody_400OnMalformedUserId(t *testing.T) {
	reader := new(testhttpwrapper.MockReader)
	reader.On("GetParameter", "user_id").Return("not a number").Once()

	c, w := testhttpwrapper.CreateTestContext()
	c.Reader = reader

	createGetUserHandler()(c)
	actual := bytes.TrimSpace([]byte(w.Body.String()))
	expected := golden.Get(t, actual, "bad_request.malformed.user_id.golden.json")

	assert.Equal(t, http.StatusBadRequest, w.Code)
	assert.Equal(t, expected, actual)
	reader.AssertExpectations(t)
}

func TestUserGetControllerBody_500OnRepositoryReadError(t *testing.T) {
	repository := new(repository.MockUserRepository)
	repository.On("GetUser", uint(1000)).Return(nil, errors.New("whoops error")).Once()

	body := user.CreateGetUserHandler("user_id", pipeline2.CreateDebugHttpPipeline(), repository)

	reader := new(testhttpwrapper.MockReader)
	reader.On("GetParameter", "user_id").Return("1000").Once()

	c, w := testhttpwrapper.CreateTestContext()
	c.Reader = reader

	body(c)

	actual := bytes.TrimSpace([]byte(w.Body.String()))
	expected := golden.Get(t, actual, "internal_server_error.error_read.repository.golden.json")

	assert.Equal(t, http.StatusInternalServerError, w.Code)
	assert.Equal(t, expected, actual)
	reader.AssertExpectations(t)
	repository.AssertExpectations(t)
}

func TestUserGetControllerBody_400OnNoUserStoredInRepository(t *testing.T) {
	repository := new(repository.MockUserRepository)
	repository.On("GetUser", uint(1000)).Return(user2.User{}, http2.CreateNotFoundError()).Once()

	body := user.CreateGetUserHandler("user_id", pipeline2.CreateDebugHttpPipeline(), repository)

	reader := new(testhttpwrapper.MockReader)
	reader.On("GetParameter", "user_id").Return("1000").Once()

	c, w := testhttpwrapper.CreateTestContext()
	c.Reader = reader

	body(c)

	actual := bytes.TrimSpace([]byte(w.Body.String()))
	expected := golden.Get(t, actual, "not_found.missing_user.read.repository.golden.json")

	assert.Equal(t, http.StatusNotFound, w.Code)
	assert.Equal(t, expected, actual)
	reader.AssertExpectations(t)
	repository.AssertExpectations(t)
}

func TestUserGetControllerBody_200OnUserStoredOnRepository(t *testing.T) {
	expectedUser := user2.User{
		Id:      1000,
		Type:    user2.TypeStudent,
		Email:   "test@email.com",
		Name:    "TestName",
		Surname: "TestSurname",
	}
	repository := new(repository.MockUserRepository)
	repository.On("GetUser", uint(1000)).Return(expectedUser, nil).Once()

	body := user.CreateGetUserHandler("user_id", pipeline2.CreateDebugHttpPipeline(), repository)

	reader := new(testhttpwrapper.MockReader)
	reader.On("GetParameter", "user_id").Return("1000").Once()

	c, w := testhttpwrapper.CreateTestContext()
	c.Reader = reader

	body(c)

	actual := bytes.TrimSpace([]byte(w.Body.String()))
	expected := golden.Get(t, actual, "ok.get_user.golden.json")

	assert.Equal(t, http.StatusOK, w.Code)
	assert.Equal(t, expected, actual)
	reader.AssertExpectations(t)
	repository.AssertExpectations(t)
}
