package gin_wrapper

import (
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCreateContext_Reader_DelegatesToGin(t *testing.T) {
	ginContext := new(gin.Context)
	ginContext.Params = gin.Params{
		gin.Param{
			Key:   "test",
			Value: "test",
		},
	}
	context := CreateContext(ginContext)

	assert.Equal(t, "test", context.Param("test"))
}

func TestCreateContext_Writer_DelegatesToGin(t *testing.T) {
	ginContext := new(gin.Context)
	context := CreateContext(ginContext)

	assert.Equal(t, context.Writer, ginContext)
}

func TestCreateContext_Middleware_DelegatesToGin(t *testing.T) {
	ginContext := new(gin.Context)
	context := CreateContext(ginContext)

	assert.Equal(t, context.Middleware, ginContext)
}
