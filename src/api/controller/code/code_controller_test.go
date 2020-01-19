package code_test

import (
	"github.com/arpb2/C-3PO/src/api/controller/code"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

type SharedInMemoryCodeService struct {
	codeId string
	code   *string
	err    error
}
func (s *SharedInMemoryCodeService) Read(userId string, codeId string) (code *string, err error) {
	s.codeId = codeId
	return s.code, s.err
}

func (s *SharedInMemoryCodeService) Write(userId string, code *string) (codeId string, err error) {
	s.code = code
	return s.codeId, s.err
}

func (s *SharedInMemoryCodeService) Replace(userId string, codeId string, code *string) error {
	s.code = code
	s.codeId = codeId
	return s.err
}

func init() {
	gin.SetMode(gin.TestMode)
}

func TestFetchCodeId_RetrievesFromParam(t *testing.T) {
	recorder := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(recorder)
	c.Params = append(c.Params, gin.Param{
		Key:   "code_id",
		Value: "1234",
	})

	codeId, halt := code.FetchCodeId(c)

	assert.False(t, halt)
	assert.Equal(t, "1234", codeId)
}

func TestFetchCodeId_HaltsWith400OnError(t *testing.T) {
	recorder := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(recorder)

	codeId, halt := code.FetchCodeId(c)

	assert.True(t, halt)
	assert.Empty(t, codeId)
	assert.Equal(t, http.StatusBadRequest, recorder.Code)
}

func TestFetchUserId_RetrievesFromParam(t *testing.T) {
	recorder := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(recorder)
	c.Params = append(c.Params, gin.Param{
		Key:   "user_id",
		Value: "1234",
	})

	userId, halt := code.FetchUserId(c)

	assert.False(t, halt)
	assert.Equal(t, "1234", userId)
}

func TestFetchUserId_HaltsWith400OnError(t *testing.T) {
	recorder := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(recorder)

	userId, halt := code.FetchUserId(c)

	assert.True(t, halt)
	assert.Empty(t, userId)
	assert.Equal(t, http.StatusBadRequest, recorder.Code)
}

func TestFetchCode_RetrievesFromPart(t *testing.T) {
	recorder := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(recorder)
	c.Request, _ = http.NewRequest("GET", "/test", strings.NewReader(""))
	c.Request.PostForm = map[string][]string{}
	c.Request.PostForm.Set("code", "test code")
	_ = c.Request.ParseForm()

	rawCode, halt := code.FetchCode(c)

	assert.False(t, halt)
	assert.Equal(t, "test code", *rawCode)
}

func TestFetchCode_HaltsWith400OnError(t *testing.T) {
	recorder := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(recorder)
	c.Request, _ = http.NewRequest("GET", "/test", strings.NewReader(""))
	c.Request.PostForm = map[string][]string{}
	_ = c.Request.ParseForm()

	rawCode, halt := code.FetchCode(c)

	assert.True(t, halt)
	assert.Nil(t, rawCode)
	assert.Equal(t, http.StatusBadRequest, recorder.Code)
}