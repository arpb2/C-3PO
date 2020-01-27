package health_test

import (
	"testing"

	"github.com/arpb2/C-3PO/hack/http"
	healthcontroller "github.com/arpb2/C-3PO/pkg/controller/health"
	"github.com/stretchr/testify/assert"
)

func TestHealthControllerMethodIsGET(t *testing.T) {
	assert.Equal(t, "GET", healthcontroller.CreateGetController().Method)
}

func TestHealthControllerPathIsPing(t *testing.T) {
	assert.Equal(t, "/ping", healthcontroller.CreateGetController().Path)
}

func TestHealthControllerBodyReturnsWith200IfItsOk(t *testing.T) {
	c, w := http.CreateTestContext()

	healthcontroller.CreateGetController().Body(c)

	assert.Equal(t, 200, w.Code)
}
