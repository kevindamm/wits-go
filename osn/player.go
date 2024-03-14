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
// github:kevindamm/wits-go/osn/player.go

package osn

import (
	"github.com/kevindamm/wits-go/schema"
	"github.com/kevindamm/wits-go/witsjson"
)

// 0 = Player1, 1 = Player2, 2 = Player3, 3 = Player4
type PlayerIndex int

const (
	PLAYER_1 PlayerIndex = iota
	PLAYER_2
	PLAYER_3
	PLAYER_4
)

// The usual properties found in "settings" property of a replay.
// Satisfies the PlayerID and PlayerRole interfaces.
type PlayerSettings struct {
	Name_          string       `json:"name"`
	PlayerID       GCID         `json:"gcID"`
	Color          int          `json:"color"`
	Race_          UnitRaceOsn  `json:"race"`
	Team_          TeamIndex    `json:"team"`
	Index          PlayerIndex  `json:"id"`
	AP             int          `json:"actionPoints"`
	BasePreference BaseCosmetic `json:"basePref"`
	Invited        Boolish      `json:"isInvited"`
	Placeholder    Boolish      `json:"isPlaceHolder"`
}

func (player PlayerSettings) GCID() schema.GCID {
	return schema.GCID(player.PlayerID)
}

func (player PlayerSettings) Name() schema.PlayerName {
	return schema.PlayerName(player.Name_)
}

func (player PlayerSettings) Race() schema.UnitRaceEnum {
	return schema.UnitRaceEnum(player.Race_)
}

func (player PlayerSettings) Team() schema.FriendlyEnum {
	if player.Index == PLAYER_1 {
		return schema.FR_SELF
	} else if player.Index == PLAYER_2 {
		return schema.FR_ENEMY
	} else if player.Index == PLAYER_3 {
		return schema.FR_ALLY
	} else if player.Index == PLAYER_4 {
		return schema.FR_ENEMY2
	} else {
		return schema.FR_UNKNOWN
	}
}

// Player standings before and after the replay's match.
// Cannot satisfy both the PlayerStandings and StandingsAfter without some
// ambiguity, so instead it satisfies MatchOutcome which can furnish these
// with its Before() and After() methods.  Also satisfies PlayerID and PlayerRole
type PlayerUpdate struct {
	Name  string                    `json:"name"`
	GCID  GCID                      `json:"gcID"`
	Color PlayerColor               `json:"color"`
	Race  witsjson.UnitRaceJSON     `json:"race"`
	Team  witsjson.FriendlyEnumJSON `json:"team"`
	Index PlayerIndex               `json:"owner"`

	BaseHealth int     `json:"baseHealth"`
	Demoted    Boolish `json:"wasDemoted"`
	Promoted   Boolish `json:"wasPrmomoted"`

	OldLeague     LeagueTier `json:"oldLeague"`
	OldLeagueRank LeagueRank `json:"oldLeagueRank"`
	NewLeague     LeagueTier `json:"newLeague"`
	NewLeagueRank LeagueRank `json:"newLeagueRank"`
	Direction     Direction  `json:"rankDirection"`
	PointsDelta   int        `json:"leaguePointsDelta"`
}

func (update PlayerUpdate) Before() schema.PlayerStandings {
	before := witsjson.PlayerStandingsJSON{
		Tier_: witsjson.LeagueTierJSON(update.OldLeague),
		Rank_: witsjson.LeagueRankJSON(update.OldLeagueRank)}
	return before
}

func (update PlayerUpdate) After() schema.StandingsAfter {
	after := witsjson.StandingsAfterJSON{
		Tier_:  witsjson.LeagueTierJSON(update.NewLeague),
		Rank_:  witsjson.LeagueRankJSON(update.NewLeagueRank),
		Delta_: update.PointsDelta}
	return after
}

type Direction int

type GCID string // in /G:(\d+)/ format

type BaseCosmetic int

type PlayerColor int

type TeamIndex int

const (
	COLOR_BLUE  PlayerColor = 1
	COLOR_RED   PlayerColor = 2
	COLOR_GREEN PlayerColor = 3
	COLOR_GOLD  PlayerColor = 4
)

func (color PlayerColor) String() string {
	return []string{
		"BLUE", "RED", "GREEN", "GOLD",
	}[int(color)]
}

type LeagueTier string
type LeagueRank int
