package authenticated_user_validation_test

import (
	"github.com/arpb2/C-3PO/src/api/model"
	"github.com/arpb2/C-3PO/src/api/validation/authenticated_user_validation"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestPasswordValidation_Error(t *testing.T) {
	err := authenticated_user_validation.PasswordValidation(&model.AuthenticatedUser{
		Password: "",
	})

	assert.NotNil(t, err)
	assert.Equal(t, "no 'password' provided", err.Error())
}


func TestPasswordValidation_Success(t *testing.T) {
	err := authenticated_user_validation.PasswordValidation(&model.AuthenticatedUser{
		Password: "test password",
	})

	assert.Nil(t, err)
}