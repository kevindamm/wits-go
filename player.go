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
// github:kevindamm/wits-go/player.go

package wits

// Player information as it pertains to a single match.
type PlayerRole interface {
	PlayerID
	Name() PlayerName

	Race() UnitRaceEnum
	Team() FriendlyEnum

	Result() TerminalStatus
	BaseHP() BaseHealth
	Wits() ActionPoints
}

// For identifying players globally.
type PlayerID interface {
	PlayerKey() GCID
}

// A globally consistent identifier for players (from OML, via OSN)
type GCID string

// The human-readable representation (may contain unicode)
type PlayerName string

// Team alignment.
type FriendlyEnum byte

const (
	FR_UNKNOWN FriendlyEnum = iota
	// two roles relative to the current player.
	FR_SELF
	FR_ENEMY
	// additional roles available in `duos` game replays:
	FR_ALLY
	FR_ENEMY2
)

// Often-useful toggle for player role.
func (role FriendlyEnum) Opponent() FriendlyEnum {
	if role == FR_UNKNOWN {
		return FR_UNKNOWN
	}
	if role == FR_ENEMY || role == FR_ENEMY2 {
		return FR_SELF
	} else {
		return FR_ENEMY
	}
}
