package authenticated_user_validation_test

import (
	"github.com/arpb2/C-3PO/src/api/model"
	"github.com/arpb2/C-3PO/src/api/validation/authenticated_user_validation"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestUserValidation_Error(t *testing.T) {
	err := authenticated_user_validation.UserValidation(nil)

	assert.NotNil(t, err)
	assert.Equal(t, "no user found", err.Error())
}


func TestUserValidation_Success(t *testing.T) {
	err := authenticated_user_validation.UserValidation(&model.AuthenticatedUser{})

	assert.Nil(t, err)
}