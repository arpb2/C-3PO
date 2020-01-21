package gin_wrapper

import (
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCreateContext_Reader_Exists(t *testing.T) {
	ginContext := new(gin.Context)
	context := CreateContext(ginContext)

	assert.NotNil(t, context.Reader)
}

func TestCreateContext_Writer_Exists(t *testing.T) {
	ginContext := new(gin.Context)
	context := CreateContext(ginContext)

	assert.NotNil(t, context.Writer)
}

func TestCreateContext_Middleware_Exists(t *testing.T) {
	ginContext := new(gin.Context)
	context := CreateContext(ginContext)

	assert.NotNil(t, context.Middleware)
}
