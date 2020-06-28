package schema

import (
	"time"

	"github.com/arpb2/C-3PO/pkg/domain/model/level"

	"github.com/facebookincubator/ent"
	"github.com/facebookincubator/ent/schema/field"
)

type LevelDefinition level.Definition

// Level holds the schema definition for the Level entity.
type Level struct {
	ent.Schema
}

func (Level) Fields() []ent.Field {
	return []ent.Field{
		field.Uint("id").Unique().Immutable(),
		field.Time("created_at").Default(time.Now).Immutable(),
		field.Time("updated_at").Default(time.Now).UpdateDefault(time.Now),
		field.String("name").NotEmpty(),
		field.Text("description").NotEmpty(),
		field.JSON("definition", &LevelDefinition{}),
	}
}

func (Level) Edges() []ent.Edge {
	return []ent.Edge{}
}
