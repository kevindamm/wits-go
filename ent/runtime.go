// Code generated by ent, DO NOT EDIT.

package ent

import (
	"time"

	"github.com/kevindamm/wits-go/ent/match"
	"github.com/kevindamm/wits-go/ent/osnmap"
	"github.com/kevindamm/wits-go/ent/playerrole"
	"github.com/kevindamm/wits-go/ent/schema"
)

// The init function reads all schema descriptors with runtime code
// (default values, validators, hooks and policies) and stitches it
// to their package variables.
func init() {
	matchFields := schema.Match{}.Fields()
	_ = matchFields
	// matchDescSeason is the schema descriptor for season field.
	matchDescSeason := matchFields[2].Descriptor()
	// match.DefaultSeason holds the default value on creation for the season field.
	match.DefaultSeason = matchDescSeason.Default.(int8)
	// matchDescCreatedTs is the schema descriptor for created_ts field.
	matchDescCreatedTs := matchFields[3].Descriptor()
	// match.DefaultCreatedTs holds the default value on creation for the created_ts field.
	match.DefaultCreatedTs = matchDescCreatedTs.Default.(func() time.Time)
	osnmapFields := schema.OsnMap{}.Fields()
	_ = osnmapFields
	// osnmapDescRoleCount is the schema descriptor for role_count field.
	osnmapDescRoleCount := osnmapFields[2].Descriptor()
	// osnmap.RoleCountValidator is a validator for the "role_count" field. It is called by the builders before save.
	osnmap.RoleCountValidator = osnmapDescRoleCount.Validators[0].(func(int) error)
	playerroleFields := schema.PlayerRole{}.Fields()
	_ = playerroleFields
	// playerroleDescPosition is the schema descriptor for position field.
	playerroleDescPosition := playerroleFields[2].Descriptor()
	// playerrole.PositionValidator is a validator for the "position" field. It is called by the builders before save.
	playerrole.PositionValidator = playerroleDescPosition.Validators[0].(func(int) error)
	// playerroleDescTurnOrder is the schema descriptor for turn_order field.
	playerroleDescTurnOrder := playerroleFields[3].Descriptor()
	// playerrole.TurnOrderValidator is a validator for the "turn_order" field. It is called by the builders before save.
	playerrole.TurnOrderValidator = playerroleDescTurnOrder.Validators[0].(func(int) error)
}