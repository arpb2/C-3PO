package controller_test

import (
	"testing"

	"github.com/arpb2/C-3PO/pkg/domain/controller"
	"github.com/arpb2/C-3PO/pkg/domain/http"
	level "github.com/arpb2/C-3PO/pkg/presentation/level/controller"
	httpmock "github.com/arpb2/C-3PO/test/mock/http"
	"github.com/stretchr/testify/assert"
)

func createPutController() controller.Controller {
	return level.CreatePutController(func(ctx *http.Context) {
		// Empty
	})
}

func TestPutController_IsPut(t *testing.T) {
	assert.Equal(t, "PUT", createPutController().Method)
}

func TestPutControllerPath_IsLevels(t *testing.T) {
	assert.Equal(t, "/levels/:id", createPutController().Path)
}

func TestPutController_IsStubbed(t *testing.T) {
	writer := new(httpmock.MockWriter)
	writer.On("WriteString", 200, "stub", []interface{}(nil)).Once()
	ctx := &http.Context{
		Writer: writer,
	}

	createPutController().Body(ctx)

	writer.AssertExpectations(t)
}
