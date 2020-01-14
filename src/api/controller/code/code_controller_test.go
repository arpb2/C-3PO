package code_test

import (
	"github.com/arpb2/C-3PO/src/api/controller/code"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"net/http/httptest"
	"testing"
)

////////////////////////////////////////// GET

func TestCodeGetControllerMethodIsGET(t *testing.T) {
	assert.Equal(t, "GET", code.GetController.Method)
}

func TestCodeGetControllerPathIsAsExpected(t *testing.T) {
	assert.Equal(t, "/users/:user_id/codes/:code_id", code.GetController.Path)
}

func TestCodeGetControllerBodyReturnsWith200IfItsOk(t *testing.T) {
	gin.SetMode(gin.TestMode)

	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)

	code.GetController.Body(c)

	assert.Equal(t, 200, c.Writer.Status())
}

////////////////////////////////////////// POST

func TestCodePostControllerMethodIsPOST(t *testing.T) {
	assert.Equal(t, "POST", code.PostController.Method)
}

func TestCodePostControllerPathIsAsExpected(t *testing.T) {
	assert.Equal(t, "/users/:user_id/codes", code.PostController.Path)
}

func TestCodePostControllerBodyReturnsWith200IfItsOk(t *testing.T) {
	gin.SetMode(gin.TestMode)

	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)

	code.PostController.Body(c)

	assert.Equal(t, 200, c.Writer.Status())
}

/////////////////////////////////////////// PUT

func TestCodePutControllerMethodIsPUT(t *testing.T) {
	assert.Equal(t, "PUT", code.PutController.Method)
}

func TestCodePutControllerPathIsAsExpected(t *testing.T) {
	assert.Equal(t, "/users/:user_id/codes/:code_id", code.PutController.Path)
}

func TestCodePutControllerBodyReturnsWith200IfItsOk(t *testing.T) {
	gin.SetMode(gin.TestMode)

	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)

	code.PutController.Body(c)

	assert.Equal(t, 200, c.Writer.Status())
}