package validation_test

import (
	"testing"

	"github.com/arpb2/C-3PO/pkg/data/usecase/user/validation"
	"github.com/arpb2/C-3PO/pkg/domain/model/user"

	"github.com/stretchr/testify/assert"
)

func TestIdProvidedValidation_Error(t *testing.T) {
	err := validation.IdProvided(&user.AuthenticatedUser{
		User: user.User{
			Id: 1000,
		},
		Password: "",
	})

	assert.NotNil(t, err)
	assert.Equal(t, "client can't define user 'id'", err.Error())
}

func TestIdProvidedValidation_Success(t *testing.T) {
	err := validation.IdProvided(&user.AuthenticatedUser{
		User:     user.User{},
		Password: "",
	})

	assert.Nil(t, err)
}
