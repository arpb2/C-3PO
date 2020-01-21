package code_test

import (
	"github.com/arpb2/C-3PO/src/api/controller/code"
	"github.com/arpb2/C-3PO/src/api/http_wrapper"
	"github.com/arpb2/C-3PO/src/api/http_wrapper/gin_wrapper"
	"github.com/stretchr/testify/assert"
	"net/http"
	"testing"
)

type SharedInMemoryCodeService struct {
	codeId string
	code   *string
	err    error
}
func (s *SharedInMemoryCodeService) GetCode(userId string, codeId string) (code *string, err error) {
	s.codeId = codeId
	return s.code, s.err
}

func (s *SharedInMemoryCodeService) CreateCode(userId string, code *string) (codeId string, err error) {
	s.code = code
	return s.codeId, s.err
}

func (s *SharedInMemoryCodeService) ReplaceCode(userId string, codeId string, code *string) error {
	s.code = code
	s.codeId = codeId
	return s.err
}

func TestFetchCodeId_RetrievesFromParam(t *testing.T) {
	reader := new(http_wrapper.TestReader)
	reader.On("Param", "code_id").Return("1234").Once()

	c, _ := gin_wrapper.CreateTestContext()
	c.Reader = reader

	codeId, halt := code.FetchCodeId(c)

	assert.False(t, halt)
	assert.Equal(t, "1234", codeId)
	reader.AssertExpectations(t)
}

func TestFetchCodeId_HaltsWith400OnError(t *testing.T) {
	c, recorder := gin_wrapper.CreateTestContext()

	codeId, halt := code.FetchCodeId(c)

	assert.True(t, halt)
	assert.Empty(t, codeId)
	assert.Equal(t, http.StatusBadRequest, recorder.Code)
}

func TestFetchUserId_RetrievesFromParam(t *testing.T) {
	reader := new(http_wrapper.TestReader)
	reader.On("Param", "user_id").Return("1234").Once()

	c, _ := gin_wrapper.CreateTestContext()
	c.Reader = reader

	userId, halt := code.FetchUserId(c)

	assert.False(t, halt)
	assert.Equal(t, "1234", userId)
	reader.AssertExpectations(t)
}

func TestFetchUserId_HaltsWith400OnError(t *testing.T) {
	c, recorder := gin_wrapper.CreateTestContext()

	userId, halt := code.FetchUserId(c)

	assert.True(t, halt)
	assert.Empty(t, userId)
	assert.Equal(t, http.StatusBadRequest, recorder.Code)
}

func TestFetchCode_RetrievesFromPart(t *testing.T) {
	reader := new(http_wrapper.TestReader)
	reader.On("GetPostForm", "code").Return("test code", true).Once()

	c, _ := gin_wrapper.CreateTestContext()
	c.Reader = reader

	rawCode, halt := code.FetchCode(c)

	assert.False(t, halt)
	assert.Equal(t, "test code", *rawCode)
	reader.AssertExpectations(t)
}

func TestFetchCode_HaltsWith400_OnError(t *testing.T) {
	reader := new(http_wrapper.TestReader)
	reader.On("GetPostForm", "code").Return("", false).Once()

	c, recorder := gin_wrapper.CreateTestContext()
	c.Reader = reader

	rawCode, halt := code.FetchCode(c)

	assert.True(t, halt)
	assert.Nil(t, rawCode)
	assert.Equal(t, http.StatusBadRequest, recorder.Code)
	reader.AssertExpectations(t)
}