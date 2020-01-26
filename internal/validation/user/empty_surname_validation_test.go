package user_validation_test

import (
	"github.com/arpb2/C-3PO/api/model"
	"github.com/arpb2/C-3PO/internal/validation/user"
	"github.com/stretchr/testify/assert"
	"testing"
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
