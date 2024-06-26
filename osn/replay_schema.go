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
// github:kevindamm/wits-go/osn/replay_schema.go

package osn

import (
	"encoding/json"
	"fmt"

	"github.com/kevindamm/wits-go"
)

type GameReplay struct {
	OsnGameID string           `json:"game_id"`
	MapName   string           `json:"map_name"`
	MapTheme  string           `json:"map_theme"`
	Settings  []PlayerSettings `json:"settings"`
	GameState

	GameOverData GameOverData `json:"gameOverData"`
	Replay       ReplayData   `json:"replay"`
}

type GameState struct {
	TurnCount  int          `json:"turnCount,omitempty"`
	Units      []UnitStatus `json:"units"`
	UsedSpawns []UsedSpawn  `json:"used_spawns,omitempty"`

	CapturedStates []CaptureTileState `json:"captureTileStates"`
	CurrentPawnID  int                `json:"currentPawnID"`
	CurrentPlayer  PlayerIndex        `json:"currentPlayer"`
	Base0_HP       BaseHealth         `json:"hp_base0"`
	Base1_HP       BaseHealth         `json:"hp_base1"`
	Outcome        GameStatus         `json:"outcome,omitempty"`
}

type UsedSpawn struct {
	SpawnX int `json:"ix"`
	SpawnY int `json:"iy"`
}

// in-progress, destruction, extinction, forfeit
type GameStatus int
type CaptureTileState struct {
	Column int `json:"tileI"`
	Row    int `json:"tileJ"`
	Type   int `json:"tileType"`
}

type GameOverData struct {
	LeagueMatch Boolish        `json:"isLeagueMatch"`
	Online      Boolish        `json:"isOnline"`
	Winners     []PlayerUpdate `json:"winners"`
	Losers      []PlayerUpdate `json:"losers"`
}

type BaseHealth int

