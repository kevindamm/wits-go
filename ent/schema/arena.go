package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
	"github.com/kevindamm/wits-go"
)

// Arena holds the schema definition for the Arena entity.
type Arena struct {
	ent.Schema
}

// Fields of the Arena.
func (Arena) Fields() []ent.Field {
	return []ent.Field{
		field.String("map_name"),
		field.Int("player_count").Positive(),
		field.JSON("details", &ArenaDetails{}),
	}
}

// TODO detail the typing of each terrain type
type ArenaDetails struct {
	Base  []any
	Bonus []any
	Floor []any
	Spawn []any
	Wall  []any
	Units []UnitGroup
}

// TODO relational data representation of each player's group of units
type UnitGroup interface {
	Count() int
	Get(int) wits.Unit
}

// Edges of the Arena.
func (Arena) Edges() []ent.Edge {
	return nil
}
