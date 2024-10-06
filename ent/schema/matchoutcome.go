package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// MatchOutcome holds the schema definition for the MatchOutcome entity.
type MatchOutcome struct {
	ent.Schema
}

// Fields of the MatchOutcome.
func (MatchOutcome) Fields() []ent.Field {
	return []ent.Field{
		field.Enum("league").
			Values("UNRANKED",
				"Fluffy",
				"Clever",
				"Gifted",
				"Master",
				"Supertitan").
			Default("UNRANKED"),
		field.Int("rank").Positive(),
		field.Int("points").NonNegative(),
		field.Int("delta"),
	}
}

// Edges of the MatchOutcome.
func (MatchOutcome) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("prev", MatchOutcome.Type).
			Unique().
			From("next").
			Unique(),
		edge.To("role", PlayerRole.Type).
			Unique(),
	}
}
