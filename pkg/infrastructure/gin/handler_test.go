package gin_test

import (
	"testing"

	gin2 "github.com/arpb2/C-3PO/pkg/infrastructure/gin"

	apihttpwrapper "github.com/arpb2/C-3PO/pkg/domain/http"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func TestCreateHandler_Delegates(t *testing.T) {
	invocations := 0
	handlerFunc := func(ctx *apihttpwrapper.Context) {
		invocations++
	}

	gin2.CreateHandler(handlerFunc)(new(gin.Context))

	assert.Equal(t, 1, invocations)
}

func TestCreateHandlers_Delegates(t *testing.T) {
	invocations := 0
	handlerFunc := func(ctx *apihttpwrapper.Context) {
		invocations++
	}

	handlers := gin2.CreateHandlers(handlerFunc, handlerFunc)

	assert.Equal(t, 2, len(handlers))

	handlers[0](new(gin.Context))
	handlers[1](new(gin.Context))

	assert.Equal(t, 2, invocations)
}
