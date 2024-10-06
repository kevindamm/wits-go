package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/index"
)

// PlayerRole holds the schema definition for the PlayerRole entity.
type PlayerRole struct {
	ent.Schema
}

// Fields of the PlayerRole.
func (PlayerRole) Fields() []ent.Field {
	return []ent.Field{
		field.Int("turn_order"),
	}
}

// Edges of the PlayerRole.
func (PlayerRole) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("match", Match.Type).
			Field("match_id"),
		edge.From("player", Player.Type).
			Ref("roles").
			Field("player_id"),
		edge.From("outcome", MatchOutcome.Type).
			Ref("role"),
	}
}

func (PlayerRole) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("match_id", "turn_order").
			Unique(),
		index.Fields("match_id", "player_id").
			Unique(),
	}
}
