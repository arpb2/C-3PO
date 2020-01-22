package session_validation_test

import (
	"github.com/arpb2/C-3PO/src/api/controller/user/user_validation"
	"github.com/arpb2/C-3PO/src/api/model"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestIdProvidedValidation_Error(t *testing.T) {
	err := session_validation.IdProvidedValidation(&model.AuthenticatedUser{
		User:     &model.User{
			Id: 1000,
		},
		Password: "",
	})

	assert.NotNil(t, err)
	assert.Equal(t, "client can't define user 'id'", err.Error())
}


func TestIdProvidedValidation_Success(t *testing.T) {
	err := session_validation.IdProvidedValidation(&model.AuthenticatedUser{
		User:     &model.User{},
		Password: "",
	})

	assert.Nil(t, err)
}

