package model_test

import (
	"encoding/json"
	"github.com/arpb2/C-3PO/src/api/model"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestAuthenticatedUser_ToJson(t *testing.T) {
	expectedJson := `{"id":0,"email":"test@email.com","name":"test name","surname":"test surname","password":"test password"}`

	user := &model.AuthenticatedUser{
		User:       &model.User{
			Id:      0,
			Name:    "test name",
			Surname: "test surname",
			Email:   "test@email.com",
		},
		Password: "test password",
	}

	data, err := json.Marshal(user)

	assert.Nil(t, err)
	assert.Equal(t, expectedJson, string(data))
}

func TestAuthenticatedUser_FromJson(t *testing.T) {
	expectedUser := &model.AuthenticatedUser{
		User:       &model.User{
			Id:      0,
			Name:    "test name",
			Surname: "test surname",
			Email:   "test@email.com",
		},
		Password: "test password",
	}

	data := `{"id":0,"email":"test@email.com","name":"test name","surname":"test surname","password":"test password"}`
	var user model.AuthenticatedUser

	err := json.Unmarshal([]byte(data), &user)

	assert.Nil(t, err)

	assert.Equal(t, expectedUser.Id, user.Id)
	assert.Equal(t, expectedUser.Email, user.Email)
	assert.Equal(t, expectedUser.Name, user.Name)
	assert.Equal(t, expectedUser.Surname, user.Surname)
	assert.Equal(t, expectedUser.Password, user.Password)
}