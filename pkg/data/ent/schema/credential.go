package schema

import (
	"github.com/facebookincubator/ent"
	"github.com/facebookincubator/ent/schema/edge"
	"github.com/facebookincubator/ent/schema/field"
	"github.com/facebookincubator/ent/schema/index"
)

// Credential holds the schema definition for the Credential entity.
type Credential struct {
	ent.Schema
}

func (Credential) Fields() []ent.Field {
	return []ent.Field{
		field.Bytes("salt").Immutable(),
		field.Bytes("password_hash"),
	}
}

func (Credential) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("holder", User.Type).Ref("credentials").
			Required().
			Unique(),
	}
}

func (Credential) Indexes() []ent.Index {
	return []ent.Index{
		index.Edges("holder").Unique(),
	}
}
