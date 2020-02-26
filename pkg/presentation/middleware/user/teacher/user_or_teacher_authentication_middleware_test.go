package teacher_test

import (
	"errors"
	"github.com/arpb2/C-3PO/pkg/domain/session/repository"
	"net/http"
	"testing"

	"github.com/arpb2/C-3PO/pkg/presentation/user"

	http3 "github.com/arpb2/C-3PO/pkg/domain/architecture/http"
	model2 "github.com/arpb2/C-3PO/pkg/domain/user/model"
	http2 "github.com/arpb2/C-3PO/test/mock/http"
	"github.com/arpb2/C-3PO/test/mock/token"

	"github.com/arpb2/C-3PO/pkg/presentation/middleware"

	"github.com/arpb2/C-3PO/pkg/presentation/middleware/user/teacher"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type MockTeacherRepository struct {
	mock.Mock
}

func (m *MockTeacherRepository) GetUser(userId uint) (user model2.User, err error) {
	panic("implement me")
}

func (m *MockTeacherRepository) CreateUser(user model2.User) (result model2.User, err error) {
	panic("implement me")
}

func (m *MockTeacherRepository) UpdateUser(user model2.User) (result model2.User, err error) {
	panic("implement me")
}

func (m *MockTeacherRepository) DeleteUser(userId uint) error {
	panic("implement me")
}

func (m *MockTeacherRepository) GetStudents(userId uint) (students []model2.User, err error) {
	args := m.Called(userId)

	firstArg := args.Get(0)
	if firstArg != nil {
		students = firstArg.([]model2.User)
	}

	err = args.Error(1)
	return
}

func Test_Multi_HandlingOfAuthentication_NoHeader(t *testing.T) {
	reader := new(http2.MockReader)
	reader.On("GetHeader", middleware.HeaderAuthorization).Return("")

	c, recorder := http2.CreateTestContext()
	c.Reader = reader

	middle := teacher.CreateMiddleware(&token.MockTokenHandler{}, &MockTeacherRepository{})

	middle(c)

	assert.Equal(t, http.StatusUnauthorized, recorder.Code)
	assert.Equal(t, "{\"error\":\"unauthorized\"}", recorder.Body.String())
}

func Test_Multi_HandlingOfAuthentication_BadHeader(t *testing.T) {
	tokenHandler := new(token.MockTokenHandler)
	tokenHandler.On("Retrieve", "bad token").Return(nil, http3.CreateBadRequestError("malformed token"))

	reader := new(http2.MockReader)
	reader.On("GetHeader", middleware.HeaderAuthorization).Return("bad token")

	c, recorder := http2.CreateTestContext()
	c.Reader = reader

	middle := teacher.CreateMiddleware(tokenHandler, &MockTeacherRepository{})

	middle(c)

	assert.Equal(t, http.StatusBadRequest, recorder.Code)
	assert.Equal(t, "{\"error\":\"malformed token\"}", recorder.Body.String())
	reader.AssertExpectations(t)
	tokenHandler.AssertExpectations(t)
}

func Test_Multi_HandlingOfAuthentication_UnauthorizedUser(t *testing.T) {
	tokenHandler := new(token.MockTokenHandler)
	tokenHandler.On("Retrieve", "token").Return(&repository.Token{
		UserId: 1000,
	}, nil)

	repository := new(MockTeacherRepository)
	repository.On("GetStudents", uint(1000)).Return(nil, nil).Once()

	reader := new(http2.MockReader)
	reader.On("GetParameter", user.ParamUserId).Return("1", nil).Once()
	reader.On("GetHeader", middleware.HeaderAuthorization).Return("token")

	c, recorder := http2.CreateTestContext()
	c.Reader = reader

	middle := teacher.CreateMiddleware(tokenHandler, repository)

	middle(c)

	assert.Equal(t, http.StatusUnauthorized, recorder.Code)
	assert.Equal(t, "{\"error\":\"unauthorized\"}", recorder.Body.String())
	repository.AssertExpectations(t)
	reader.AssertExpectations(t)
	tokenHandler.AssertExpectations(t)
}

func Test_Multi_HandlingOfAuthentication_Authorized_SameUser(t *testing.T) {
	tokenHandler := new(token.MockTokenHandler)
	tokenHandler.On("Retrieve", "token").Return(&repository.Token{
		UserId: 1000,
	}, nil)

	reader := new(http2.MockReader)
	reader.On("GetParameter", user.ParamUserId).Return("1000", nil).Once()
	reader.On("GetHeader", middleware.HeaderAuthorization).Return("token")

	c, recorder := http2.CreateTestContext()
	c.Reader = reader

	middle := teacher.CreateMiddleware(tokenHandler, &MockTeacherRepository{})

	middle(c)

	assert.Equal(t, http.StatusOK, recorder.Code)
	reader.AssertExpectations(t)
	tokenHandler.AssertExpectations(t)
}

func Test_Multi_HandlingOfAuthentication_Authorized_Student(t *testing.T) {
	tokenHandler := new(token.MockTokenHandler)
	tokenHandler.On("Retrieve", "token").Return(&repository.Token{
		UserId: 1001,
	}, nil)

	repository := new(MockTeacherRepository)
	repository.On("GetStudents", uint(1001)).Return([]model2.User{
		{
			Id: 999,
		},
		{
			Id: 1000,
		},
	}, nil).Once()

	reader := new(http2.MockReader)
	reader.On("GetParameter", user.ParamUserId).Return("1000", nil).Once()
	reader.On("GetHeader", middleware.HeaderAuthorization).Return("token")

	c, recorder := http2.CreateTestContext()
	c.Reader = reader

	middle := teacher.CreateMiddleware(tokenHandler, repository)

	middle(c)

	assert.Equal(t, http.StatusOK, recorder.Code)
	repository.AssertExpectations(t)
	reader.AssertExpectations(t)
	tokenHandler.AssertExpectations(t)
}

func Test_Multi_HandlingOfAuthentication_Unauthorized_Student(t *testing.T) {
	tokenHandler := new(token.MockTokenHandler)
	tokenHandler.On("Retrieve", "token").Return(&repository.Token{
		UserId: 1002,
	}, nil)

	repository := new(MockTeacherRepository)
	repository.On("GetStudents", uint(1002)).Return([]model2.User{
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

	reader := new(http2.MockReader)
	reader.On("GetParameter", user.ParamUserId).Return("1000", nil).Once()
	reader.On("GetHeader", middleware.HeaderAuthorization).Return("token")

	c, recorder := http2.CreateTestContext()
	c.Reader = reader

	middle := teacher.CreateMiddleware(tokenHandler, repository)

	middle(c)

	assert.Equal(t, http.StatusUnauthorized, recorder.Code)
	assert.Equal(t, "{\"error\":\"unauthorized\"}", recorder.Body.String())
	repository.AssertExpectations(t)
	reader.AssertExpectations(t)
	tokenHandler.AssertExpectations(t)
}

func Test_Multi_HandlingOfAuthentication_Repository_Error(t *testing.T) {
	tokenHandler := new(token.MockTokenHandler)
	tokenHandler.On("Retrieve", "token").Return(&repository.Token{
		UserId: 1001,
	}, nil)

	repository := new(MockTeacherRepository)
	repository.On("GetStudents", uint(1001)).Return([]model2.User{}, errors.New("whoops this fails")).Once()

	reader := new(http2.MockReader)
	reader.On("GetParameter", user.ParamUserId).Return("1000", nil).Once()
	reader.On("GetHeader", middleware.HeaderAuthorization).Return("token")

	c, recorder := http2.CreateTestContext()
	c.Reader = reader

	middle := teacher.CreateMiddleware(tokenHandler, repository)

	middle(c)

	assert.Equal(t, http.StatusInternalServerError, recorder.Code)
	assert.Equal(t, "{\"error\":\"internal error\"}", recorder.Body.String())
	repository.AssertExpectations(t)
	reader.AssertExpectations(t)
	tokenHandler.AssertExpectations(t)
}
