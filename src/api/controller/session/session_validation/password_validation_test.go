package session_validation_test

import (
	"github.com/arpb2/C-3PO/src/api/controller/session/session_validation"
	"github.com/arpb2/C-3PO/src/api/model"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestPasswordValidation_Error(t *testing.T) {
	err := session_validation.PasswordValidation(&model.AuthenticatedUser{
		Password: "",
	})

	assert.NotNil(t, err)
	assert.Equal(t, "no 'password' provided", err.Error())
}


func TestPasswordValidation_Success(t *testing.T) {
	err := session_validation.PasswordValidation(&model.AuthenticatedUser{
		Password: "test password",
	})

	assert.Nil(t, err)
}