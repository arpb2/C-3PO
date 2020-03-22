package schema

import (
	"time"

	"github.com/arpb2/C-3PO/pkg/domain/model/user"

	"github.com/facebookincubator/ent"
	"github.com/facebookincubator/ent/schema/edge"
	"github.com/facebookincubator/ent/schema/field"
)

// User holds the schema definition for the User entity.
type User struct {
	ent.Schema
}

func (User) Fields() []ent.Field {
	return []ent.Field{
		field.Uint("id").Unique().Immutable(),
		field.Enum("type").Values(
			string(user.TypeTeacher),
			string(user.TypeStudent),
		).Immutable(),
		field.String("email").NotEmpty().Unique(),
		field.String("name").NotEmpty(),
		field.String("surname").NotEmpty(),
		field.Time("created_at").Default(time.Now).Immutable(),
		field.Time("updated_at").Default(time.Now).UpdateDefault(time.Now),
	}
}

func (User) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("levels", UserLevel.Type),
		edge.To("credentials", Credential.Type).Unique(),
	}
}
