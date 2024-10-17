package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/index"
)

// PlayerRole holds the schema definition for the PlayerRole entity.
type TeamRole struct {
	ent.Schema
}

// Fields of the TeamRole.
func (TeamRole) Fields() []ent.Field {
	return []ent.Field{
		field.Int("turn_order"),
	}
}

// Edges of the TeamRole.
func (TeamRole) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("match", Match.Type).
			Field("match_id"),
		edge.To("player_left", Player.Type).
			Field("player_id"),
		edge.To("player_right", Player.Type).
			Field("player_id"),
	}
}

func (TeamRole) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("match_id", "turn_order").
			Unique(),
		index.Fields("match_id", "player_id").
			Unique(),
	}
}
