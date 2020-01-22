package user_validation_test

import (
	"github.com/arpb2/C-3PO/src/api/controller/user/user_validation"
	"github.com/arpb2/C-3PO/src/api/model"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSecurePasswordValidation_Error(t *testing.T) {
	err := user_validation.SecurePasswordValidation(&model.AuthenticatedUser{
		User:     &model.User{},
		Password: "1234567",
	})

	assert.NotNil(t, err)
	assert.Equal(t, "password must have at least 8 characters", err.Error())
}


func TestSecurePasswordValidation_Success(t *testing.T) {
	err := user_validation.SecurePasswordValidation(&model.AuthenticatedUser{
		User:     &model.User{},
		Password: "12345678",
	})

	assert.Nil(t, err)
}

