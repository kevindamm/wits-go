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
// github:kevindamm/wits-go/osn/units.go

package osn

import (
	"github.com/kevindamm/wits-go"
	"github.com/kevindamm/wits-go/witsjson"
)

type UnitStatus struct { // also see wits.UnitPlacement, wits.UnitTurnState
	Index UnitIndex    `json:"identifier"`
	Owner PlayerIndex  `json:"owner"`
	Team  TeamIndex    `json:"team"`
	Color PlayerColor  `json:"color"`
	Class UnitClassOsn `json:"class"`
	Race  UnitRaceOsn  `json:"race"`

	HexCoord
	Attacked    Boolish `json:"hasAttacked"`
	Moved       Boolish `json:"hasMoved"`
	Transformed Boolish `json:"hasTransformed"`

	Health    UnitHealth `json:"health"`
	AltStatus Boolish    `json:"isAlt"`
	HealthAlt UnitHealth `json:"altHealth"`

	Parent      UnitIndex `json:"parent"`
	SpawnedFrom UnitIndex `json:"spawnedFrom"`
}

type UnitIndex int

type AlternateForm struct {
	Enabled Boolish    `json:"enabled"`
	Health  UnitHealth `json:"health"`
}

type Possession struct {
	Parent UnitIndex `json:"parent"`
	From   UnitIndex `json:"from"`
}

type UnitHealth int

// The unit's class (indicating their strength, distance, health, etc.) is
// represented differently in OSN than it is in other representations, so
// it gets its own enum here instead of relying only on wits.UnitClassEnum.
type UnitClassOsn byte

const (
	CLASS_UNKNOWN UnitClassOsn = iota
	CLASS_RUNNER
	CLASS_SOLDIER
	CLASS_MEDIC
	CLASS_SNIPER
	CLASS_HEAVY
	CLASS_SCRAMBLER
	CLASS_MOBI
	CLASS_BOMBSHELL
	CLASS_BRAMBLE
	CLASS_BRAMBLE_THORN
)

func (class UnitClassOsn) String() string {
	if class > CLASS_BRAMBLE_THORN {
		class = CLASS_UNKNOWN
	}
	return []string{
		"UNKNOWN",
		"RUNNER",
		"SOLDIER",
		"MEDIC",
		"SNIPER",
		"HEAVY",
		"SCRAMBLER",
		"MOBI",
		"BOMBSHELL",
		"BRAMBLE",
		"BRAMBLE_THORN",
	}[int(class)]
}

func (class UnitClassOsn) AsWitsJSON() witsjson.UnitClassJSON {
	if class > CLASS_BRAMBLE_THORN {
		class = CLASS_UNKNOWN
	}
	return []witsjson.UnitClassJSON{
		witsjson.UnitClassJSON(wits.CLASS_UNKNOWN),
		witsjson.UnitClassJSON(wits.CLASS_RUNNER),
		witsjson.UnitClassJSON(wits.CLASS_SOLDIER),
		witsjson.UnitClassJSON(wits.CLASS_MEDIC),
		witsjson.UnitClassJSON(wits.CLASS_SNIPER),
		witsjson.UnitClassJSON(wits.CLASS_HEAVY),
		witsjson.UnitClassJSON(wits.CLASS_SPECIAL),
		witsjson.UnitClassJSON(wits.CLASS_SPECIAL),
		witsjson.UnitClassJSON(wits.CLASS_SPECIAL),
		witsjson.UnitClassJSON(wits.CLASS_SPECIAL),
		witsjson.UnitClassJSON(wits.CLASS_THORN),
	}[int(class)]
}

// OSN enumeration of the different squads a player can choose from.
type UnitRaceOsn byte

const (
	RACE_UNKNOWN UnitRaceOsn = iota
	RACE_FEEDBACK
	RACE_ADORABLES
	RACE_SCALLYWAGS
	RACE_VEGGIENAUTS
)

func (race UnitRaceOsn) String() string {
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

func (race UnitRaceOsn) AsWitsJSON() witsjson.UnitRaceJSON {
	if race > RACE_VEGGIENAUTS {
		race = RACE_UNKNOWN
	}
	return []witsjson.UnitRaceJSON{
		witsjson.UnitRaceJSON(wits.RACE_UNKNOWN),
		witsjson.UnitRaceJSON(wits.RACE_FEEDBACK),
		witsjson.UnitRaceJSON(wits.RACE_ADORABLES),
		witsjson.UnitRaceJSON(wits.RACE_SCALLYWAGS),
		witsjson.UnitRaceJSON(wits.RACE_VEGGIENAUTS),
	}[race]
}
