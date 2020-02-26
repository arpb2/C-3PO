package gin_test

import (
	"net/http/httptest"
	"testing"

	this "github.com/arpb2/C-3PO/cmd/c3po/infrastructure/http/gin"
	"github.com/arpb2/C-3PO/pkg/domain/architecture/http"
	gin "github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func TestWriter_GivenOne_WhenWritingString_ThenRecorderGetsTheString(t *testing.T) {
	gin.SetMode(gin.TestMode)
	recorder := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(recorder)
	writer := this.CreateWriter(c)

	writer.WriteString(201, "some content")

	assert.Equal(t, 201, recorder.Code)
	assert.Equal(t, "some content", recorder.Body.String())
}

func TestWriter_GivenOne_WhenWritingJson_ThenRecorderGetsTheJson(t *testing.T) {
	gin.SetMode(gin.TestMode)
	recorder := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(recorder)
	writer := this.CreateWriter(c)

	writer.WriteJson(201, http.Json{
		"key": "value",
	})

	assert.Equal(t, 201, recorder.Code)
	assert.Equal(t, "{\"key\":\"value\"}\n", recorder.Body.String())
}

func TestWriter_GivenOne_WhenWritingStatus_ThenRecorderGetsTheStatus(t *testing.T) {
	gin.SetMode(gin.TestMode)
	recorder := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(recorder)
	writer := this.CreateWriter(c)

	writer.WriteStatus(200)

	assert.Equal(t, 200, recorder.Code)
}
