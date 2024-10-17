package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
)

// Arena holds the schema definition for the Arena entity.
type OsnMap struct {
	ent.Schema
}

// Fields of the Arena.
func (OsnMap) Fields() []ent.Field {
	return []ent.Field{
		field.Int("id"),
		field.String("map_slug"),
		field.String("name"),
		field.Int("role_count").Positive(),
		field.Int("width").Positive(),
		field.Int("height").Positive(),
	}
}

// Edges of the Arena.
func (OsnMap) Edges() []ent.Edge {
	return nil
}
