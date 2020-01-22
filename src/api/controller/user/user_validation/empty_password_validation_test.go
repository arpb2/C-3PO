package session_validation_test

import (
	"github.com/arpb2/C-3PO/src/api/controller/user/user_validation"
	"github.com/arpb2/C-3PO/src/api/model"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestEmptyPasswordValidation_Error(t *testing.T) {
	err := session_validation.EmptyPasswordValidation(&model.AuthenticatedUser{
		User:     &model.User{},
		Password: "",
	})

	assert.NotNil(t, err)
	assert.Equal(t, "no 'password' provided", err.Error())
}


func TestEmptyPasswordValidation_Success(t *testing.T) {
	err := session_validation.EmptyPasswordValidation(&model.AuthenticatedUser{
		User:     &model.User{},
		Password: "test password",
	})

	assert.Nil(t, err)
}

