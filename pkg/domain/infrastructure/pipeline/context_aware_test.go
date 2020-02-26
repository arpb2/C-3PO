package pipeline_test

import (
	"testing"

	"github.com/arpb2/C-3PO/pkg/domain/infrastructure/pipeline"
	model3 "github.com/arpb2/C-3PO/pkg/domain/session/model"
	model2 "github.com/arpb2/C-3PO/pkg/domain/user/model"
	model4 "github.com/arpb2/C-3PO/pkg/domain/user_level/model"

	"github.com/arpb2/C-3PO/pkg/domain/infrastructure/http"
	http2 "github.com/arpb2/C-3PO/test/mock/http"
	gopipeline "github.com/saantiaguilera/go-pipeline"
	"github.com/stretchr/testify/assert"
)

func TestContextAware_GivenOneWithNoReader_WhenGettingIt_ThenItReturnsNilAndInternalError(t *testing.T) {
	ctx := gopipeline.CreateContext()
	ctxAware := pipeline.CreateContextAware(ctx)

	val, err := ctxAware.GetReader()

	assert.Nil(t, val)
	assert.Equal(t, http.CreateInternalError(), err)
}

func TestContextAware_GivenOneWithSomethingDifferentThanReader_WhenGettingIt_ThenItReturnsNilAndInternalError(t *testing.T) {
	ctx := gopipeline.CreateContext()
	ctx.Set(pipeline.TagHttpReader, "string")
	ctxAware := pipeline.CreateContextAware(ctx)

	val, err := ctxAware.GetReader()

	assert.Nil(t, val)
	assert.Equal(t, http.CreateInternalError(), err)
}

func TestContextAware_GivenOneWithAReader_WhenGettingIt_ThenItReturnsItAndNoError(t *testing.T) {
	expectedVal := &http2.MockReader{}
	ctx := gopipeline.CreateContext()
	ctx.Set(pipeline.TagHttpReader, expectedVal)
	ctxAware := pipeline.CreateContextAware(ctx)

	val, err := ctxAware.GetReader()

	assert.Equal(t, expectedVal, val)
	assert.Nil(t, err)
}

func TestContextAware_GivenOneWithNoWriter_WhenGettingIt_ThenItReturnsNilAndInternalError(t *testing.T) {
	ctx := gopipeline.CreateContext()
	ctxAware := pipeline.CreateContextAware(ctx)

	val, err := ctxAware.GetWriter()

	assert.Nil(t, val)
	assert.Equal(t, http.CreateInternalError(), err)
}

func TestContextAware_GivenOneWithSomethingDifferentThanWriter_WhenGettingIt_ThenItReturnsNilAndInternalError(t *testing.T) {
	ctx := gopipeline.CreateContext()
	ctx.Set(pipeline.TagHttpWriter, "string")
	ctxAware := pipeline.CreateContextAware(ctx)

	val, err := ctxAware.GetWriter()

	assert.Nil(t, val)
	assert.Equal(t, http.CreateInternalError(), err)
}

func TestContextAware_GivenOneWithAWriter_WhenGettingIt_ThenItReturnsItAndNoError(t *testing.T) {
	expectedVal := &http2.MockWriter{}
	ctx := gopipeline.CreateContext()
	ctx.Set(pipeline.TagHttpWriter, expectedVal)
	ctxAware := pipeline.CreateContextAware(ctx)

	val, err := ctxAware.GetWriter()

	assert.Equal(t, expectedVal, val)
	assert.Nil(t, err)
}

func TestContextAware_GivenOneWithNoMiddleware_WhenGettingIt_ThenItReturnsNilAndInternalError(t *testing.T) {
	ctx := gopipeline.CreateContext()
	ctxAware := pipeline.CreateContextAware(ctx)

	val, err := ctxAware.GetMiddleware()

	assert.Nil(t, val)
	assert.Equal(t, http.CreateInternalError(), err)
}

