package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// Arena holds the schema definition for the Arena entity.
type OsnMap struct {
	ent.Schema
}

// Fields of the Arena.
func (OsnMap) Fields() []ent.Field {
	return []ent.Field{
		field.String("name"),
		field.String("shortname"),
		field.Int("role_count").Positive(),
	}
}

// Edges of the Arena.
func (OsnMap) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("matches", Match.Type),
	}
}
