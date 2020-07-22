package level_test

import (
	"encoding/json"
	"testing"

	level2 "github.com/arpb2/C-3PO/pkg/domain/model/level"

	"github.com/stretchr/testify/assert"
)

func TestLevel_ToJson(t *testing.T) {
	expectedJson := `{"id":0,"name":"test name","description":"test description","definition":{"origin":{"position":{"x":0,"y":0},"orientation":""},"destination":{"position":{"x":0,"y":0},"active":false,"conditions":null},"minimal_dimensions":{"rows":0,"columns":0},"collectibles":null,"gates":null,"pads":null}}`

	level := &level2.Level{
		Id:          0,
		Name:        "test name",
		Description: "test description",
	}

	data, err := json.Marshal(level)

	assert.Nil(t, err)
	assert.Equal(t, expectedJson, string(data))
}

func TestLevel_FromJson(t *testing.T) {
	expectedLevel := &level2.Level{
		Id:          0,
		Name:        "test name",
		Description: "test description",
	}

	data := `{"id":0,"name":"test name","description":"test description","definition":{"origin":{"position":{"x":0,"y":0},"orientation":""},"destination":{"position":{"x":0,"y":0},"active":false,"conditions":null},"minimal_dimensions":{"rows":0,"columns":0},"collectibles":null,"gates":null,"pads":null}}`
	var level level2.Level

	err := json.Unmarshal([]byte(data), &level)

	assert.Nil(t, err)

	assert.Equal(t, expectedLevel.Id, level.Id)
	assert.Equal(t, expectedLevel.Name, level.Name)
	assert.Equal(t, expectedLevel.Description, level.Description)
}
