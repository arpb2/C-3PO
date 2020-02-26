package gin_test

import (
	"testing"

	internalhttpwrapper "github.com/arpb2/C-3PO/cmd/c3po/infrastructure/http/gin"
	apihttpwrapper "github.com/arpb2/C-3PO/pkg/domain/architecture/http"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func TestCreateHandler_Delegates(t *testing.T) {
	invocations := 0
	handlerFunc := func(ctx *apihttpwrapper.Context) {
		invocations++
	}

	internalhttpwrapper.CreateHandler(handlerFunc)(new(gin.Context))

	assert.Equal(t, 1, invocations)
}

func TestCreateHandlers_Delegates(t *testing.T) {
	invocations := 0
	handlerFunc := func(ctx *apihttpwrapper.Context) {
		invocations++
	}

	handlers := internalhttpwrapper.CreateHandlers(handlerFunc, handlerFunc)

	assert.Equal(t, 2, len(handlers))

	handlers[0](new(gin.Context))
	handlers[1](new(gin.Context))

	assert.Equal(t, 2, invocations)
}
