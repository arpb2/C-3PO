package health_test

import (
	"github.com/arpb2/C-3PO/src/api/controller/health"
	"github.com/arpb2/C-3PO/src/api/http_wrapper/gin_wrapper"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestHealthControllerMethodIsGET(t *testing.T) {
	assert.Equal(t, "GET", health.GetController.Method)
}

func TestHealthControllerPathIsPing(t *testing.T) {
	assert.Equal(t, "/ping", health.GetController.Path)
}

func TestHealthControllerBodyReturnsWith200IfItsOk(t *testing.T) {
	c, w := gin_wrapper.CreateTestContext()

	health.GetController.Body(c)

	assert.Equal(t, 200, w.Code)
}
