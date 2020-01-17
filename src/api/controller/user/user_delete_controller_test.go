package user_test

import (
	"github.com/arpb2/C-3PO/src/api/controller/user"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestUserDeleteControllerMethodIsGET(t *testing.T) {
	assert.Equal(t, "DELETE", user.CreateDeleteController().Method)
}

func TestUserDeleteControllerPathIsAsExpected(t *testing.T) {
	assert.Equal(t, "/users/:user_id", user.CreateDeleteController().Path)
}