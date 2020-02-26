package teacher_test

import (
	"errors"
	"net/http"
	"testing"

	http3 "github.com/arpb2/C-3PO/pkg/domain/infrastructure/http"
	token2 "github.com/arpb2/C-3PO/pkg/domain/session/token"
	"github.com/arpb2/C-3PO/pkg/domain/user/controller"
	model2 "github.com/arpb2/C-3PO/pkg/domain/user/model"
	http2 "github.com/arpb2/C-3PO/test/mock/http"
	"github.com/arpb2/C-3PO/test/mock/token"

	"github.com/arpb2/C-3PO/pkg/presentation/middleware"

	"github.com/arpb2/C-3PO/pkg/presentation/middleware/user/teacher"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type MockTeacherService struct {
	mock.Mock
}

func (m *MockTeacherService) GetUser(userId uint) (user model2.User, err error) {
	panic("implement me")
}

func (m *MockTeacherService) CreateUser(user model2.User) (result model2.User, err error) {
	panic("implement me")
}

func (m *MockTeacherService) UpdateUser(user model2.User) (result model2.User, err error) {
	panic("implement me")
}

func (m *MockTeacherService) DeleteUser(userId uint) error {
	panic("implement me")
}

func (m *MockTeacherService) GetStudents(userId uint) (students []model2.User, err error) {
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

	middle := teacher.CreateMiddleware(&token.MockTokenHandler{}, &MockTeacherService{})

	middle(c)

	assert.Equal(t, http.StatusUnauthorized, recorder.Code)
	assert.Equal(t, "{\"error\":\"unauthorized\"}\n", recorder.Body.String())
}

func Test_Multi_HandlingOfAuthentication_BadHeader(t *testing.T) {
	tokenHandler := new(token.MockTokenHandler)
	tokenHandler.On("Retrieve", "bad token").Return(nil, http3.CreateBadRequestError("malformed token"))

	reader := new(http2.MockReader)
	reader.On("GetHeader", middleware.HeaderAuthorization).Return("bad token")

	c, recorder := http2.CreateTestContext()
	c.Reader = reader

	middle := teacher.CreateMiddleware(tokenHandler, &MockTeacherService{})

	middle(c)

	assert.Equal(t, http.StatusBadRequest, recorder.Code)
	assert.Equal(t, "{\"error\":\"malformed token\"}\n", recorder.Body.String())
	reader.AssertExpectations(t)
	tokenHandler.AssertExpectations(t)
}

func Test_Multi_HandlingOfAuthentication_UnauthorizedUser(t *testing.T) {
	tokenHandler := new(token.MockTokenHandler)
	tokenHandler.On("Retrieve", "token").Return(&token2.Token{
		UserId: 1000,
	}, nil)

	service := new(MockTeacherService)
	service.On("GetStudents", uint(1000)).Return(nil, nil).Once()

	reader := new(http2.MockReader)
	reader.On("GetParameter", controller.ParamUserId).Return("1", nil).Once()
	reader.On("GetHeader", middleware.HeaderAuthorization).Return("token")

	c, recorder := http2.CreateTestContext()
	c.Reader = reader

	middle := teacher.CreateMiddleware(tokenHandler, service)

	middle(c)

	assert.Equal(t, http.StatusUnauthorized, recorder.Code)
	assert.Equal(t, "{\"error\":\"unauthorized\"}\n", recorder.Body.String())
	service.AssertExpectations(t)
	reader.AssertExpectations(t)
	tokenHandler.AssertExpectations(t)
}

func Test_Multi_HandlingOfAuthentication_Authorized_SameUser(t *testing.T) {
	tokenHandler := new(token.MockTokenHandler)
	tokenHandler.On("Retrieve", "token").Return(&token2.Token{
		UserId: 1000,
	}, nil)

	reader := new(http2.MockReader)
	reader.On("GetParameter", controller.ParamUserId).Return("1000", nil).Once()
	reader.On("GetHeader", middleware.HeaderAuthorization).Return("token")

	c, recorder := http2.CreateTestContext()
	c.Reader = reader

	middle := teacher.CreateMiddleware(tokenHandler, &MockTeacherService{})

	middle(c)

	assert.Equal(t, http.StatusOK, recorder.Code)
	reader.AssertExpectations(t)
	tokenHandler.AssertExpectations(t)
}

func Test_Multi_HandlingOfAuthentication_Authorized_Student(t *testing.T) {
	tokenHandler := new(token.MockTokenHandler)
	tokenHandler.On("Retrieve", "token").Return(&token2.Token{
		UserId: 1001,
	}, nil)

	service := new(MockTeacherService)
	service.On("GetStudents", uint(1001)).Return([]model2.User{
		{
			Id: 999,
		},
		{
			Id: 1000,
		},
	}, nil).Once()

	reader := new(http2.MockReader)
	reader.On("GetParameter", controller.ParamUserId).Return("1000", nil).Once()
	reader.On("GetHeader", middleware.HeaderAuthorization).Return("token")

	c, recorder := http2.CreateTestContext()
	c.Reader = reader

	middle := teacher.CreateMiddleware(tokenHandler, service)

	middle(c)

	assert.Equal(t, http.StatusOK, recorder.Code)
	service.AssertExpectations(t)
	reader.AssertExpectations(t)
	tokenHandler.AssertExpectations(t)
}

func Test_Multi_HandlingOfAuthentication_Unauthorized_Student(t *testing.T) {
	tokenHandler := new(token.MockTokenHandler)
	tokenHandler.On("Retrieve", "token").Return(&token2.Token{
		UserId: 1002,
	}, nil)

	service := new(MockTeacherService)
	service.On("GetStudents", uint(1002)).Return([]model2.User{
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
	reader.On("GetParameter", controller.ParamUserId).Return("1000", nil).Once()
	reader.On("GetHeader", middleware.HeaderAuthorization).Return("token")

	c, recorder := http2.CreateTestContext()
	c.Reader = reader

	middle := teacher.CreateMiddleware(tokenHandler, service)

	middle(c)

	assert.Equal(t, http.StatusUnauthorized, recorder.Code)
	assert.Equal(t, "{\"error\":\"unauthorized\"}\n", recorder.Body.String())
	service.AssertExpectations(t)
	reader.AssertExpectations(t)
	tokenHandler.AssertExpectations(t)
}

func Test_Multi_HandlingOfAuthentication_Service_Error(t *testing.T) {
	tokenHandler := new(token.MockTokenHandler)
	tokenHandler.On("Retrieve", "token").Return(&token2.Token{
		UserId: 1001,
	}, nil)

	service := new(MockTeacherService)
	service.On("GetStudents", uint(1001)).Return([]model2.User{}, errors.New("whoops this fails")).Once()

	reader := new(http2.MockReader)
	reader.On("GetParameter", controller.ParamUserId).Return("1000", nil).Once()
	reader.On("GetHeader", middleware.HeaderAuthorization).Return("token")

	c, recorder := http2.CreateTestContext()
	c.Reader = reader

	middle := teacher.CreateMiddleware(tokenHandler, service)

	middle(c)

	assert.Equal(t, http.StatusInternalServerError, recorder.Code)
	assert.Equal(t, "{\"error\":\"internal error\"}\n", recorder.Body.String())
	service.AssertExpectations(t)
	reader.AssertExpectations(t)
	tokenHandler.AssertExpectations(t)
}
