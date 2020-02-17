package model_test

import (
	"encoding/json"
	"testing"

	"github.com/arpb2/C-3PO/api/model"
	"github.com/stretchr/testify/assert"
)

func TestCode_ToJson(t *testing.T) {
	expectedJson := `{"user_id":5,"level_id":1,"code":"some code written here","workspace":"test"}`

	code := &model.UserLevel{
		LevelId: uint(1),
		UserId:  uint(5),
		UserLevelData: model.UserLevelData{
			Code:      "some code written here",
			Workspace: "test",
		},
	}

	data, err := json.Marshal(code)

	assert.Nil(t, err)
	assert.Equal(t, expectedJson, string(data))
}

func TestCode_FromJson(t *testing.T) {
	expectedCode := &model.UserLevel{
		LevelId: 1,
		UserId:  5,
		UserLevelData: model.UserLevelData{
			Code:      "some code written here",
			Workspace: "test",
		},
	}

	data := `{"user_id":5,"level_id":1,"code":"some code written here","workspace":"test"}`
	var code model.UserLevel

	err := json.Unmarshal([]byte(data), &code)

	assert.Nil(t, err)

	assert.Equal(t, expectedCode.LevelId, code.LevelId)
	assert.Equal(t, expectedCode.UserId, code.UserId)
	assert.Equal(t, expectedCode.Code, code.Code)
}
