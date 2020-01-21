package authenticated_user_validation_test

import (
	"github.com/arpb2/C-3PO/src/api/model"
	"github.com/arpb2/C-3PO/src/api/validation/authenticated_user_validation"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestEmailValidation_Error(t *testing.T) {
	err := authenticated_user_validation.EmailValidation(&model.AuthenticatedUser{
		User: &model.User{
			Email: "",
		},
	})

	assert.NotNil(t, err)
	assert.Equal(t, "no 'email' provided", err.Error())
}


func TestEmailValidation_Success(t *testing.T) {
	err := authenticated_user_validation.EmailValidation(&model.AuthenticatedUser{
		User: &model.User{
			Email: "test@email.com",
		},
	})

	assert.Nil(t, err)
}