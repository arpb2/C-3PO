package session_validation_test

import (
	"github.com/arpb2/C-3PO/src/api/controller/user/user_validation"
	"github.com/arpb2/C-3PO/src/api/model"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestEmptyNameValidation_Error(t *testing.T) {
	err := session_validation.EmptyNameValidation(&model.AuthenticatedUser{
		User:     &model.User{},
		Password: "",
	})

	assert.NotNil(t, err)
	assert.Equal(t, "no 'name' provided", err.Error())
}


func TestEmptyNameValidation_Success(t *testing.T) {
	err := session_validation.EmptyNameValidation(&model.AuthenticatedUser{
		User:     &model.User{
			Name: "TestName",
		},
		Password: "",
	})

	assert.Nil(t, err)
}
