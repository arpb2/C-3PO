package teacher_auth_test

import (
	"errors"
	"github.com/arpb2/C-3PO/src/api/auth/jwt"
	"github.com/arpb2/C-3PO/src/api/controller"
	"github.com/arpb2/C-3PO/src/api/engine/gin_engine"
	"github.com/arpb2/C-3PO/src/api/http_wrapper"
	"github.com/arpb2/C-3PO/src/api/middleware/auth/teacher_auth"
	"github.com/arpb2/C-3PO/src/api/model"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

var MultiTokenHandler = jwt.CreateTokenHandler()

type MockTeacherService struct{
	mock.Mock
}

func (m MockTeacherService) GetUser(userId uint) (user *model.User, err error) {
	panic("implement me")
}

func (m MockTeacherService) CreateUser(authenticatedUser *model.AuthenticatedUser) (user *model.User, err error) {
	panic("implement me")
}

func (m MockTeacherService) UpdateUser(authenticatedUser *model.AuthenticatedUser) (user *model.User, err error) {
	panic("implement me")
}

func (m MockTeacherService) DeleteUser(userId uint) error {
	panic("implement me")
}

func (m MockTeacherService) GetStudents(userId uint) (students *[]model.User, err error) {
	args := m.Called(userId)

	firstArg := args.Get(0)
	if firstArg != nil {
		students = firstArg.(*[]model.User)
	}

	err = args.Error(1)
	return
}

func performRequest(r http.Handler, method, path, body string, headers map[string][]string) *httptest.ResponseRecorder {
	req, _ := http.NewRequest(method, path, strings.NewReader(body))
	req.Header = headers

	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)
	return w
}

func Test_Multi_HandlingOfAuthentication_NoHeader(t *testing.T) {
	e := gin_engine.New()
	e.Register(controller.Controller{
		Method:        "GET",
		Path:          "/test",
		Middleware:    []http_wrapper.Handler{
			teacher_auth.CreateMiddleware(MultiTokenHandler, MockTeacherService{}),
		},
		Body:          func(ctx *http_wrapper.Context) {
			panic("Shouldn't reach here!")
		},
	})

	recorder := performRequest(e, "GET", "/test", "", map[string][]string{})

	assert.Equal(t, http.StatusUnauthorized, recorder.Code)
	assert.Equal(t, "{\"error\":\"unauthorized\"}\n", recorder.Body.String())
}

func Test_Multi_HandlingOfAuthentication_BadHeader(t *testing.T) {
	e := gin_engine.New()
	e.Register(controller.Controller{
		Method:        "GET",
		Path:          "/test",
		Middleware:    []http_wrapper.Handler{
			teacher_auth.CreateMiddleware(MultiTokenHandler, MockTeacherService{}),
		},
		Body:          func(ctx *http_wrapper.Context) {
			panic("Shouldn't reach here!")
		},
	})

	headers := map[string][]string{}
	headers["Authorization"] = []string{"bad token"}
	recorder := performRequest(e, "GET", "/test", "", headers)

	assert.Equal(t, http.StatusBadRequest, recorder.Code)
	assert.Equal(t, "{\"error\":\"malformed token\"}\n", recorder.Body.String())
}

func Test_Multi_HandlingOfAuthentication_UnauthorizedUser(t *testing.T) {
	service := new(MockTeacherService)
	service.On("GetStudents", uint(1000)).Return(nil, nil).Once()

	e := gin_engine.New()
	e.Register(controller.Controller{
		Method:        "GET",
		Path:          "/test/:user_id",
		Middleware:    []http_wrapper.Handler{
			teacher_auth.CreateMiddleware(MultiTokenHandler, service),
		},
		Body:          func(ctx *http_wrapper.Context) {
			panic("Shouldn't reach here!")
		},
	})

	headers := map[string][]string{}
	// Token for user 1000
	headers["Authorization"] = []string{"eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJVc2VySWQiOjEwMDB9.GVS-KC5nOCHybzzFIIH864u4KcGu-ZSd-96krqTUGWo"}
	recorder := performRequest(e, "GET", "/test/1", "", headers)

	assert.Equal(t, http.StatusUnauthorized, recorder.Code)
	assert.Equal(t, "{\"error\":\"unauthorized\"}\n", recorder.Body.String())
	service.AssertExpectations(t)

}

