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
// github:kevindamm/wits-go/replay.go

package wits

type GameReplay interface {
	GameID() MatchID
	MapID() GameMapID
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

type BaseHealth byte

type PlayerTurn interface {
	TurnCount() uint
	Actions() []PlayerAction

	// Temporarily here so that we can validate the simulation against the intermediate states.
	State() GameState
}

type PlayerAction interface {
	ActionName() string
	RelVarEncoding() string
	Visit(*GameState) error
}

// Non-negative integer, the amount of "wits" (action points) available, or cost.
type ActionPoints byte

//
// UNKNOWN ACTION
//

type UnknownActionError struct {
	Name string
}

func (e UnknownActionError) Error() string { return "unknown action name " + e.Name }

//
// PASS (retain remaining wits)
//

// While not absolutely necessary to have the encoding work out, as the player
// turns have dedicated lists they are in, this "no-op" action is crucial.  It
// allows for simplifying assumptions later about empty sets/subsets vis-a-vis
// the recurrence relation of reordering action subgroups (when canonicalizing).
type PassAction struct{}

func (PassAction) ActionName() string     { return "Pass" }
func (PassAction) RelVarEncoding() string { return `["pass"]` }
func (PassAction) Visit(*GameState) error { return nil }
