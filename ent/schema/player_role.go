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
		field.Int("match_id").
			StorageKey("role_match").
			Comment("foreign key to the matches table (Many-to-Many)"),

		field.Int("player_id").
			StorageKey("role_player").
			Comment("foreign key to the players table (Many-to-Many)"),

		field.Int("position").
			Range(1, 4).
			Comment("enumerated map position where base is located"),
		field.Int("turn_order").
			Range(1, 4).
			Comment("turn order is independent of base's map position"),
	}
}

// Edges of the PlayerRole.
func (PlayerRole) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("match", Match.Type).
			StorageKey(edge.Columns("role_match", "id")).
			Required(),

		edge.To("players", Player.Type).
			StorageKey(edge.Columns("role_player", "id")).
			Required(),
	}
}

func (PlayerRole) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("match_id", "position").Unique(),
		index.Fields("match_id", "turn_order").Unique(),
		index.Fields("match_id", "player_id").Unique(),
	}
}
