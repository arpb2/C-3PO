package user_validation_test

import (
	"testing"

	"github.com/arpb2/C-3PO/api/model"
	user_validation "github.com/arpb2/C-3PO/pkg/validation/user"
	"github.com/stretchr/testify/assert"
)

func TestEmptySurnameValidation_Error(t *testing.T) {
	err := user_validation.EmptySurname(&model.AuthenticatedUser{
		User:     &model.User{},
		Password: "",
	})

	assert.NotNil(t, err)
	assert.Equal(t, "no 'surname' provided", err.Error())
}

func TestEmptySurnameValidation_Success(t *testing.T) {
	err := user_validation.EmptySurname(&model.AuthenticatedUser{
		User: &model.User{
			Surname: "TestSurname",
		},
		Password: "",
	})

	assert.Nil(t, err)
}
