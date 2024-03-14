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
// github:kevindamm/wits-go/schema/replay.go

package schema

type GameReplay interface {
	GameID() OsnGameID
	Map() GameMap
	MapTheme() string
	Players() []PlayerRole

	InitState() GameInit
	MatchReplay() []PlayerTurn
	MatchResult() TerminalStatus
}

type GameState interface {
	BaseHP(player FriendlyEnum) BaseHealth
	BonusWits() []HexCoord
	Units() []UnitPlacement
}

type GameInit interface {
	Units() []UnitInit
}

// Outcome enum relative to the player (player 1 if without context)
// Values include win (destruction), win (extinction), ...lose, forfeit
// These
type TerminalStatus byte

const (
	STATUS_UNKNOWN TerminalStatus = iota
	VICTORY_DESTRUCTION
	VICTORY_EXTINCTION
	VICTORY_RESIGNATION
	LOSS_DELAY_OF_GAME
	LOSS_DESTRUCTION
	LOSS_EXTINCTION
	LOSS_RESIGNATION
)

type BaseHealth byte