func TestContextAware_GivenOneWithSomethingDifferentThanMiddleware_WhenGettingIt_ThenItReturnsNilAndInternalError(t *testing.T) {
	ctx := gopipeline.CreateContext()
	ctx.Set(pipeline.TagHttpMiddleware, "string")
	ctxAware := pipeline.CreateContextAware(ctx)

	val, err := ctxAware.GetMiddleware()

	assert.Nil(t, val)
	assert.Equal(t, http.CreateInternalError(), err)
}

func TestContextAware_GivenOneWithAMiddleware_WhenGettingIt_ThenItReturnsItAndNoError(t *testing.T) {
	expectedVal := &http2.MockMiddleware{}
	ctx := gopipeline.CreateContext()
	ctx.Set(pipeline.TagHttpMiddleware, expectedVal)
	ctxAware := pipeline.CreateContextAware(ctx)

	val, err := ctxAware.GetMiddleware()

	assert.Equal(t, expectedVal, val)
	assert.Nil(t, err)
}

func TestContextAware_GivenOneWithNoAuthenticatedUser_WhenGettingIt_ThenItReturnsNilAndInternalError(t *testing.T) {
	ctx := gopipeline.CreateContext()
	ctxAware := pipeline.CreateContextAware(ctx)

	val, err := ctxAware.GetAuthenticatedUser("tag")

	assert.Equal(t, model2.AuthenticatedUser{}, val)
	assert.Equal(t, http.CreateInternalError(), err)
}

func TestContextAware_GivenOneWithSomethingDifferentThanAuthenticatedUser_WhenGettingIt_ThenItReturnsNilAndInternalError(t *testing.T) {
	ctx := gopipeline.CreateContext()
	ctx.Set("tag", "string")
	ctxAware := pipeline.CreateContextAware(ctx)

	val, err := ctxAware.GetAuthenticatedUser("tag")

	assert.Equal(t, model2.AuthenticatedUser{}, val)
	assert.Equal(t, http.CreateInternalError(), err)
}

func TestContextAware_GivenOneWithAAuthenticatedUser_WhenGettingIt_ThenItReturnsItAndNoError(t *testing.T) {
	expectedVal := model2.AuthenticatedUser{}
	ctx := gopipeline.CreateContext()
	ctx.Set("tag", expectedVal)
	ctxAware := pipeline.CreateContextAware(ctx)

	val, err := ctxAware.GetAuthenticatedUser("tag")

	assert.Equal(t, expectedVal, val)
	assert.Nil(t, err)
}

func TestContextAware_GivenOneWithNoUser_WhenGettingIt_ThenItReturnsNilAndInternalError(t *testing.T) {
	ctx := gopipeline.CreateContext()
	ctxAware := pipeline.CreateContextAware(ctx)

	val, err := ctxAware.GetUser("tag")

	assert.Equal(t, model2.User{}, val)
	assert.Equal(t, http.CreateInternalError(), err)
}

func TestContextAware_GivenOneWithSomethingDifferentThanUser_WhenGettingIt_ThenItReturnsNilAndInternalError(t *testing.T) {
	ctx := gopipeline.CreateContext()
	ctx.Set("tag", "string")
	ctxAware := pipeline.CreateContextAware(ctx)

	val, err := ctxAware.GetUser("tag")

	assert.Equal(t, model2.User{}, val)
	assert.Equal(t, http.CreateInternalError(), err)
}

func TestContextAware_GivenOneWithAUser_WhenGettingIt_ThenItReturnsItAndNoError(t *testing.T) {
	expectedVal := model2.User{}
	ctx := gopipeline.CreateContext()
	ctx.Set("tag", expectedVal)
	ctxAware := pipeline.CreateContextAware(ctx)

	val, err := ctxAware.GetUser("tag")

	assert.Equal(t, expectedVal, val)
	assert.Nil(t, err)
}

