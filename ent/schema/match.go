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
			Immutable().
			Unique().
			Comment("Identifier, a base-64 encoded hash-id of the match."),

		field.Int("version").
			Immutable().
			Comment("Which runtime version this match was played on.  Latest version is 1603."),

		field.Int8("season").
			Default(0).
			Comment("Zero value also implies non-competitive play."),

		field.Time("created_ts").
			Default(time.Now).
			Nillable().
			Comment("The timestamp when this match was recorded.  Nillable so it is not required in JSON responses."),

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
			Default("LISTED").
			Nillable().
			Comment("Where in the fetch/minify/convert process the replay is.  Nillable so that it is not required in JSON responses."),
	}
}

// Edges of the Match.
func (Match) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("map", OsnMap.Type).
			Ref("matches").
			Unique(),

		edge.From("roles", PlayerRole.Type).
			Ref("match"),
	}
}
