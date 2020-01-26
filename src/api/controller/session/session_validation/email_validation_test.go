package session_validation_test

import (
	"github.com/arpb2/C-3PO/src/api/controller/session/session_validation"
	"github.com/arpb2/C-3PO/src/api/model"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestEmailValidation_Error(t *testing.T) {
	err := session_validation.EmailValidation(&model.AuthenticatedUser{
		User: &model.User{
			Email: "",
		},
	})

	assert.NotNil(t, err)
	assert.Equal(t, "no 'email' provided", err.Error())
}


func TestEmailValidation_Success(t *testing.T) {
	err := session_validation.EmailValidation(&model.AuthenticatedUser{
		User: &model.User{
			Email: "test@email.com",
		},
	})

	assert.Nil(t, err)
}