func TestContextAware_GivenOneWithNoUserLevel_WhenGettingIt_ThenItReturnsNilAndInternalError(t *testing.T) {
	ctx := gopipeline.CreateContext()
	ctxAware := pipeline.CreateContextAware(ctx)

	val, err := ctxAware.GetUserLevel("tag")

	assert.Equal(t, model4.UserLevel{}, val)
	assert.Equal(t, http.CreateInternalError(), err)
}

func TestContextAware_GivenOneWithSomethingDifferentThanUserLevel_WhenGettingIt_ThenItReturnsNilAndInternalError(t *testing.T) {
	ctx := gopipeline.CreateContext()
	ctx.Set("tag", "string")
	ctxAware := pipeline.CreateContextAware(ctx)

	val, err := ctxAware.GetUserLevel("tag")

	assert.Equal(t, model4.UserLevel{}, val)
	assert.Equal(t, http.CreateInternalError(), err)
}

func TestContextAware_GivenOneWithAUserLevel_WhenGettingIt_ThenItReturnsItAndNoError(t *testing.T) {
	expectedVal := model4.UserLevel{}
	ctx := gopipeline.CreateContext()
	ctx.Set("tag", expectedVal)
	ctxAware := pipeline.CreateContextAware(ctx)

	val, err := ctxAware.GetUserLevel("tag")

	assert.Equal(t, expectedVal, val)
	assert.Nil(t, err)
}

func TestContextAware_GivenOneWithNoSession_WhenGettingIt_ThenItReturnsNilAndInternalError(t *testing.T) {
	ctx := gopipeline.CreateContext()
	ctxAware := pipeline.CreateContextAware(ctx)

	val, err := ctxAware.GetSession("tag")

	assert.Equal(t, model3.Session{}, val)
	assert.Equal(t, http.CreateInternalError(), err)
}

func TestContextAware_GivenOneWithSomethingDifferentThanSession_WhenGettingIt_ThenItReturnsNilAndInternalError(t *testing.T) {
	ctx := gopipeline.CreateContext()
	ctx.Set("tag", "string")
	ctxAware := pipeline.CreateContextAware(ctx)

	val, err := ctxAware.GetSession("tag")

	assert.Equal(t, model3.Session{}, val)
	assert.Equal(t, http.CreateInternalError(), err)
}

func TestContextAware_GivenOneWithASession_WhenGettingIt_ThenItReturnsItAndNoError(t *testing.T) {
	expectedVal := model3.Session{}
	ctx := gopipeline.CreateContext()
	ctx.Set("tag", expectedVal)
	ctxAware := pipeline.CreateContextAware(ctx)

	val, err := ctxAware.GetSession("tag")

	assert.Equal(t, expectedVal, val)
	assert.Nil(t, err)
}

func TestContextAware_GivenOneWithNoUserLevelData_WhenGettingIt_ThenItReturnsNilAndInternalError(t *testing.T) {
	ctx := gopipeline.CreateContext()
	ctxAware := pipeline.CreateContextAware(ctx)

	val, err := ctxAware.GetUserLevelData("tag")

	assert.Equal(t, model4.UserLevelData{}, val)
	assert.Equal(t, http.CreateInternalError(), err)
}

func TestContextAware_GivenOneWithSomethingDifferentThanUserLevelData_WhenGettingIt_ThenItReturnsNilAndInternalError(t *testing.T) {
	ctx := gopipeline.CreateContext()
	ctx.Set("tag", "string")
	ctxAware := pipeline.CreateContextAware(ctx)

	val, err := ctxAware.GetUserLevelData("tag")

	assert.Equal(t, model4.UserLevelData{}, val)
	assert.Equal(t, http.CreateInternalError(), err)
}

func TestContextAware_GivenOneWithAUserLevelData_WhenGettingIt_ThenItReturnsItAndNoError(t *testing.T) {
	expectedVal := model4.UserLevelData{}
	ctx := gopipeline.CreateContext()
	ctx.Set("tag", expectedVal)
	ctxAware := pipeline.CreateContextAware(ctx)

	val, err := ctxAware.GetUserLevelData("tag")

	assert.Equal(t, expectedVal, val)
	assert.Nil(t, err)
}
