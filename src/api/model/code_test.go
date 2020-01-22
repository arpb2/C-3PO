package model_test

import (
	"encoding/json"
	"github.com/arpb2/C-3PO/src/api/model"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCode_ToJson(t *testing.T) {
	expectedJson := `{"user_id":5,"id":1,"code":"some code written here"}`

	code := &model.Code{
		Id:     uint(1),
		UserId: uint(5),
		Code:   "some code written here",
	}

	data, err := json.Marshal(code)

	assert.Nil(t, err)
	assert.Equal(t, expectedJson, string(data))
}

func TestCode_FromJson(t *testing.T) {
	expectedCode := &model.Code{
		Id:     1,
		UserId: 5,
		Code:   "some code written here",
	}

	data := `{"user_id":5,"id":1,"code":"some code written here"}`
	var code model.Code

	err := json.Unmarshal([]byte(data), &code)

	assert.Nil(t, err)

	assert.Equal(t, expectedCode.Id, code.Id)
	assert.Equal(t, expectedCode.UserId, code.UserId)
	assert.Equal(t, expectedCode.Code, code.Code)
}