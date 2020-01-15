package auth_test

import (
	"errors"
	"github.com/arpb2/C-3PO/src/api/controller"
	"github.com/arpb2/C-3PO/src/api/engine"
	"github.com/arpb2/C-3PO/src/api/middleware/auth"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"net/http"
	"testing"
)

func init() {
	auth.TeacherService = MockTeacherService{}
}

func Test_Multi_HandlingOfAuthentication_NoHeader(t *testing.T) {
	e := engine.CreateBasicServerEngine()
	e.Register(controller.Controller{
		Method:        "GET",
		Path:          "/test",
		Middleware:    []gin.HandlerFunc{
			auth.UserOrTeacherAuthenticationMiddleware,
		},
		Body:          func(ctx *gin.Context) {
			panic("Shouldn't reach here!")
		},
	})

	recorder := performRequest(e, "GET", "/test", "", map[string][]string{})

	assert.Equal(t, http.StatusUnauthorized, recorder.Code)
	assert.Equal(t, "{\"error\":\"no 'Authorization' header provided\"}\n", recorder.Body.String())
}

func Test_Multi_HandlingOfAuthentication_BadHeader(t *testing.T) {
	e := engine.CreateBasicServerEngine()
	e.Register(controller.Controller{
		Method:        "GET",
		Path:          "/test",
		Middleware:    []gin.HandlerFunc{
			auth.UserOrTeacherAuthenticationMiddleware,
		},
		Body:          func(ctx *gin.Context) {
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
	e := engine.CreateBasicServerEngine()
	e.Register(controller.Controller{
		Method:        "GET",
		Path:          "/test/:user_id",
		Middleware:    []gin.HandlerFunc{
			auth.UserOrTeacherAuthenticationMiddleware,
		},
		Body:          func(ctx *gin.Context) {
			panic("Shouldn't reach here!")
		},
	})

	headers := map[string][]string{}
	// Token for user 1000
	headers["Authorization"] = []string{"eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJVc2VySWQiOjEwMDB9.GVS-KC5nOCHybzzFIIH864u4KcGu-ZSd-96krqTUGWo"}
	recorder := performRequest(e, "GET", "/test/1", "", headers)

	assert.Equal(t, http.StatusUnauthorized, recorder.Code)
	assert.Equal(t, "{\"error\":\"unauthorized\"}\n", recorder.Body.String())
}

func Test_Multi_HandlingOfAuthentication_Authorized_SameUser(t *testing.T) {
	e := engine.CreateBasicServerEngine()
	e.Register(controller.Controller{
		Method:        "GET",
		Path:          "/test/:user_id",
		Middleware:    []gin.HandlerFunc{
			auth.UserOrTeacherAuthenticationMiddleware,
		},
		Body:          func(ctx *gin.Context) {
			ctx.String(http.StatusOK, "Returned success")
		},
	})

	headers := map[string][]string{}
	// Token for user 1000
	headers["Authorization"] = []string{"eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJVc2VySWQiOjEwMDB9.GVS-KC5nOCHybzzFIIH864u4KcGu-ZSd-96krqTUGWo"}
	recorder := performRequest(e, "GET", "/test/1000", "", headers)

	assert.Equal(t, http.StatusOK, recorder.Code)
	assert.Equal(t, "Returned success", recorder.Body.String())
}

type MockTeacherService struct{}
func (s MockTeacherService) GetUser(userId string) (user *interface{}, err error) {
	panic("implement me")
}
func (s MockTeacherService) CreateUser(data ...interface{}) (user *interface{}, err error) {
	panic("implement me")
}
func (s MockTeacherService) UpdateUser(user interface{}) error {
	panic("implement me")
}
func (s MockTeacherService) DeleteUser(userId string) error {
	panic("implement me")
}
func (s MockTeacherService) GetStudents(userId string) (students *[]int, err error) {
	if userId == "1001" {
		students = &[]int{
			1000,
		}
	} else {
		students = nil
	}
	err = nil
	return
}

func Test_Multi_HandlingOfAuthentication_Authorized_Student(t *testing.T) {
	e := engine.CreateBasicServerEngine()
	e.Register(controller.Controller{
		Method:        "GET",
		Path:          "/test/:user_id",
		Middleware:    []gin.HandlerFunc{
			auth.UserOrTeacherAuthenticationMiddleware,
		},
		Body:          func(ctx *gin.Context) {
			ctx.String(http.StatusOK, "Returned success")
		},
	})

	auth.TeacherService = MockTeacherService{}

	headers := map[string][]string{}
	// Token for user 1001
	headers["Authorization"] = []string{"eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJVc2VySWQiOjEwMDF9.Vx0MXNKC_A5s7rWZ_LfcwEc7rVgbns62mfFbq3RwSk0"}
	recorder := performRequest(e, "GET", "/test/1000", "", headers)

	assert.Equal(t, http.StatusOK, recorder.Code)
	assert.Equal(t, "Returned success", recorder.Body.String())
}


func Test_Multi_HandlingOfAuthentication_Unauthorized_Student(t *testing.T) {
	e := engine.CreateBasicServerEngine()
	e.Register(controller.Controller{
		Method:        "GET",
		Path:          "/test/:user_id",
		Middleware:    []gin.HandlerFunc{
			auth.UserOrTeacherAuthenticationMiddleware,
		},
		Body:          func(ctx *gin.Context) {
			ctx.String(http.StatusOK, "Returned success")
		},
	})

	auth.TeacherService = MockTeacherService{}

	headers := map[string][]string{}
	// Token for user 1002
	headers["Authorization"] = []string{"eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJVc2VySWQiOjEwMDJ9.YRQ2drFXz-apv85QOyMjNybmxsizVnfwImTWKwIVqso"}
	recorder := performRequest(e, "GET", "/test/1000", "", headers)

	assert.Equal(t, http.StatusUnauthorized, recorder.Code)
	assert.Equal(t, "{\"error\":\"unauthorized\"}\n", recorder.Body.String())
}

type FailingMockTeacherService struct{}
func (s FailingMockTeacherService) GetUser(userId string) (user *interface{}, err error) {
	panic("implement me")
}
func (s FailingMockTeacherService) CreateUser(data ...interface{}) (user *interface{}, err error) {
	panic("implement me")
}
func (s FailingMockTeacherService) UpdateUser(user interface{}) error {
	panic("implement me")
}
func (s FailingMockTeacherService) DeleteUser(userId string) error {
	panic("implement me")
}
func (s FailingMockTeacherService) GetStudents(userId string) (students *[]int, err error) {
	students = nil
	err = errors.New("woops this fails")
	return
}

func Test_Multi_HandlingOfAuthentication_Service_Error(t *testing.T) {
	e := engine.CreateBasicServerEngine()
	e.Register(controller.Controller{
		Method:        "GET",
		Path:          "/test/:user_id",
		Middleware:    []gin.HandlerFunc{
			auth.UserOrTeacherAuthenticationMiddleware,
		},
		Body:          func(ctx *gin.Context) {
			ctx.String(http.StatusOK, "Returned success")
		},
	})

	auth.TeacherService = FailingMockTeacherService{}

	headers := map[string][]string{}
	// Token for user 1001
	headers["Authorization"] = []string{"eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJVc2VySWQiOjEwMDF9.Vx0MXNKC_A5s7rWZ_LfcwEc7rVgbns62mfFbq3RwSk0"}
	recorder := performRequest(e, "GET", "/test/1000", "", headers)

	assert.Equal(t, http.StatusInternalServerError, recorder.Code)
	assert.Equal(t, "{\"error\":\"internal error\"}\n", recorder.Body.String())
}
