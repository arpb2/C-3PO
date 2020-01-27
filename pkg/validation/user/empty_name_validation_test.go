package user_validation_test

import (
	"github.com/arpb2/C-3PO/api/model"
	"github.com/arpb2/C-3PO/pkg/validation/user"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestEmptyNameValidation_Error(t *testing.T) {
	err := user_validation.EmptyName(&model.AuthenticatedUser{
		User:     &model.User{},
		Password: "",
	})

	assert.NotNil(t, err)
	assert.Equal(t, "no 'name' provided", err.Error())
}

func TestEmptyNameValidation_Success(t *testing.T) {
	err := user_validation.EmptyName(&model.AuthenticatedUser{
		User: &model.User{
			Name: "TestName",
		},
		Password: "",
	})

	assert.Nil(t, err)
}
