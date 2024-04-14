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

	"github.com/kevindamm/wits-go"
)

// Compatible with wits.PlayerRole interface, from a JSON-formatted replay.
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

func (role PlayerRoleJSON) Name() wits.PlayerName             { return wits.PlayerName(role.Name_) }
func (role PlayerRoleJSON) Race() wits.UnitRaceEnum           { return wits.UnitRaceEnum(role.Race_) }
func (role PlayerRoleJSON) Team() wits.FriendlyEnum           { return wits.FriendlyEnum(role.Team_) }
func (role PlayerRoleJSON) Result() wits.TerminalStatus       { return wits.TerminalStatus(role.Result_) }
func (role PlayerRoleJSON) Before() wits.PlayerStandings      { return role.Before_ }
func (role PlayerRoleJSON) After() wits.PlayerStandingsUpdate { return role.After_ }
func (role PlayerRoleJSON) BaseHP() wits.BaseHealth           { return wits.BaseHealth(role.BaseHP_) }
func (role PlayerRoleJSON) Wits() wits.ActionPoints           { return wits.ActionPoints(role.Wits_) }

// May be inlined by other structs (see PlayerRoleJSON and player standings).
type PlayerID struct {
	GCID_ wits.GCID `json:"gcID"`
}

func (id PlayerID) GCID() wits.GCID { return id.GCID_ }

// A JSON-compatible representation wrapping the team-association enum.
type FriendlyEnumJSON wits.FriendlyEnum

func (team FriendlyEnumJSON) String() string {
	return map[wits.FriendlyEnum]string{
		wits.FR_SELF:    "RED",
		wits.FR_ENEMY:   "BLUE",
		wits.FR_ALLY:    "GOLD",
		wits.FR_ENEMY2:  "GREEN",
		wits.FR_UNKNOWN: "UNKNOWN",
	}[wits.FriendlyEnum(team)]
}

func ParseTeam(color string) FriendlyEnumJSON {
	return map[string]FriendlyEnumJSON{
		"RED":   FriendlyEnumJSON(wits.FR_SELF),
		"BLUE":  FriendlyEnumJSON(wits.FR_ENEMY),
		"GOLD":  FriendlyEnumJSON(wits.FR_ALLY),
		"GREEN": FriendlyEnumJSON(wits.FR_ENEMY2),
	}[color]
}

// Read and decode the JSON representation, accepting either integer or string.
// When reading from JSON, if it is an integer it is interpreted as 1-indexed
// sequence, and if it is a string then the canonical team coloring is used.
func (team *FriendlyEnumJSON) UnmarshalJSON(encoded []byte) error {
	var intVal int
	if err := json.Unmarshal(encoded, &intVal); err == nil {
		// The wits uses 0 (the default value) as UNKNOWN,
		// and OSN 0-indexed values are shifted when read.
		*team = FriendlyEnumJSON(wits.FriendlyEnum(intVal))
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

type LeagueTierJSON wits.LeagueTier

func (standings PlayerStandingsJSON) Tier() wits.LeagueTier {
	return wits.LeagueTier(standings.Tier_)
}

type LeagueRankJSON wits.LeagueRank

func (standings PlayerStandingsJSON) Rank() wits.LeagueRank {
	return wits.LeagueRank(standings.Rank_)
}

type StandingsAfterJSON struct {
	Tier_  LeagueTierJSON `json:"tier"`
	Rank_  LeagueRankJSON `json:"rank"`
	Delta_ int            `json:"delta"`
}

func (standings StandingsAfterJSON) Tier() wits.LeagueTier {
	return wits.LeagueTier(standings.Tier_)
}

func (standings StandingsAfterJSON) Rank() wits.LeagueRank {
	return wits.LeagueRank(standings.Rank_)
}

func (standings StandingsAfterJSON) Delta() int {
	return standings.Delta_
}

// This value type has hard-coded limits of 0..5 checked when decoding.
type BaseHealth wits.BaseHealth

func (health *BaseHealth) UnmarshalJSON(encoded []byte) error {
	var hp int
	if err := json.Unmarshal(encoded, &hp); err != nil {
		return err
	}
	if hp < 0 || hp > 5 {
		return fmt.Errorf("invalid base HP: %d", hp)
	}

	*health = BaseHealth(wits.BaseHealth(hp))
	return nil
}

type TerminalStatusJSON wits.TerminalStatus