func Test_Multi_HandlingOfAuthentication_Authorized_SameUser(t *testing.T) {
	e := gin_engine.New()
	e.Register(controller.Controller{
		Method:        "GET",
		Path:          "/test/:user_id",
		Middleware:    []http_wrapper.Handler{
			teacher_auth.CreateMiddleware(MultiTokenHandler, MockTeacherService{}),
		},
		Body:          func(ctx *http_wrapper.Context) {
			ctx.WriteString(http.StatusOK, "Returned success")
		},
	})

	headers := map[string][]string{}
	// Token for user 1000
	headers["Authorization"] = []string{"eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJVc2VySWQiOjEwMDB9.GVS-KC5nOCHybzzFIIH864u4KcGu-ZSd-96krqTUGWo"}
	recorder := performRequest(e, "GET", "/test/1000", "", headers)

	assert.Equal(t, http.StatusOK, recorder.Code)
	assert.Equal(t, "Returned success", recorder.Body.String())
}

func Test_Multi_HandlingOfAuthentication_Authorized_Student(t *testing.T) {
	service := new(MockTeacherService)
	service.On("GetStudents", uint(1001)).Return(&[]model.User{
		{
			Id:      999,
		},
		{
			Id:      1000,
		},
	}, nil).Once()

	e := gin_engine.New()
	e.Register(controller.Controller{
		Method:        "GET",
		Path:          "/test/:user_id",
		Middleware:    []http_wrapper.Handler{
			teacher_auth.CreateMiddleware(MultiTokenHandler, service),
		},
		Body:          func(ctx *http_wrapper.Context) {
			ctx.WriteString(http.StatusOK, "Returned success")
		},
	})

	headers := map[string][]string{}
	// Token for user 1001
	headers["Authorization"] = []string{"eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJVc2VySWQiOjEwMDF9.Vx0MXNKC_A5s7rWZ_LfcwEc7rVgbns62mfFbq3RwSk0"}
	recorder := performRequest(e, "GET", "/test/1000", "", headers)

	assert.Equal(t, http.StatusOK, recorder.Code)
	assert.Equal(t, "Returned success", recorder.Body.String())
	service.AssertExpectations(t)
}


func Test_Multi_HandlingOfAuthentication_Unauthorized_Student(t *testing.T) {
	service := new(MockTeacherService)
	service.On("GetStudents", uint(1002)).Return(&[]model.User{
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

	e := gin_engine.New()
	e.Register(controller.Controller{
		Method:        "GET",
		Path:          "/test/:user_id",
		Middleware:    []http_wrapper.Handler{
			teacher_auth.CreateMiddleware(MultiTokenHandler, service),
		},
		Body:          func(ctx *http_wrapper.Context) {
			ctx.WriteString(http.StatusOK, "Returned success")
		},
	})

	headers := map[string][]string{}
	// Token for user 1002
	headers["Authorization"] = []string{"eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJVc2VySWQiOjEwMDJ9.YRQ2drFXz-apv85QOyMjNybmxsizVnfwImTWKwIVqso"}
	recorder := performRequest(e, "GET", "/test/1000", "", headers)

	assert.Equal(t, http.StatusUnauthorized, recorder.Code)
	assert.Equal(t, "{\"error\":\"unauthorized\"}\n", recorder.Body.String())
	service.AssertExpectations(t)
}

func Test_Multi_HandlingOfAuthentication_Service_Error(t *testing.T) {
	service := new(MockTeacherService)
	service.On("GetStudents", uint(1001)).Return(nil, errors.New("whoops this fails")).Once()

	e := gin_engine.New()
	e.Register(controller.Controller{
		Method:        "GET",
		Path:          "/test/:user_id",
		Middleware:    []http_wrapper.Handler{
			teacher_auth.CreateMiddleware(MultiTokenHandler, service),
		},
		Body:          func(ctx *http_wrapper.Context) {
			ctx.WriteString(http.StatusOK, "Returned success")
		},
	})

	headers := map[string][]string{}
	// Token for user 1001
	headers["Authorization"] = []string{"eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJVc2VySWQiOjEwMDF9.Vx0MXNKC_A5s7rWZ_LfcwEc7rVgbns62mfFbq3RwSk0"}
	recorder := performRequest(e, "GET", "/test/1000", "", headers)

	assert.Equal(t, http.StatusInternalServerError, recorder.Code)
	assert.Equal(t, "{\"error\":\"internal error\"}\n", recorder.Body.String())
	service.AssertExpectations(t)
}
