// Copyright (c) 2024 Kevin Damm
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
//
// github:kevindamm/wits-go/witsjson/player.go

package witsjson

import (
	"encoding/json"
	"fmt"

	"github.com/kevindamm/wits-go/schema"
)

// Compatible with schema.PlayerRole interface, from a JSON-formatted replay.
type PlayerRoleJSON struct {
	PlayerID
	Name_   string              `json:"name"`
	Race_   UnitRaceJSON        `json:"race"`
	Team_   FriendlyEnumJSON    `json:"team"`
	Result_ TerminalStatusJSON  `json:"result"`
	Before_ PlayerStandingsJSON `json:"before"`
	After_  StandingsAfterJSON  `json:"after"`
	BaseHP_ BaseHealth          `json:"base_hp"`
	Wits_   int                 `json:"wits"`
}

func (role PlayerRoleJSON) Name() schema.PlayerName        { return schema.PlayerName(role.Name_) }
func (role PlayerRoleJSON) Race() schema.UnitRaceEnum      { return schema.UnitRaceEnum(role.Race_) }
func (role PlayerRoleJSON) Team() schema.FriendlyEnum      { return schema.FriendlyEnum(role.Team_) }
func (role PlayerRoleJSON) Result() schema.TerminalStatus  { return schema.TerminalStatus(role.Result_) }
func (role PlayerRoleJSON) Before() schema.PlayerStandings { return role.Before_ }
func (role PlayerRoleJSON) After() schema.StandingsAfter   { return role.After_ }
func (role PlayerRoleJSON) BaseHP() schema.BaseHealth      { return schema.BaseHealth(role.BaseHP_) }
func (role PlayerRoleJSON) Wits() schema.ActionPoints      { return schema.ActionPoints(role.Wits_) }

// May be inlined by other structs (see PlayerRoleJSON and player standings).
type PlayerID struct {
	GCID_ schema.GCID `json:"gcID"`
}

func (id PlayerID) GCID() schema.GCID { return id.GCID_ }

// A JSON-compatible representation wrapping the team-association enum.
type FriendlyEnumJSON schema.FriendlyEnum

func (team FriendlyEnumJSON) String() string {
	return map[schema.FriendlyEnum]string{
		schema.FR_SELF:    "RED",
		schema.FR_ENEMY:   "BLUE",
		schema.FR_ALLY:    "GOLD",
		schema.FR_ENEMY2:  "GREEN",
		schema.FR_UNKNOWN: "UNKNOWN",
	}[schema.FriendlyEnum(team)]
}

func ParseTeam(color string) FriendlyEnumJSON {
	return map[string]FriendlyEnumJSON{
		"RED":   FriendlyEnumJSON(schema.FR_SELF),
		"BLUE":  FriendlyEnumJSON(schema.FR_ENEMY),
		"GOLD":  FriendlyEnumJSON(schema.FR_ALLY),
		"GREEN": FriendlyEnumJSON(schema.FR_ENEMY2),
	}[color]
}

// Read and decode the JSON representation, accepting either integer or string.
// When reading from JSON, if it is an integer it is interpreted as 1-indexed
// sequence, and if it is a string then the canonical team coloring is used.
func (team *FriendlyEnumJSON) UnmarshalJSON(encoded []byte) error {
	var intVal int
	if err := json.Unmarshal(encoded, &intVal); err == nil {
		// The schema uses 0 (the default value) as UNKNOWN,
		// and OSN 0-indexed values are shifted when read.
		*team = FriendlyEnumJSON(schema.FriendlyEnum(intVal))
		return nil
	}
	var strVal string
	if err := json.Unmarshal(encoded, &strVal); err != nil {
		return err
	}
	*team = ParseTeam(strVal)
	return nil
}

func (team FriendlyEnumJSON) MarshalJSON() ([]byte, error) {
	encoded := team.String()
	if encoded == "UNKNOWN" {
		return []byte{}, fmt.Errorf("unknown team %d", byte(team))
	}
	return []byte(encoded), nil
}

// Player standings is the tier/rank of the player before or after the match.
type PlayerStandingsJSON struct {
	Tier_ LeagueTierJSON `json:"tier"`
	Rank_ LeagueRankJSON `json:"rank"`
}

type LeagueTierJSON schema.LeagueTier

func (standings PlayerStandingsJSON) Tier() schema.LeagueTier {
	return schema.LeagueTier(standings.Tier_)
}

type LeagueRankJSON schema.LeagueRank

func (standings PlayerStandingsJSON) Rank() schema.LeagueRank {
	return schema.LeagueRank(standings.Rank_)
}

type StandingsAfterJSON struct {
	Tier_  LeagueTierJSON `json:"tier"`
	Rank_  LeagueRankJSON `json:"rank"`
	Delta_ int            `json:"delta"`
}

func (standings StandingsAfterJSON) Tier() schema.LeagueTier {
	return schema.LeagueTier(standings.Tier_)
}

func (standings StandingsAfterJSON) Rank() schema.LeagueRank {
	return schema.LeagueRank(standings.Rank_)
}

func (standings StandingsAfterJSON) Delta() int {
	return standings.Delta_
}

// This value type has hard-coded limits of 0..5 checked when decoding.
type BaseHealth schema.BaseHealth

func (health *BaseHealth) UnmarshalJSON(encoded []byte) error {
	var hp int
	if err := json.Unmarshal(encoded, &hp); err != nil {
		return err
	}
	if hp < 0 || hp > 5 {
		return fmt.Errorf("invalid base HP: %d", hp)
	}

	*health = BaseHealth(schema.BaseHealth(hp))
	return nil
}

type TerminalStatusJSON schema.TerminalStatus
