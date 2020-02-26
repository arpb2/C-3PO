package validation_test

import (
	"testing"

	model2 "github.com/arpb2/C-3PO/pkg/domain/user/model"

	"github.com/arpb2/C-3PO/pkg/presentation/user/validation"

	"github.com/stretchr/testify/assert"
)

func TestIdProvidedValidation_Error(t *testing.T) {
	err := validation.IdProvided(&model2.AuthenticatedUser{
		User: model2.User{
			Id: 1000,
		},
		Password: "",
	})

	assert.NotNil(t, err)
	assert.Equal(t, "client can't define user 'id'", err.Error())
}

func TestIdProvidedValidation_Success(t *testing.T) {
	err := validation.IdProvided(&model2.AuthenticatedUser{
		User:     model2.User{},
		Password: "",
	})

	assert.Nil(t, err)
}
