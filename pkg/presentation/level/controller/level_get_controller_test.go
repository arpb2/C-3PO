package controller_test

import (
	"testing"

	"github.com/arpb2/C-3PO/pkg/domain/controller"
	"github.com/arpb2/C-3PO/pkg/domain/http"
	level "github.com/arpb2/C-3PO/pkg/presentation/level/controller"
	httpmock "github.com/arpb2/C-3PO/test/mock/http"
	"github.com/stretchr/testify/assert"
)

func createGetController() controller.Controller {
	return level.CreateGetController()
}

func TestGetController_IsGet(t *testing.T) {
	assert.Equal(t, "GET", createGetController().Method)
}

func TestGetControllerPath_IsLevels(t *testing.T) {
	assert.Equal(t, "/levels/:id", createGetController().Path)
}

func TestGetController_IsStubbed(t *testing.T) {
	writer := new(httpmock.MockWriter)
	writer.On("WriteString", 200, "stub", []interface{}(nil)).Once()
	ctx := &http.Context{
		Writer: writer,
	}

	createGetController().Body(ctx)

	writer.AssertExpectations(t)
}
