package schema

import (
	"time"

	"github.com/facebookincubator/ent"
	"github.com/facebookincubator/ent/schema/edge"
	"github.com/facebookincubator/ent/schema/field"
	"github.com/facebookincubator/ent/schema/index"
)

// UserLevel holds the schema definition for the UserLevel entity.
type UserLevel struct {
	ent.Schema
}

func (UserLevel) Fields() []ent.Field {
	return []ent.Field{
		field.Time("created_at").Default(time.Now).Immutable(),
		field.Time("updated_at").Default(time.Now).UpdateDefault(time.Now),
		field.Text("code").NotEmpty(),
		field.Text("workspace").NotEmpty(),
	}
}

func (UserLevel) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("developer", User.Type).Unique().Required(),
		edge.To("level", Level.Type).Unique().Required(),
	}
}

func (UserLevel) Indexes() []ent.Index {
	return []ent.Index{
		index.Edges("developer", "level").Unique(),
	}
}
