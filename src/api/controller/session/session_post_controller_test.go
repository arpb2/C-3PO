package session_test

import (
	"github.com/arpb2/C-3PO/src/api/controller/session"
	"github.com/arpb2/C-3PO/src/api/http_wrapper/gin_wrapper"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestHealthControllerMethodIsGET(t *testing.T) {
	assert.Equal(t, "POST", session.CreatePostController().Method)
}

func TestHealthControllerPathIsPing(t *testing.T) {
	assert.Equal(t, "/session", session.CreatePostController().Path)
}

func TestHealthControllerBodyReturnsWith200IfItsOk(t *testing.T) {
	c, w := gin_wrapper.CreateTestContext()

	session.CreatePostController().Body(c)

	assert.Equal(t, 200, w.Code)
}
