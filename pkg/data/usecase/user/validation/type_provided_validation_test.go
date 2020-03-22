package validation_test

import (
	"testing"

	"github.com/arpb2/C-3PO/pkg/data/usecase/user/validation"
	"github.com/arpb2/C-3PO/pkg/domain/model/user"

	"github.com/stretchr/testify/assert"
)

func TestTypeProvidedValidation_Error(t *testing.T) {
	err := validation.TypeProvided(&user.AuthenticatedUser{
		User: user.User{
			Type: "test",
		},
	})

	assert.NotNil(t, err)
	assert.Equal(t, "client can't define user 'type'", err.Error())
}

func TestTypeProvidedValidation_Success(t *testing.T) {
	err := validation.TypeProvided(&user.AuthenticatedUser{
		User: user.User{},
	})

	assert.Nil(t, err)
}
