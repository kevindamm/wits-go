// Code generated by ent, DO NOT EDIT.

package migrate

import (
	"entgo.io/ent/dialect/sql/schema"
	"entgo.io/ent/schema/field"
)

var (
	// MatchesColumns holds the columns for the "matches" table.
	MatchesColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "match_hash", Type: field.TypeString, Unique: true},
		{Name: "version", Type: field.TypeInt},
		{Name: "season", Type: field.TypeInt8, Default: 0},
		{Name: "created_ts", Type: field.TypeTime},
		{Name: "turn_count", Type: field.TypeInt},
		{Name: "fetch_status", Type: field.TypeEnum, Enums: []string{"UNKNOWN", "LISTED", "FETCHED", "UNWRAPPED", "CONVERTED", "CANONICAL", "VALIDATED", "INDEXED", "INVALID", "LEGACY"}, Default: "LISTED"},
		{Name: "osn_map_matches", Type: field.TypeInt, Nullable: true},
	}
	// MatchesTable holds the schema information for the "matches" table.
	MatchesTable = &schema.Table{
		Name:       "matches",
		Columns:    MatchesColumns,
		PrimaryKey: []*schema.Column{MatchesColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "matches_osn_maps_matches",
				Columns:    []*schema.Column{MatchesColumns[7]},
				RefColumns: []*schema.Column{OsnMapsColumns[0]},
				OnDelete:   schema.SetNull,
			},
		},
	}
	// OsnMapsColumns holds the columns for the "osn_maps" table.
	OsnMapsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "name", Type: field.TypeString},
		{Name: "shortname", Type: field.TypeString},
		{Name: "role_count", Type: field.TypeInt},
	}
	// OsnMapsTable holds the schema information for the "osn_maps" table.
	OsnMapsTable = &schema.Table{
		Name:       "osn_maps",
		Columns:    OsnMapsColumns,
		PrimaryKey: []*schema.Column{OsnMapsColumns[0]},
	}
	// PlayersColumns holds the columns for the "players" table.
	PlayersColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "gcid", Type: field.TypeString, Unique: true},
		{Name: "name", Type: field.TypeString, Unique: true},
	}
	// PlayersTable holds the schema information for the "players" table.
	PlayersTable = &schema.Table{
		Name:       "players",
		Columns:    PlayersColumns,
		PrimaryKey: []*schema.Column{PlayersColumns[0]},
	}
	// PlayerRolesColumns holds the columns for the "player_roles" table.
	PlayerRolesColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "role_match", Type: field.TypeInt},
		{Name: "role_player", Type: field.TypeInt},
		{Name: "position", Type: field.TypeInt},
		{Name: "turn_order", Type: field.TypeInt},
	}
	// PlayerRolesTable holds the schema information for the "player_roles" table.
	PlayerRolesTable = &schema.Table{
		Name:       "player_roles",
		Columns:    PlayerRolesColumns,
		PrimaryKey: []*schema.Column{PlayerRolesColumns[0]},
		Indexes: []*schema.Index{
			{
				Name:    "playerrole_role_match_position",
				Unique:  true,
				Columns: []*schema.Column{PlayerRolesColumns[1], PlayerRolesColumns[3]},
			},
			{
				Name:    "playerrole_role_match_role_player",
				Unique:  true,
				Columns: []*schema.Column{PlayerRolesColumns[1], PlayerRolesColumns[2]},
			},
		},
	}
	// PlayerRoleMatchColumns holds the columns for the "player_role_match" table.
	PlayerRoleMatchColumns = []*schema.Column{
		{Name: "role_match", Type: field.TypeInt},
		{Name: "id", Type: field.TypeInt},
	}
	// PlayerRoleMatchTable holds the schema information for the "player_role_match" table.
	PlayerRoleMatchTable = &schema.Table{
		Name:       "player_role_match",
		Columns:    PlayerRoleMatchColumns,
		PrimaryKey: []*schema.Column{PlayerRoleMatchColumns[0], PlayerRoleMatchColumns[1]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "player_role_match_role_match",
				Columns:    []*schema.Column{PlayerRoleMatchColumns[0]},
				RefColumns: []*schema.Column{PlayerRolesColumns[0]},
				OnDelete:   schema.Cascade,
			},
			{
				Symbol:     "player_role_match_id",
				Columns:    []*schema.Column{PlayerRoleMatchColumns[1]},
				RefColumns: []*schema.Column{MatchesColumns[0]},
				OnDelete:   schema.Cascade,
			},
		},
	}
	// PlayerRolePlayersColumns holds the columns for the "player_role_players" table.
	PlayerRolePlayersColumns = []*schema.Column{
		{Name: "role_player", Type: field.TypeInt},
		{Name: "id", Type: field.TypeInt},
	}
	// PlayerRolePlayersTable holds the schema information for the "player_role_players" table.
	PlayerRolePlayersTable = &schema.Table{
		Name:       "player_role_players",
		Columns:    PlayerRolePlayersColumns,
		PrimaryKey: []*schema.Column{PlayerRolePlayersColumns[0], PlayerRolePlayersColumns[1]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "player_role_players_role_player",
				Columns:    []*schema.Column{PlayerRolePlayersColumns[0]},
				RefColumns: []*schema.Column{PlayerRolesColumns[0]},
				OnDelete:   schema.Cascade,
			},
			{
				Symbol:     "player_role_players_id",
				Columns:    []*schema.Column{PlayerRolePlayersColumns[1]},
				RefColumns: []*schema.Column{PlayersColumns[0]},
				OnDelete:   schema.Cascade,
			},
		},
	}
	// Tables holds all the tables in the schema.
	Tables = []*schema.Table{
		MatchesTable,
		OsnMapsTable,
		PlayersTable,
		PlayerRolesTable,
		PlayerRoleMatchTable,
		PlayerRolePlayersTable,
	}
)

func init() {
	MatchesTable.ForeignKeys[0].RefTable = OsnMapsTable
	PlayerRoleMatchTable.ForeignKeys[0].RefTable = PlayerRolesTable
	PlayerRoleMatchTable.ForeignKeys[1].RefTable = MatchesTable
	PlayerRolePlayersTable.ForeignKeys[0].RefTable = PlayerRolesTable
	PlayerRolePlayersTable.ForeignKeys[1].RefTable = PlayersTable
}
