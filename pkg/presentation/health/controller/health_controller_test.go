package controller_test

import (
	"github.com/arpb2/C-3PO/pkg/presentation/health/controller"
	"testing"

	"github.com/arpb2/C-3PO/test/mock/http"
	"github.com/stretchr/testify/assert"
)

func TestHealthControllerMethodIsGET(t *testing.T) {
	assert.Equal(t, "GET", controller.CreateGetController().Method)
}

func TestHealthControllerPathIsPing(t *testing.T) {
	assert.Equal(t, "/ping", controller.CreateGetController().Path)
}

func TestHealthControllerBodyReturnsWith200IfItsOk(t *testing.T) {
	c, w := http.CreateTestContext()

	controller.CreateGetController().Body(c)

	assert.Equal(t, 200, w.Code)
}
