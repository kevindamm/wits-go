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
// github:kevindamm/wits-go/units.go

package wits

// Intrinsic (and invariant) properties of the Unit
type Unit interface {
	Class() UnitClassEnum
	IsSpecial() bool
	Race() UnitRaceEnum
	Team() FriendlyEnum

	Cost() ActionPoints
	Strength() UnitHealth
	Distance() TileDistance
}

// Unit health (and strength) does not exceed 5, its max is determined by the
// unit's constructor; some units have a maix health of 2, all units have a max
// (and constant) UnitHealth equivalent measuring their strength.
type UnitHealth int

// Common interface for JSON-loaded and binary-deflated representations.  This
// is turn-intransient state (excluding status bits like HasMoved, HasActed).
// This state and a position are enough to make reasonable strategy decisions.
// Some information, such as parentage, would need to be derived and is subject
// to the partial information visibility of each player (the thorn that spawned
// a thorn may be known only if the parent was visible during spawning).  These
// properties were chosen because they are minimally-satisfying and as much as
// I could pack within just eight bits.
type UnitState interface {
	Unit
	Health() UnitHealth

	// Two special units (bombshell and bramble) have a "toggle" special action
	// that they need to perform before being able to do a more powerful action.
	IsAlternate() bool
	Toggle() UnitState

	// Health and Team are not directly assigned, they are affected by effects.
	ReceiveBoost() UnitState
	ReceiveDamage(other Unit) UnitState
	ReceiveCharm(other Unit) UnitState
	DoAction(action PlayerAction) (UnitState, error)
}

type UnitStateExtended interface {
	UnitState
	UnitTurnStatus
	UnitParentage
}

type UnitParentage interface {
	HasParent() bool
	Parent() HexCoordIndex
}

type UnitTurnStatus interface {
	HasActed() bool
	HasMoved() bool
	HasAlted() bool
}

type UnitPlacement interface {
	UnitState
	Placement
}

type UnitInit interface {
	Positional
	Team() FriendlyEnum
	Class() UnitClassEnum
	Health() UnitHealth // if 0 value, use default
}

// The type of unit (determining its movement, health, actions, ...) can actually fit
// in three bits (including an UNKNOWN enum) and is usually part of the Unit data.
// In the case of specials, the tribe data also needs to be known.
type UnitClassEnum byte

// These values are inherited from OML enumeration.
const (
	CLASS_UNKNOWN UnitClassEnum = iota
	CLASS_RUNNER
	CLASS_SOLDIER
	CLASS_MEDIC
	CLASS_SNIPER
	CLASS_HEAVY
	CLASS_THORN
	CLASS_SPECIAL
)

func (class UnitClassEnum) String() string {
	if class > CLASS_SPECIAL {
		class = CLASS_UNKNOWN
	}
	return []string{
		"UNKNOWN",
		"RUNNER",
		"SOLDIER",
		"MEDIC",
		"SNIPER",
		"HEAVY",
		"THORN",
		"SPECIAL",
	}[int(class)]
}

// UnitRace is an enumeration with both string and integer representations.
// In contexts where the explicit UNKNOWN value is unnecessary (e.g., in Unit)
// the other values may be squeezed into two bits without loss of generality.
type UnitRaceEnum byte

const (
	RACE_UNKNOWN UnitRaceEnum = iota
	RACE_FEEDBACK
	RACE_ADORABLES
	RACE_SCALLYWAGS
	RACE_VEGGIENAUTS
)

func (race UnitRaceEnum) String() string {
	if race > RACE_VEGGIENAUTS {
		race = RACE_UNKNOWN
	}
	return []string{
		"UNKNOWN",
		"FEEDBACK",
		"ADORABLES",
		"SCALLYWAGS",
		"VEGGIENAUTS",
	}[int(race)]
}

// Each unit class costs a consistent amount regardless of race or team.
func CostForUnit(class UnitClassEnum) ActionPoints {
	return cost[class]
}

var cost []ActionPoints

func init() {
	cost = []ActionPoints{
		0, // CLASS_UNKNOWN
		1, // CLASS_RUNNER
		2, // CLASS_SOLDIER
		2, // CLASS_MEDIC
		3, // CLASS_SNIPER
		4, // CLASS_HEAVY
		1, // CLASS_THORN
		7, // CLASS_SPECIAL
	}
}

// Calculating unit strength is a little tricky for special units because
// most specials do not have an attack, except the bombshell which has an
// added AoE splash-damage to adjacent tiles.  We generalize with 3 here.
func StrengthForUnit(class UnitClassEnum) UnitHealth {
	return strength[class]
}

var strength []UnitHealth

func init() {
	strength = []UnitHealth{
		0, // CLASS_UNKNOWN
		1, // CLASS_RUNNER
		2, // CLASS_SOLDIER
		0, // CLASS_MEDIC
		3, // CLASS_SNIPER
		3, // CLASS_HEAVY
		1, // CLASS_THORN
		3, // CLASS_SPECIAL
	}
}

// Distance is also straightforward for standard units but the special units
// vary depending on whether they are toggled into their alt state.  At the
// point where these details matter, the actions will be invoked on a unit
// status object which can distinguish which kind of special is the active unit.
func DistanceForUnit(class UnitClassEnum) TileDistance {
	return travel[class]
}

var travel []TileDistance

func init() {
	travel = []TileDistance{
		0, // CLASS_UNKNOWN
		5, // CLASS_RUNNER
		3, // CLASS_SOLDIER
		3, // CLASS_MEDIC
		1, // CLASS_SNIPER
		2, // CLASS_HEAVY
		2, // CLASS_THORN
		3, // CLASS_SPECIAL
	}
}
