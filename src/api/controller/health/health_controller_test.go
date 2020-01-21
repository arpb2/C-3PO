package health_test

import (
	"github.com/arpb2/C-3PO/src/api/controller/health"
	"github.com/arpb2/C-3PO/src/api/http_wrapper/gin_wrapper"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestHealthControllerMethodIsGET(t *testing.T) {
	assert.Equal(t, "GET", health.CreateGetController().Method)
}

func TestHealthControllerPathIsPing(t *testing.T) {
	assert.Equal(t, "/ping", health.CreateGetController().Path)
}

func TestHealthControllerBodyReturnsWith200IfItsOk(t *testing.T) {
	c, w := gin_wrapper.CreateTestContext()

	health.CreateGetController().Body(c)

	assert.Equal(t, 200, w.Code)
}
