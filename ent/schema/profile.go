package schema

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/dialect"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// Profile holds the schema definition for the Profile entity.
type Profile struct {
	ent.Schema
}

// Fields of the Profile.
func (Profile) Fields() []ent.Field {
	return []ent.Field{
		field.String("id").MinLen(28).MaxLen(28).NotEmpty(), // varchar(28)
		field.String("name").MinLen(32).MaxLen(32).NotEmpty(),
		field.String("email").SchemaType(map[string]string{
			dialect.Postgres: "text",
		}).NotEmpty(),
		field.String("photo_url").SchemaType(map[string]string{
			dialect.Postgres: "text",
		}).Optional().Nillable(),
		field.Time("create_at").Default(time.Now).Immutable(),
		field.Time("update_at").Default(time.Now).UpdateDefault(time.Now),
		field.Time("delete_at").Optional(),
	}
}

// Edges of the Profile.
func (Profile) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("Tasks", Task.Type),
	}
}
