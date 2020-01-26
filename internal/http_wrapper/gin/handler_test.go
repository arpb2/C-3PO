package gin_wrapper_test

import (
	"github.com/arpb2/C-3PO/api/http_wrapper"
	api_http_wrapper "github.com/arpb2/C-3PO/api/http_wrapper"
	internal_http_wrapper "github.com/arpb2/C-3PO/internal/http_wrapper/gin"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCreateHandler_Delegates(t *testing.T) {
	invocations := 0
	handlerFunc := func(ctx *api_http_wrapper.Context) {
		invocations++
	}

	internal_http_wrapper.CreateHandler(handlerFunc)(new(gin.Context))

	assert.Equal(t, 1, invocations)
}

func TestCreateHandlers_Delegates(t *testing.T) {
	invocations := 0
	handlerFunc := func(ctx *http_wrapper.Context) {
		invocations++
	}

	handlers := internal_http_wrapper.CreateHandlers(handlerFunc, handlerFunc)

	assert.Equal(t, 2, len(handlers))

	handlers[0](new(gin.Context))
	handlers[1](new(gin.Context))

	assert.Equal(t, 2, invocations)
}
