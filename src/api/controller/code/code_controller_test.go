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
	codeId uint
	code   *string
	err    error
}
func (s *SharedInMemoryCodeService) GetCode(userId uint, codeId uint) (code *string, err error) {
	s.codeId = codeId
	return s.code, s.err
}

func (s *SharedInMemoryCodeService) CreateCode(userId uint, code *string) (codeId uint, err error) {
	s.code = code
	return s.codeId, s.err
}

func (s *SharedInMemoryCodeService) ReplaceCode(userId uint, codeId uint, code *string) error {
	s.code = code
	s.codeId = codeId
	return s.err
}

func TestFetchCodeId_RetrievesFromParam(t *testing.T) {
	reader := new(http_wrapper.TestReader)
	reader.On("GetParameter", "code_id").Return("1234").Once()

	c, _ := gin_wrapper.CreateTestContext()
	c.Reader = reader

	codeId, halt := code.FetchCodeId(c)

	assert.False(t, halt)
	assert.Equal(t, uint(1234), codeId)
	reader.AssertExpectations(t)
}

func TestFetchCodeId_HaltsWith400OnError(t *testing.T) {
	c, recorder := gin_wrapper.CreateTestContext()

	codeId, halt := code.FetchCodeId(c)

	assert.True(t, halt)
	assert.Zero(t, codeId)
	assert.Equal(t, http.StatusBadRequest, recorder.Code)
}

func TestFetchCodeId_RetrievesFromParam_400IfMalformed(t *testing.T) {
	reader := new(http_wrapper.TestReader)
	reader.On("GetParameter", "code_id").Return("not a number").Once()

	c, _ := gin_wrapper.CreateTestContext()
	c.Reader = reader

	codeId, halt := code.FetchCodeId(c)

	assert.True(t, halt)
	assert.Zero(t, codeId)
	reader.AssertExpectations(t)
}

func TestFetchUserId_RetrievesFromParam_400IfMalformed(t *testing.T) {
	reader := new(http_wrapper.TestReader)
	reader.On("GetParameter", "user_id").Return("not a number").Once()

	c, _ := gin_wrapper.CreateTestContext()
	c.Reader = reader

	userId, halt := code.FetchUserId(c)

	assert.True(t, halt)
	assert.Zero(t, userId)
	reader.AssertExpectations(t)
}

func TestFetchUserId_RetrievesFromParam(t *testing.T) {
	reader := new(http_wrapper.TestReader)
	reader.On("GetParameter", "user_id").Return("1234").Once()

	c, _ := gin_wrapper.CreateTestContext()
	c.Reader = reader

	userId, halt := code.FetchUserId(c)

	assert.False(t, halt)
	assert.Equal(t, uint(1234), userId)
	reader.AssertExpectations(t)
}

func TestFetchUserId_HaltsWith400OnError(t *testing.T) {
	c, recorder := gin_wrapper.CreateTestContext()

	userId, halt := code.FetchUserId(c)

	assert.True(t, halt)
	assert.Zero(t, userId)
	assert.Equal(t, http.StatusBadRequest, recorder.Code)
}

func TestFetchCode_RetrievesFromPart(t *testing.T) {
	reader := new(http_wrapper.TestReader)
	reader.On("GetFormData", "code").Return("test code", true).Once()

	c, _ := gin_wrapper.CreateTestContext()
	c.Reader = reader

	rawCode, halt := code.FetchCode(c)

	assert.False(t, halt)
	assert.Equal(t, "test code", *rawCode)
	reader.AssertExpectations(t)
}

func TestFetchCode_HaltsWith400_OnError(t *testing.T) {
	reader := new(http_wrapper.TestReader)
	reader.On("GetFormData", "code").Return("", false).Once()

	c, recorder := gin_wrapper.CreateTestContext()
	c.Reader = reader

	rawCode, halt := code.FetchCode(c)

	assert.True(t, halt)
	assert.Nil(t, rawCode)
	assert.Equal(t, http.StatusBadRequest, recorder.Code)
	reader.AssertExpectations(t)
}