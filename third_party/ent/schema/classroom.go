package schema

import (
	"time"

	"github.com/facebookincubator/ent"
	"github.com/facebookincubator/ent/schema/edge"
	"github.com/facebookincubator/ent/schema/field"
)

// Classroom holds the schema definition for the Level entity.
type Classroom struct {
	ent.Schema
}

func (Classroom) Fields() []ent.Field {
	return []ent.Field{
		field.Uint("id").Unique().Immutable(),
		field.Time("created_at").Default(time.Now).Immutable(),
		field.Time("updated_at").Default(time.Now).UpdateDefault(time.Now),
	}
}

func (Classroom) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("teacher", User.Type).Required().Unique(),
		edge.To("students", User.Type),
		edge.To("level", Level.Type).Unique(),
	}
}
