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
	"strings"
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

type HexCoord struct {
	Column int `json:"positionI"`
	Row    int `json:"positionJ"`
}

func (coord HexCoord) MarshalJSON() ([]byte, error) {
	asArray := []int{coord.Column, coord.Row}
	return json.Marshal(asArray)
}

// Able to unmarshal from both int and bool representations,
// always marshals as boolean.
type Boolish struct {
	Value bool
}

func (b *Boolish) UnmarshalJSON(encoded []byte) error {
	var boolVal bool
	if err := json.Unmarshal(encoded, &boolVal); err != nil {
		b.Value = boolVal
	} else {
		var intVal int
		if err := json.Unmarshal(encoded, &boolVal); err != nil {
			b.Value = (intVal != 0)
		} else {
			return err
		}
	}
	return nil
}

func (b Boolish) MarshalJSON() ([]byte, error) {
	return json.Marshal(b.Value)
}

type UsedSpawn struct {
	SpawnX int `json:"ix"`
	SpawnY int `json:"iy"`
}

type OsnGameID string

// Automatically trims the prefix when encoding.
func (id OsnGameID) MarshalJSON() ([]byte, error) {
	// Trim the common prefix off when writing the game ID.
	return json.Marshal(id.ShortID())
}

// This 48~character string is the same across ALL game-replay identifiers.
const COMMON_PREFIX string = "ahRzfm91dHdpdHRlcnNnYW1lLWhyZHIVCxIIR2FtZVJvb20Y"

func (id OsnGameID) ShortID() string {
	return strings.TrimPrefix(string(id), COMMON_PREFIX)
}

func (id *OsnGameID) UnmarshalJSON(encoded []byte) error {
	return json.Unmarshal(encoded, id)
}

// in-progress, destruction, extinction, forfeit
type GameStatus int

// 0 = Player1, 1 = Player2, 2 = Player3, 3 = Player4
type PlayerIndex int

type ReplayData struct {
	Turns []OsnPlayerTurn
}

func (data *ReplayData) UnmarshalJSON(encoded []byte) error {
	var frames []Frame
	if err := json.Unmarshal(encoded, &frames); err != nil {
		fmt.Println(err)
		return err
	}

	data.Turns = make([]OsnPlayerTurn, 0)
	turnIndex := -1

	for _, frame := range frames {
		if frame.State != nil {
			newTurn := OsnPlayerTurn{Actions: make([]OsnPlayerAction, 0)}
			newTurn.State = *frame.State
			data.Turns = append(data.Turns, newTurn)
			turnIndex += 1
		} else if frame.Action != nil {
			if turnIndex == -1 {
				if (*frame.Action).Name() != "EndTurnAction" {
					return fmt.Errorf("before-state action of type %s", (*frame.Action).Name())
				}
				continue
			}
			data.Turns[turnIndex].Actions = append(
				data.Turns[turnIndex].Actions, *frame.Action)
		}
	}
	return nil
}

type Frame struct {
	Action *OsnPlayerAction
	State  *GameState
}

type partialFrame struct {
	Action struct {
		Name string `json:"name"`
	} `json:"action"`
	State *GameState `json:"gameState"`
}

func (frame *Frame) UnmarshalJSON(encoded []byte) error {
	var partial partialFrame
	err := json.Unmarshal(encoded, &partial)
	if err != nil {
		return err
	}

	if partial.Action.Name == "" {
		frame.State = partial.State
	} else if partial.State == nil {
		action, err := ParseGenericAction(partial.Action.Name, encoded)
		if err != nil {
			return err
		}
		frame.Action = &action
	}
	return nil
}

func (data *ReplayData) String() string {
	return fmt.Sprintf("Replay with %d turns; final state\n%v",
		len(data.Turns),
		data.Turns[len(data.Turns)-1].State)
}

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

type PlayerSettings struct {
	Name           string       `json:"name"`
	PlayerID       GCID         `json:"gcID"`
	Color          int          `json:"color"`
	Race           UnitRace     `json:"race"`
	Team           TeamIndex    `json:"team"`
	Index          PlayerIndex  `json:"id"`
	AP             int          `json:"actionPoints"`
	BasePreference BaseCosmetic `json:"basePref"`
	Invited        Boolish      `json:"isInvited"`
	Placeholder    Boolish      `json:"isPlaceHolder"`
}

type PlayerUpdate struct {
	Name  string      `json:"name"`
	GCID  GCID        `json:"gcID"`
	Color PlayerColor `json:"color"`
	Race  UnitRace    `json:"race"`
	Team  int         `json:"team"`
	Index PlayerIndex `json:"owner"`

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

func (replay GameReplay) MarshalJSON() ([]byte, error) {
	type PlayerFormat struct {
		Name string    `json:"name"`
		GCID GCID      `json:"gcID"`
		Race UnitRace  `json:"race"`
		Team TeamIndex `json:"team"`
	}
	type UnitInit struct {
		Position HexCoord  `json:"position"`
		Team     TeamIndex `json:"team"`
		Class    UnitClass `json:"class"`
	}
	type UnitState struct {
		Index    UnitIndex  `json:"identifier"`
		Position HexCoord   `json:"position"`
		Team     TeamIndex  `json:"team"`
		Race     UnitRace   `json:"race"`
		Class    UnitClass  `json:"class"`
		Health   UnitHealth `json:"health"`

		Attacked    Boolish        `json:"hasAttacked"`
		Moved       Boolish        `json:"hasMoved"`
		Transformed Boolish        `json:"hasTransformed"`
		AltForm     *AlternateForm `json:"alternate,omitempty"`
		Possession  *Possession    `json:"possessed,omitempty"`
	}

	type Before struct {
		GCID   GCID       `json:"gcID"`
		League LeagueTier `json:"tier"`
		Rank   LeagueRank `json:"rank"`
	}
	type After struct {
		GCID   GCID       `json:"gcID"`
		League LeagueTier `json:"tier"`
		Rank   LeagueRank `json:"rank"`
		Delta  int        `json:"delta"`
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
		p.Name = player.Name
		p.Race = player.Race
		p.Team = player.Team
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
