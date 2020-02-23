package validation_test

import (
	"github.com/arpb2/C-3PO/pkg/presentation/user/validation"
	"testing"

	"github.com/arpb2/C-3PO/pkg/domain/model"
	"github.com/stretchr/testify/assert"
)

func TestIdProvidedValidation_Error(t *testing.T) {
	err := validation.IdProvided(&model.AuthenticatedUser{
		User: model.User{
			Id: 1000,
		},
		Password: "",
	})

	assert.NotNil(t, err)
	assert.Equal(t, "client can't define user 'id'", err.Error())
}

func TestIdProvidedValidation_Success(t *testing.T) {
	err := validation.IdProvided(&model.AuthenticatedUser{
		User:     model.User{},
		Password: "",
	})

	assert.Nil(t, err)
}
