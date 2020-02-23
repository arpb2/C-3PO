package validation_test

import (
	"github.com/arpb2/C-3PO/pkg/presentation/user/validation"
	"testing"

	"github.com/arpb2/C-3PO/pkg/domain/model"
	"github.com/stretchr/testify/assert"
)

func TestEmptyNameValidation_Error(t *testing.T) {
	err := validation.EmptyName(&model.AuthenticatedUser{
		User:     model.User{},
		Password: "",
	})

	assert.NotNil(t, err)
	assert.Equal(t, "no 'name' provided", err.Error())
}

func TestEmptyNameValidation_Success(t *testing.T) {
	err := validation.EmptyName(&model.AuthenticatedUser{
		User: model.User{
			Name: "TestName",
		},
		Password: "",
	})

	assert.Nil(t, err)
}
