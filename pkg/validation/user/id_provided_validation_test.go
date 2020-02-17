package user_test

import (
	"testing"

	"github.com/arpb2/C-3PO/api/model"
	uservalidation "github.com/arpb2/C-3PO/pkg/validation/user"
	"github.com/stretchr/testify/assert"
)

func TestIdProvidedValidation_Error(t *testing.T) {
	err := uservalidation.IdProvided(&model.AuthenticatedUser{
		User: model.User{
			Id: 1000,
		},
		Password: "",
	})

	assert.NotNil(t, err)
	assert.Equal(t, "client can't define user 'id'", err.Error())
}

func TestIdProvidedValidation_Success(t *testing.T) {
	err := uservalidation.IdProvided(&model.AuthenticatedUser{
		User:     model.User{},
		Password: "",
	})

	assert.Nil(t, err)
}
