package gin_test

import (
	"errors"
	"net/http/httptest"
	"testing"

	"github.com/arpb2/C-3PO/pkg/domain/http"
	this "github.com/arpb2/C-3PO/pkg/infra/http/gin"
	gin "github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func TestMiddleware_GivenOne_WhenAbortingWith4xxError_ThenGinAbortWithErrorIsCalled(t *testing.T) {
	gin.SetMode(gin.TestMode)
	recorder := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(recorder)
	middleware := this.CreateMiddleware(c)

	middleware.AbortTransactionWithError(http.CreateNotFoundError())

	assert.Equal(t, 404, recorder.Code)
}

func TestMiddleware_GivenOne_WhenAbortingWithExternalError_ThenGinAbortWith500ErrorIsCalled(t *testing.T) {
	gin.SetMode(gin.TestMode)
	recorder := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(recorder)
	middleware := this.CreateMiddleware(c)

	middleware.AbortTransactionWithError(errors.New("some error"))

	assert.Equal(t, 500, recorder.Code)
}

func TestMiddleware_GivenOne_WhenAbortingWith2xxError_ThenNothingHappens(t *testing.T) {
	gin.SetMode(gin.TestMode)
	recorder := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(recorder)
	middleware := this.CreateMiddleware(c)

	middleware.AbortTransactionWithError(http.Error{
		Code: 250,
	})

	assert.Equal(t, 200, recorder.Code)
}

func TestMiddleware_GivenOne_WhenAborting_ThenGinAbortIsCalledAndEndsPrematurely(t *testing.T) {
	gin.SetMode(gin.TestMode)
	recorder := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(recorder)
	middleware := this.CreateMiddleware(c)

	middleware.AbortTransaction()

	assert.Equal(t, 200, recorder.Code)
}

func TestMiddleware_GivenOne_WhenAbortingWithStatus_ThenGinAbortIsCalled(t *testing.T) {
	gin.SetMode(gin.TestMode)
	recorder := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(recorder)
	middleware := this.CreateMiddleware(c)

	middleware.AbortTransactionWithStatus(504, http.Json{})

	assert.Equal(t, 504, recorder.Code)
}

func TestMiddleware_GivenOne_WhenAborting_ThenItsAbortedAfterwards(t *testing.T) {
	gin.SetMode(gin.TestMode)
	c, _ := gin.CreateTestContext(httptest.NewRecorder())
	middleware := this.CreateMiddleware(c)

	middleware.AbortTransaction()

	assert.True(t, middleware.IsAborted())
}

func TestMiddleware_GivenOne_WhenGoingToNext_ThenItsFinishedBecauseItsTheLast(t *testing.T) {
	gin.SetMode(gin.TestMode)
	recorder := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(recorder)
	middleware := this.CreateMiddleware(c)

	middleware.NextHandler()

	assert.Equal(t, 200, recorder.Code)
}
