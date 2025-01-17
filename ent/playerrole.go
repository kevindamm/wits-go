// Code generated by ent, DO NOT EDIT.

package ent

import (
	"fmt"
	"strings"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
	"github.com/kevindamm/wits-go/ent/playerrole"
)

// PlayerRole is the model entity for the PlayerRole schema.
type PlayerRole struct {
	config `json:"-"`
	// ID of the ent.
	ID int `json:"id,omitempty"`
	// MatchID holds the value of the "match_id" field.
	MatchID int `json:"match_id,omitempty"`
	// PlayerID holds the value of the "player_id" field.
	PlayerID int `json:"player_id,omitempty"`
	// Position holds the value of the "position" field.
	Position int `json:"position,omitempty"`
	// TurnOrder holds the value of the "turn_order" field.
	TurnOrder int `json:"turn_order,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the PlayerRoleQuery when eager-loading is set.
	Edges        PlayerRoleEdges `json:"edges"`
	selectValues sql.SelectValues
}

// PlayerRoleEdges holds the relations/edges for other nodes in the graph.
type PlayerRoleEdges struct {
	// Match holds the value of the match edge.
	Match []*Match `json:"match,omitempty"`
	// Players holds the value of the players edge.
	Players []*Player `json:"players,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [2]bool
}

// MatchOrErr returns the Match value or an error if the edge
// was not loaded in eager-loading.
func (e PlayerRoleEdges) MatchOrErr() ([]*Match, error) {
	if e.loadedTypes[0] {
		return e.Match, nil
	}
	return nil, &NotLoadedError{edge: "match"}
}

// PlayersOrErr returns the Players value or an error if the edge
// was not loaded in eager-loading.
func (e PlayerRoleEdges) PlayersOrErr() ([]*Player, error) {
	if e.loadedTypes[1] {
		return e.Players, nil
	}
	return nil, &NotLoadedError{edge: "players"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*PlayerRole) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case playerrole.FieldID, playerrole.FieldMatchID, playerrole.FieldPlayerID, playerrole.FieldPosition, playerrole.FieldTurnOrder:
			values[i] = new(sql.NullInt64)
		default:
			values[i] = new(sql.UnknownType)
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the PlayerRole fields.
func (pr *PlayerRole) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case playerrole.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			pr.ID = int(value.Int64)
		case playerrole.FieldMatchID:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field match_id", values[i])
			} else if value.Valid {
				pr.MatchID = int(value.Int64)
			}
		case playerrole.FieldPlayerID:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field player_id", values[i])
			} else if value.Valid {
				pr.PlayerID = int(value.Int64)
			}
		case playerrole.FieldPosition:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field position", values[i])
			} else if value.Valid {
				pr.Position = int(value.Int64)
			}
		case playerrole.FieldTurnOrder:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field turn_order", values[i])
			} else if value.Valid {
				pr.TurnOrder = int(value.Int64)
			}
		default:
			pr.selectValues.Set(columns[i], values[i])
		}
	}
	return nil
}

// Value returns the ent.Value that was dynamically selected and assigned to the PlayerRole.
// This includes values selected through modifiers, order, etc.
func (pr *PlayerRole) Value(name string) (ent.Value, error) {
	return pr.selectValues.Get(name)
}

// QueryMatch queries the "match" edge of the PlayerRole entity.
func (pr *PlayerRole) QueryMatch() *MatchQuery {
	return NewPlayerRoleClient(pr.config).QueryMatch(pr)
}

// QueryPlayers queries the "players" edge of the PlayerRole entity.
func (pr *PlayerRole) QueryPlayers() *PlayerQuery {
	return NewPlayerRoleClient(pr.config).QueryPlayers(pr)
}

// Update returns a builder for updating this PlayerRole.
// Note that you need to call PlayerRole.Unwrap() before calling this method if this PlayerRole
// was returned from a transaction, and the transaction was committed or rolled back.
func (pr *PlayerRole) Update() *PlayerRoleUpdateOne {
	return NewPlayerRoleClient(pr.config).UpdateOne(pr)
}

// Unwrap unwraps the PlayerRole entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (pr *PlayerRole) Unwrap() *PlayerRole {
	_tx, ok := pr.config.driver.(*txDriver)
	if !ok {
		panic("ent: PlayerRole is not a transactional entity")
	}
	pr.config.driver = _tx.drv
	return pr
}

// String implements the fmt.Stringer.
func (pr *PlayerRole) String() string {
	var builder strings.Builder
	builder.WriteString("PlayerRole(")
	builder.WriteString(fmt.Sprintf("id=%v, ", pr.ID))
	builder.WriteString("match_id=")
	builder.WriteString(fmt.Sprintf("%v", pr.MatchID))
	builder.WriteString(", ")
	builder.WriteString("player_id=")
	builder.WriteString(fmt.Sprintf("%v", pr.PlayerID))
	builder.WriteString(", ")
	builder.WriteString("position=")
	builder.WriteString(fmt.Sprintf("%v", pr.Position))
	builder.WriteString(", ")
	builder.WriteString("turn_order=")
	builder.WriteString(fmt.Sprintf("%v", pr.TurnOrder))
	builder.WriteByte(')')
	return builder.String()
}

// PlayerRoles is a parsable slice of PlayerRole.
type PlayerRoles []*PlayerRole
