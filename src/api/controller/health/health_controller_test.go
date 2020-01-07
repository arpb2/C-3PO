package health_test

import (
	"github.com/arpb2/C-3PO/src/api/controller/health"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"net/http/httptest"
	"testing"
)

func TestHealthControllerMethodIsGET(t *testing.T) {
	assert.Equal(t, "GET", health.GetController.Method)
}

func TestHealthControllerPathIsPing(t *testing.T) {
	assert.Equal(t, "/ping", health.GetController.Path)
}

func TestHealthControllerBodyReturnsWith200IfItsOk(t *testing.T) {
	gin.SetMode(gin.TestMode)

	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)

	health.GetController.Body(c)

	assert.Equal(t, 200, c.Writer.Status())
}
