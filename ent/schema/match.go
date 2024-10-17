package schema

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// Match holds the schema definition for the Match entity.
type Match struct {
	ent.Schema
}

// Fields of the Match.
func (Match) Fields() []ent.Field {
	return []ent.Field{
		field.String("match_hash").
			Unique(),
		field.Int("version"),
		field.Int8("season").
			Default(0), // Zero value also implies non-competitive.
		field.Time("created_ts").
			Default(time.Now),
		field.Int("turn_count"),
		field.Enum("fetch_status").
			Values(
				"UNKNOWN",
				"LISTED",
				"FETCHED",
				"UNWRAPPED",
				"CONVERTED",
				"CANONICAL",
				"VALIDATED",
				"INDEXED",
				"INVALID",
				"LEGACY").
			Default("LISTED"),
	}
}

// Edges of the Match.
func (Match) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("map", OsnMap.Type),
		edge.From("roles", PlayerRole.Type).Ref("match"),
		edge.From("roles", TeamRole.Type).Ref("match"),
	}
}
