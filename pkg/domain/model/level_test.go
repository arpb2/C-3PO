package model_test

import (
	"encoding/json"
	"testing"

	"github.com/arpb2/C-3PO/pkg/domain/model"
	"github.com/stretchr/testify/assert"
)

func TestLevel_ToJson(t *testing.T) {
	expectedJson := `{"id":0,"name":"test name","description":"test description"}`

	level := &model.Level{
		Id:          0,
		Name:        "test name",
		Description: "test description",
	}

	data, err := json.Marshal(level)

	assert.Nil(t, err)
	assert.Equal(t, expectedJson, string(data))
}

func TestLevel_FromJson(t *testing.T) {
	expectedLevel := &model.Level{
		Id:          0,
		Name:        "test name",
		Description: "test description",
	}

	data := `{"id":0,"name":"test name","description":"test description"}`
	var level model.Level

	err := json.Unmarshal([]byte(data), &level)

	assert.Nil(t, err)

	assert.Equal(t, expectedLevel.Id, level.Id)
	assert.Equal(t, expectedLevel.Name, level.Name)
	assert.Equal(t, expectedLevel.Description, level.Description)
}
