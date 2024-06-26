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
// github:kevindamm/wits-go/witsjson/replay.go

package witsjson

import (
	"encoding/json"
	"os"

	"github.com/kevindamm/wits-go"
)

// This is not a method offered in wits.GameReplay but perhaps it should be.
func (replay GameReplayJSON) WriteJSON(filename string) error {
	encoded, err := json.Marshal(replay)
	if err != nil {
		return err
	}
	return os.WriteFile(filename, encoded, 0644)
}

type GameReplayJSON struct {
	GameID_  OsnGameID        `json:"game_id"`
	GameMap_ wits.GameMapName `json:"map_name"`
	Init_    GameInitJSON     `json:"init,omitempty"`
	Turns_   []PlayerTurnJSON `json:"replay"`

	Players_ []PlayerRoleJSON `json:"players"`
}

func (replay GameReplayJSON) GameID() wits.MatchID {
	return wits.MatchID(replay.GameID_.ShortID())
}

func (replay GameReplayJSON) MapName() wits.GameMapName {
	return replay.GameMap_
}

func (replay GameReplayJSON) MatchReplay() []wits.PlayerTurn {
	turns := make([]wits.PlayerTurn, len(replay.Turns_))
	// Unfortunately need to make a shallow copy here
	// because golang does not have covariant return types.
	for i, turn := range replay.Turns_ {
		turns[i] = turn
	}
	return turns
}

type GameInitJSON struct {
	// Defaults for all these values are defined in the map (see GameMap)
	Units_      []wits.UnitInit `json:"units,omitempty"`
	UsedSpawns_ []wits.HexCoord `json:"used_spawns,omitempty"`
	BonusWits_  []wits.HexCoord `json:"bonus_wits,omitempty"`
	BaseHP_     []BaseHealth    `json:"base_hp,omitempty"` // all bases default 5hp
}

func (init GameInitJSON) Units() []wits.UnitInit {
	return init.Units_
}

func (init GameInitJSON) UsedSpawns() []wits.HexCoordIndex {
	// TODO map needed for conversion, serialized format maybe shouldn't depend on indices
	return []wits.HexCoordIndex{}
}

func (init GameInitJSON) BonusWits() []wits.HexCoordIndex {
	// TODO map needed for conversion, serialized format maybe shouldn't depend on indices
	return []wits.HexCoordIndex{}
}
