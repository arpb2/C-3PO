package session_test

import (
	"errors"
	"testing"

	session2 "github.com/arpb2/C-3PO/pkg/data/repository/session"
	"github.com/arpb2/C-3PO/pkg/data/usecase/session"
	user2 "github.com/arpb2/C-3PO/pkg/domain/model/user"

	http3 "github.com/arpb2/C-3PO/pkg/domain/http"
	"github.com/arpb2/C-3PO/test/mock/token"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type MockTeacherRepository struct {
	mock.Mock
}

func (m *MockTeacherRepository) GetUser(userId uint) (user user2.User, err error) {
	panic("implement me")
}

func (m *MockTeacherRepository) CreateUser(user user2.User) (result user2.User, err error) {
	panic("implement me")
}

func (m *MockTeacherRepository) UpdateUser(user user2.User) (result user2.User, err error) {
	panic("implement me")
}

func (m *MockTeacherRepository) DeleteUser(userId uint) error {
	panic("implement me")
}

func (m *MockTeacherRepository) GetStudents(userId uint) (students []user2.User, err error) {
	args := m.Called(userId)

	firstArg := args.Get(0)
	if firstArg != nil {
		students = firstArg.([]user2.User)
	}

	err = args.Error(1)
	return
}

func Test_Multi_HandlingOfAuthentication_NoHeader(t *testing.T) {
	middle := session.CreateTeacherAuthenticationUseCase(&token.MockTokenHandler{}, &MockTeacherRepository{})

	err := middle("", "")

	assert.Equal(t, http3.CreateUnauthorizedError(), err)
}

func Test_Multi_HandlingOfAuthentication_BadHeader(t *testing.T) {
	tokenHandler := new(token.MockTokenHandler)
	tokenHandler.On("Retrieve", "bad token").Return(nil, http3.CreateBadRequestError("malformed token"))

	middle := session.CreateTeacherAuthenticationUseCase(tokenHandler, &MockTeacherRepository{})

	err := middle("bad token", "")

	assert.Equal(t, http3.CreateBadRequestError("malformed token"), err)
	tokenHandler.AssertExpectations(t)
}

func Test_Multi_HandlingOfAuthentication_UnauthorizedUser(t *testing.T) {
	tokenHandler := new(token.MockTokenHandler)
	tokenHandler.On("Retrieve", "token").Return(&session2.Token{
		UserId: 1000,
	}, nil)

	repository := new(MockTeacherRepository)
	repository.On("GetStudents", uint(1000)).Return(nil, nil).Once()

	middle := session.CreateTeacherAuthenticationUseCase(tokenHandler, repository)

	err := middle("token", "1")

	assert.Equal(t, http3.CreateUnauthorizedError(), err)
	repository.AssertExpectations(t)
	tokenHandler.AssertExpectations(t)
}

func Test_Multi_HandlingOfAuthentication_Authorized_SameUser(t *testing.T) {
	tokenHandler := new(token.MockTokenHandler)
	tokenHandler.On("Retrieve", "token").Return(&session2.Token{
		UserId: 1000,
	}, nil)

	middle := session.CreateTeacherAuthenticationUseCase(tokenHandler, &MockTeacherRepository{})

	err := middle("token", "1000")

	assert.Nil(t, err)
	tokenHandler.AssertExpectations(t)
}

func Test_Multi_HandlingOfAuthentication_Authorized_Student(t *testing.T) {
	tokenHandler := new(token.MockTokenHandler)
	tokenHandler.On("Retrieve", "token").Return(&session2.Token{
		UserId: 1001,
	}, nil)

	repository := new(MockTeacherRepository)
	repository.On("GetStudents", uint(1001)).Return([]user2.User{
		{
			Id: 999,
		},
		{
			Id: 1000,
		},
	}, nil).Once()

	middle := session.CreateTeacherAuthenticationUseCase(tokenHandler, repository)

	err := middle("token", "1000")

	assert.Nil(t, err)
	repository.AssertExpectations(t)
	tokenHandler.AssertExpectations(t)
}

func Test_Multi_HandlingOfAuthentication_Unauthorized_Student(t *testing.T) {
	tokenHandler := new(token.MockTokenHandler)
	tokenHandler.On("Retrieve", "token").Return(&session2.Token{
		UserId: 1002,
	}, nil)

	repository := new(MockTeacherRepository)
	repository.On("GetStudents", uint(1002)).Return([]user2.User{
		{
			Id: 1,
		},
		{
			Id: 2,
		},
		{
			Id: 3,
		},
	}, nil).Once()

	middle := session.CreateTeacherAuthenticationUseCase(tokenHandler, repository)

	err := middle("token", "1000")

	assert.Equal(t, http3.CreateUnauthorizedError(), err)
	repository.AssertExpectations(t)
	tokenHandler.AssertExpectations(t)
}

func Test_Multi_HandlingOfAuthentication_Repository_Error(t *testing.T) {
	tokenHandler := new(token.MockTokenHandler)
	tokenHandler.On("Retrieve", "token").Return(&session2.Token{
		UserId: 1001,
	}, nil)

	expectedErr := errors.New("whoops this fails")
	repository := new(MockTeacherRepository)
	repository.On("GetStudents", uint(1001)).Return([]user2.User{}, expectedErr).Once()

	middle := session.CreateTeacherAuthenticationUseCase(tokenHandler, repository)

	err := middle("token", "1000")

	assert.Equal(t, expectedErr, err)
	repository.AssertExpectations(t)
	tokenHandler.AssertExpectations(t)
}