func (replay GameReplay) MarshalJSON() ([]byte, error) {
	type PlayerFormat struct {
		Name string      `json:"name"`
		GCID GCID        `json:"gcID"`
		Race UnitRaceOsn `json:"race"`
		Team TeamIndex   `json:"team"`
	}
	type UnitInit struct {
		Position HexCoord     `json:"position"`
		Team     TeamIndex    `json:"team"`
		Class    UnitClassOsn `json:"class"`
	}
	type UnitState struct {
		Index    UnitIndex    `json:"identifier"`
		Position HexCoord     `json:"position"`
		Team     TeamIndex    `json:"team"`
		Race     UnitRaceOsn  `json:"race"`
		Class    UnitClassOsn `json:"class"`
		Health   UnitHealth   `json:"health"`

		Attacked    Boolish        `json:"hasAttacked"`
		Moved       Boolish        `json:"hasMoved"`
		Transformed Boolish        `json:"hasTransformed"`
		AltForm     *AlternateForm `json:"alternate,omitempty"`
		Possession  *Possession    `json:"possessed,omitempty"`
	}

	type Before struct {
		GCID   GCID            `json:"gcID"`
		League wits.LeagueTier `json:"tier"`
		Rank   wits.LeagueRank `json:"rank"`
	}
	type After struct {
		GCID   GCID            `json:"gcID"`
		League wits.LeagueTier `json:"tier"`
		Rank   wits.LeagueRank `json:"rank"`
		Delta  int             `json:"delta"`
	}
	type CapturedFormat struct {
		Position HexCoord `json:"position"`
		Tile     int      `json:"tile"`
	}
	type CheckpointFormat struct {
		Base0_HP   BaseHealth       `json:"base0_hp"`
		Base1_HP   BaseHealth       `json:"base1_hp"`
		UsedSpawns []HexCoord       `json:"used_spawns"`
		Captured   []CapturedFormat `json:"captured"`
		Units      []UnitState      `json:"units"`
	}
	type OutcomeFormat struct {
		Result GameStatus `json:"result"`
		Before []Before   `json:"before"`
		After  []After    `json:"after"`
		CheckpointFormat
	}
	type TurnFormat struct {
		Turn       int                      `json:"turn"`
		Actions    []map[string]interface{} `json:"actions"`
		Checkpoint CheckpointFormat         `json:"checkpoint"`
	}
	type OutputFormat struct {
		OsnGameID string         `json:"game_id"`
		MapName   string         `json:"map_name"`
		MapTheme  string         `json:"map_theme"`
		Players   []PlayerFormat `json:"players"`
		Units     []UnitInit     `json:"units"`
		Outcome   OutcomeFormat  `json:"outcome"`
		Replay    []TurnFormat   `json:"replay"`
	}

	var output OutputFormat
	output.OsnGameID = replay.OsnGameID
	output.MapName = replay.MapName
	output.MapTheme = replay.MapTheme
	output.Players = make([]PlayerFormat, 0)
	for _, player := range replay.Settings {
		var p PlayerFormat
		p.GCID = player.PlayerID
		p.Name = player.Name_
		p.Race = player.Race_
		p.Team = player.Team_
		output.Players = append(output.Players, p)
	}
	if len(replay.Replay.Turns) > 0 {
		output.Units = make([]UnitInit, 0)
		for _, unit := range replay.Replay.Turns[0].State.Units {
			var u UnitInit
			u.Position = unit.HexCoord
			u.Team = unit.Team
			u.Class = unit.Class
			output.Units = append(output.Units, u)
		}
	} else {
		return nil, fmt.Errorf("short replay turns? no initial units. skipping write")
	}
	output.Outcome.Result = replay.Outcome
	output.Outcome.Before = make([]Before, 2)
	output.Outcome.After = make([]After, 2)
	for _, winner := range replay.GameOverData.Winners {
		// add entry for before
		output.Outcome.Before[0] = Before{
			winner.GCID, winner.OldLeague, winner.OldLeagueRank}
		// and an entry for after, with positive delta
		output.Outcome.After[0] = After{
			winner.GCID, winner.NewLeague, winner.NewLeagueRank, winner.PointsDelta}
	}
	for _, loser := range replay.GameOverData.Losers {
		// add entry for before
		output.Outcome.Before[1] = Before{
			loser.GCID, loser.OldLeague, loser.OldLeagueRank}
		// and an entry for after, with negative delta
		output.Outcome.After[1] = After{
			loser.GCID, loser.NewLeague, loser.NewLeagueRank, loser.PointsDelta}
	}
	output.Outcome.Base0_HP = replay.Base0_HP
	output.Outcome.Base1_HP = replay.Base1_HP
	output.Outcome.UsedSpawns = make([]HexCoord, len(replay.UsedSpawns))
	for i, spawn := range replay.UsedSpawns {
		output.Outcome.UsedSpawns[i] = HexCoord{spawn.SpawnX, spawn.SpawnY}
	}
	output.Outcome.Captured = make([]CapturedFormat, len(replay.CapturedStates))
	for i, captured := range replay.CapturedStates {
		output.Outcome.Captured[i].Position = HexCoord{captured.Column, captured.Row}
		output.Outcome.Captured[i].Tile = captured.Type
	}
	output.Outcome.Units = make([]UnitState, len(replay.Units))
	for i, unit := range replay.Units {
		output.Outcome.Units[i].Index = unit.Index
		output.Outcome.Units[i].Position = unit.HexCoord
		output.Outcome.Units[i].Team = unit.Team
		output.Outcome.Units[i].Class = unit.Class
		output.Outcome.Units[i].Health = unit.Health
		output.Outcome.Units[i].Attacked = unit.Attacked
		output.Outcome.Units[i].Moved = unit.Moved
		output.Outcome.Units[i].Transformed = unit.Transformed
		if unit.AltStatus.Value || unit.HealthAlt > 0 {
			output.Outcome.Units[i].AltForm = &AlternateForm{unit.AltStatus, unit.HealthAlt}
		}
		if unit.Parent != -1 || unit.SpawnedFrom != -1 {
			output.Outcome.Units[i].Possession = &Possession{unit.Parent, unit.SpawnedFrom}
		}
	}
	output.Replay = make([]TurnFormat, len(replay.Replay.Turns))
	for i, turn := range replay.Replay.Turns {
		output.Replay[i].Turn = turn.State.TurnCount
		reducedActions := ActionReduction(replay.Replay.Turns[i].Actions)
		// Skip adding an empty action list if it's the last turn.
		if len(reducedActions) == 0 && i == (len(replay.Replay.Turns)-1) {
			output.Replay = output.Replay[:i]
			continue
		}
		output.Replay[i].Actions = reducedActions
		output.Replay[i].Checkpoint = CheckpointFormat{
			turn.State.Base0_HP,
			turn.State.Base1_HP,
			make([]HexCoord, len(turn.State.UsedSpawns)),
			make([]CapturedFormat, len(turn.State.CapturedStates)),
			make([]UnitState, len(turn.State.Units))}
		for j, coord := range turn.State.UsedSpawns {
			output.Replay[i].Checkpoint.UsedSpawns[j] = HexCoord{coord.SpawnX, coord.SpawnY}
		}
		for j, captured := range turn.State.CapturedStates {
			output.Replay[i].Checkpoint.Captured[j] = CapturedFormat{
				HexCoord{captured.Column, captured.Row}, captured.Type}
		}
		for j, unit := range turn.State.Units {
			output.Replay[i].Checkpoint.Units[j] = UnitState{
				unit.Index, HexCoord{unit.Column, unit.Row}, unit.Team, unit.Race, unit.Class,
				unit.Health, unit.Attacked, unit.Moved, unit.Transformed, nil, nil}
			if unit.AltStatus.Value || unit.HealthAlt > 0 {
				output.Replay[i].Checkpoint.Units[j].AltForm = &AlternateForm{unit.AltStatus, unit.HealthAlt}
			}
			if unit.Parent != -1 || unit.SpawnedFrom != -1 {
				output.Replay[i].Checkpoint.Units[j].Possession = &Possession{unit.Parent, unit.SpawnedFrom}
			}
		}
	}

	bytes, err := json.Marshal(output)
	if err != nil {
		fmt.Println(err)
	}
	return bytes, err
}

func ActionReduction(actions []OsnPlayerAction) []map[string]interface{} {
	reduced := make([]map[string]interface{}, 0)
	for i, action := range actions {
		var dict map[string]interface{}
		name := action.Name()
		if name == "StartTurnAction" ||
			name == "EndTurnAction" ||
			name == "SelectUnitAction" ||
			name == "SelectSpawnTileAction" ||
			name == "EatAction" {
			// These actions are not useful by themself but
			// sometimes adjacent actions will make use of their properties.
			// We skip them when encountering them directly.
			continue
		}

		if action.Name() == "SpawnUnitAction" {
			dict = action.AsDict()
			// Always preceded by a SelectSpawnTileAction, they can be combined.
			selectSpawn := actions[i-1].AsDict()
			dict["spawn"] = selectSpawn["position"]
		} else if action.Name() == "SpitAction" {
			dict = action.AsDict()
			eatAction := actions[i-1].AsDict()
			dict["name"] = "TeleportAction"
			dict["from"] = eatAction["from"]
		} else if action.Name() == "ToggleArtilleryModeAction" || action.Name() == "RootBrambleAction" {
			dict = action.AsDict()
			dict["name"] = "ToggleAction"
		} else {
			dict = action.AsDict()
		}

		reduced = append(reduced, dict)
	}
	return reduced
}
