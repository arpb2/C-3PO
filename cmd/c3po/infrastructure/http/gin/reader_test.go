package gin_test

import (
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
	"testing"

	ginwrapper "github.com/arpb2/C-3PO/cmd/c3po/infrastructure/http/gin"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func TestCreateReader_Url(t *testing.T) {
	ctx := new(gin.Context)
	ctx.Request = &http.Request{
		URL: &url.URL{
			Scheme: "https",
			Host:   "host",
			Path:   "/path",
		},
	}
	reader := ginwrapper.CreateReader(ctx)

	assert.Equal(t, "https://host/path", reader.GetUrl())
}

func TestCreateReader_Method(t *testing.T) {
	ctx := new(gin.Context)
	ctx.Request = &http.Request{
		Method: "POST",
	}
	reader := ginwrapper.CreateReader(ctx)

	assert.Equal(t, "POST", reader.GetMethod())
}

func TestCreateReader_GetHeader(t *testing.T) {
	expectedKey := "test-key"
	expectedValue := "test-value"

	ctx := new(gin.Context)
	ctx.Request = &http.Request{}
	ctx.Request.Header = map[string][]string{}
	ctx.Request.Header.Set(expectedKey, expectedValue)
	reader := ginwrapper.CreateReader(ctx)

	assert.Equal(t, expectedValue, reader.GetHeader(expectedKey))
}

func TestCreateReader_Param(t *testing.T) {
	expectedKey := "test-key"
	expectedValue := "test-value"

	ctx := new(gin.Context)
	ctx.Params = gin.Params{
		gin.Param{
			Key:   expectedKey,
			Value: expectedValue,
		},
	}
	reader := ginwrapper.CreateReader(ctx)

	assert.Equal(t, expectedValue, reader.GetParameter(expectedKey))
}

func TestCreateReader_GetPostForm(t *testing.T) {
	expectedKey := "test-key"
	expectedValue := "test-value"

	ctx, _ := gin.CreateTestContext(nil)
	ctx.Request = &http.Request{}
	ctx.Request.PostForm = map[string][]string{}
	ctx.Request.PostForm.Set(expectedKey, expectedValue)

	err := ctx.Request.ParseForm()

	assert.Nil(t, err)

	reader := ginwrapper.CreateReader(ctx)

	resultValue, exists := reader.GetFormData(expectedKey)

	assert.True(t, exists)
	assert.Equal(t, expectedValue, resultValue)
}

type InnerStruct struct {
	Name string `json:"name"`
	Age  uint   `json:"age"`
}

type testStruct struct {
	*InnerStruct
	Wealth int `json:"wealth"`
}

func TestCreateReader_ShouldBindJson(t *testing.T) {
	structJson := `
{
  "name": "test name",
  "age": 20,
  "wealth": 123456
}
`

	ctx := new(gin.Context)
	ctx.Request = &http.Request{
		Body: ioutil.NopCloser(strings.NewReader(structJson)),
	}
	reader := ginwrapper.CreateReader(ctx)

	var result testStruct
	err := reader.ReadBody(&result)

	assert.Nil(t, err)
	assert.Equal(t, "test name", result.Name)
	assert.Equal(t, uint(20), result.Age)
	assert.Equal(t, 123456, result.Wealth)
}
