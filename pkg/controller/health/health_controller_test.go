package health_controller_test

import (
	"testing"

	"github.com/arpb2/C-3PO/hack/http_wrapper"
	health_controller "github.com/arpb2/C-3PO/pkg/controller/health"
	"github.com/stretchr/testify/assert"
)

func TestHealthControllerMethodIsGET(t *testing.T) {
	assert.Equal(t, "GET", health_controller.CreateGetController().Method)
}

func TestHealthControllerPathIsPing(t *testing.T) {
	assert.Equal(t, "/ping", health_controller.CreateGetController().Path)
}

func TestHealthControllerBodyReturnsWith200IfItsOk(t *testing.T) {
	c, w := http_wrapper.CreateTestContext()

	health_controller.CreateGetController().Body(c)

	assert.Equal(t, 200, w.Code)
}
