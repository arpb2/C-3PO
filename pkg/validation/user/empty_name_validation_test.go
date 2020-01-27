package user_test

import (
	"testing"

	"github.com/arpb2/C-3PO/api/model"
	uservalidation "github.com/arpb2/C-3PO/pkg/validation/user"
	"github.com/stretchr/testify/assert"
)

func TestEmptyNameValidation_Error(t *testing.T) {
	err := uservalidation.EmptyName(&model.AuthenticatedUser{
		User:     &model.User{},
		Password: "",
	})

	assert.NotNil(t, err)
	assert.Equal(t, "no 'name' provided", err.Error())
}

func TestEmptyNameValidation_Success(t *testing.T) {
	err := uservalidation.EmptyName(&model.AuthenticatedUser{
		User: &model.User{
			Name: "TestName",
		},
		Password: "",
	})

	assert.Nil(t, err)
}
