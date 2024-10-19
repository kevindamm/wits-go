// Code generated by ent, DO NOT EDIT.

package ent

import (
	"fmt"
	"strings"
	"time"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
	"github.com/kevindamm/wits-go/ent/match"
	"github.com/kevindamm/wits-go/ent/osnmap"
)

// Match is the model entity for the Match schema.
type Match struct {
	config `json:"-"`
	// ID of the ent.
	ID int `json:"id,omitempty"`
	// Identifier, a base-64 encoded hash-id of the match.
	MatchHash string `json:"match_hash,omitempty"`
	// Which runtime version this match was played on.  Latest version is 1603.
	Version int `json:"version,omitempty"`
	// Zero value also implies non-competitive play.
	Season int8 `json:"season,omitempty"`
	// The timestamp when this match was recorded.  Nillable so it is not required in JSON responses.
	CreatedTs *time.Time `json:"created_ts,omitempty"`
	// TurnCount holds the value of the "turn_count" field.
	TurnCount int `json:"turn_count,omitempty"`
	// Where in the fetch/minify/convert process the replay is.  Nillable so that it is not required in JSON responses.
	FetchStatus *match.FetchStatus `json:"fetch_status,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the MatchQuery when eager-loading is set.
	Edges           MatchEdges `json:"edges"`
	osn_map_matches *int
	selectValues    sql.SelectValues
}

// MatchEdges holds the relations/edges for other nodes in the graph.
type MatchEdges struct {
	// Map holds the value of the map edge.
	Map *OsnMap `json:"map,omitempty"`
	// Roles holds the value of the roles edge.
	Roles []*PlayerRole `json:"roles,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [2]bool
}

// MapOrErr returns the Map value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e MatchEdges) MapOrErr() (*OsnMap, error) {
	if e.Map != nil {
		return e.Map, nil
	} else if e.loadedTypes[0] {
		return nil, &NotFoundError{label: osnmap.Label}
	}
	return nil, &NotLoadedError{edge: "map"}
}

// RolesOrErr returns the Roles value or an error if the edge
// was not loaded in eager-loading.
func (e MatchEdges) RolesOrErr() ([]*PlayerRole, error) {
	if e.loadedTypes[1] {
		return e.Roles, nil
	}
	return nil, &NotLoadedError{edge: "roles"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*Match) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case match.FieldID, match.FieldVersion, match.FieldSeason, match.FieldTurnCount:
			values[i] = new(sql.NullInt64)
		case match.FieldMatchHash, match.FieldFetchStatus:
			values[i] = new(sql.NullString)
		case match.FieldCreatedTs:
			values[i] = new(sql.NullTime)
		case match.ForeignKeys[0]: // osn_map_matches
			values[i] = new(sql.NullInt64)
		default:
			values[i] = new(sql.UnknownType)
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the Match fields.
func (m *Match) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case match.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			m.ID = int(value.Int64)
		case match.FieldMatchHash:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field match_hash", values[i])
			} else if value.Valid {
				m.MatchHash = value.String
			}
		case match.FieldVersion:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field version", values[i])
			} else if value.Valid {
				m.Version = int(value.Int64)
			}
		case match.FieldSeason:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field season", values[i])
			} else if value.Valid {
				m.Season = int8(value.Int64)
			}
		case match.FieldCreatedTs:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field created_ts", values[i])
			} else if value.Valid {
				m.CreatedTs = new(time.Time)
				*m.CreatedTs = value.Time
			}
		case match.FieldTurnCount:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field turn_count", values[i])
			} else if value.Valid {
				m.TurnCount = int(value.Int64)
			}
		case match.FieldFetchStatus:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field fetch_status", values[i])
			} else if value.Valid {
				m.FetchStatus = new(match.FetchStatus)
				*m.FetchStatus = match.FetchStatus(value.String)
			}
		case match.ForeignKeys[0]:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for edge-field osn_map_matches", value)
			} else if value.Valid {
				m.osn_map_matches = new(int)
				*m.osn_map_matches = int(value.Int64)
			}
		default:
			m.selectValues.Set(columns[i], values[i])
		}
	}
	return nil
}

// Value returns the ent.Value that was dynamically selected and assigned to the Match.
// This includes values selected through modifiers, order, etc.
func (m *Match) Value(name string) (ent.Value, error) {
	return m.selectValues.Get(name)
}

// QueryMap queries the "map" edge of the Match entity.
func (m *Match) QueryMap() *OsnMapQuery {
	return NewMatchClient(m.config).QueryMap(m)
}

// QueryRoles queries the "roles" edge of the Match entity.
func (m *Match) QueryRoles() *PlayerRoleQuery {
	return NewMatchClient(m.config).QueryRoles(m)
}

// Update returns a builder for updating this Match.
// Note that you need to call Match.Unwrap() before calling this method if this Match
// was returned from a transaction, and the transaction was committed or rolled back.
func (m *Match) Update() *MatchUpdateOne {
	return NewMatchClient(m.config).UpdateOne(m)
}

// Unwrap unwraps the Match entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (m *Match) Unwrap() *Match {
	_tx, ok := m.config.driver.(*txDriver)
	if !ok {
		panic("ent: Match is not a transactional entity")
	}
	m.config.driver = _tx.drv
	return m
}

// String implements the fmt.Stringer.
func (m *Match) String() string {
	var builder strings.Builder
	builder.WriteString("Match(")
	builder.WriteString(fmt.Sprintf("id=%v, ", m.ID))
	builder.WriteString("match_hash=")
	builder.WriteString(m.MatchHash)
	builder.WriteString(", ")
	builder.WriteString("version=")
	builder.WriteString(fmt.Sprintf("%v", m.Version))
	builder.WriteString(", ")
	builder.WriteString("season=")
	builder.WriteString(fmt.Sprintf("%v", m.Season))
	builder.WriteString(", ")
	if v := m.CreatedTs; v != nil {
		builder.WriteString("created_ts=")
		builder.WriteString(v.Format(time.ANSIC))
	}
	builder.WriteString(", ")
	builder.WriteString("turn_count=")
	builder.WriteString(fmt.Sprintf("%v", m.TurnCount))
	builder.WriteString(", ")
	if v := m.FetchStatus; v != nil {
		builder.WriteString("fetch_status=")
		builder.WriteString(fmt.Sprintf("%v", *v))
	}
	builder.WriteByte(')')
	return builder.String()
}

// Matches is a parsable slice of Match.
type Matches []*Match
