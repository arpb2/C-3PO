package level_test

import (
	"encoding/json"
	"testing"

	level2 "github.com/arpb2/C-3PO/pkg/domain/model/level"

	"github.com/stretchr/testify/assert"
)

func TestLevelDefinition_ToJson(t *testing.T) {
	expectedJson := `{"origin":{"position":{"x":1,"y":2},"orientation":"E"},"destination":{"position":{"x":2,"y":3},"active":true,"conditions":[{"type":"coin","value":3},{"type":"colour","value":"blue"},{"type":"key","value":true}]},"minimal_dimensions":{"rows":1,"columns":2},"collectibles":[{"position":{"x":1,"y":2},"type":"coin"},{"position":{"x":4,"y":3},"type":"key"}],"gates":[{"position":{"x":1,"y":2,"axis":"x"},"type":"key"},{"position":{"x":1,"y":8,"axis":"y"},"type":"verbal","opening_value":"holis"}],"pads":[{"position":{"x":1,"y":2},"type":"colour"}]}`

	definition := &level2.Definition{
		Origin: level2.Origin{
			Position: level2.Position{
				X: 1,
				Y: 2,
			},
			Orientation: level2.OrientationEast,
		},
		Destination: level2.Destination{
			Position: level2.Position{
				X: 2,
				Y: 3,
			},
			Active: true,
			Conditions: []level2.Condition{
				{
					Type:  "coin",
					Value: float64(3),
				},
				{
					Type:  "colour",
					Value: "blue",
				},
				{
					Type:  "key",
					Value: true,
				},
			},
		},
		MinimalDimension: level2.Dimension{
			Rows:    1,
			Columns: 2,
		},
		Collectibles: []level2.Element{
			{
				Position: level2.Position{
					X: 1,
					Y: 2,
				},
				Type: "coin",
			},
			{
				Position: level2.Position{
					X: 4,
					Y: 3,
				},
				Type: "key",
			},
		},
		Gates: []level2.Gate{
			{
				Element: level2.Element{
					Position: level2.Position{
						X:    1,
						Y:    2,
						Axis: level2.AxisX,
					},
					Type: "key",
				},
			},
			{
				Element: level2.Element{
					Position: level2.Position{
						X:    1,
						Y:    8,
						Axis: level2.AxisY,
					},
					Type: "verbal",
				},
				OpeningValue: "holis",
			},
		},
		Pads: []level2.Element{
			{
				Position: level2.Position{
					X: 1,
					Y: 2,
				},
				Type: "colour",
			},
		},
	}

	data, err := json.Marshal(definition)

	assert.Nil(t, err)
	assert.Equal(t, expectedJson, string(data))
}

func TestLevelDefinition_FromJson(t *testing.T) {
	data := `{"origin":{"position":{"x":1,"y":2},"orientation":"E"},"destination":{"position":{"x":2,"y":3},"active":true,"conditions":[{"type":"coin","value":3},{"type":"colour","value":"blue"},{"type":"key","value":true}]},"minimal_dimensions":{"rows":1,"columns":2},"collectibles":[{"position":{"x":1,"y":2},"type":"coin"},{"position":{"x":4,"y":3},"type":"key"}],"gates":[{"position":{"x":1,"y":2,"axis":"x"},"type":"key"},{"position":{"x":1,"y":8,"axis":"y"},"type":"verbal","opening_value":"holis"}],"pads":[{"position":{"x":1,"y":2},"type":"colour"}]}`
	expectedDefinition := level2.Definition{
		Origin: level2.Origin{
			Position: level2.Position{
				X: 1,
				Y: 2,
			},
			Orientation: level2.OrientationEast,
		},
		Destination: level2.Destination{
			Position: level2.Position{
				X: 2,
				Y: 3,
			},
			Active: true,
			Conditions: []level2.Condition{
				{
					Type:  "coin",
					Value: float64(3),
				},
				{
					Type:  "colour",
					Value: "blue",
				},
				{
					Type:  "key",
					Value: true,
				},
			},
		},
		MinimalDimension: level2.Dimension{
			Rows:    1,
			Columns: 2,
		},
		Collectibles: []level2.Element{
			{
				Position: level2.Position{
					X: 1,
					Y: 2,
				},
				Type: "coin",
			},
			{
				Position: level2.Position{
					X: 4,
					Y: 3,
				},
				Type: "key",
			},
		},
		Gates: []level2.Gate{
			{
				Element: level2.Element{
					Position: level2.Position{
						X:    1,
						Y:    2,
						Axis: level2.AxisX,
					},
					Type: "key",
				},
			},
			{
				Element: level2.Element{
					Position: level2.Position{
						X:    1,
						Y:    8,
						Axis: level2.AxisY,
					},
					Type: "verbal",
				},
				OpeningValue: "holis",
			},
		},
		Pads: []level2.Element{
			{
				Position: level2.Position{
					X: 1,
					Y: 2,
				},
				Type: "colour",
			},
		},
	}
	var definition level2.Definition

	err := json.Unmarshal([]byte(data), &definition)

	assert.Nil(t, err)

	assert.Equal(t, expectedDefinition, definition)
}
