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

import "encoding/json"

type UnitStatus struct { // also see schema.UnitPlacement, schema.UnitTurnState
	Index UnitIndex   `json:"identifier"`
	Owner PlayerIndex `json:"owner"`
	Team  TeamIndex   `json:"team"`
	Color PlayerColor `json:"color"`
	Class UnitClass   `json:"class"`
	Race  UnitRace    `json:"race"`

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

type AlternateForm struct {
	Enabled Boolish    `json:"enabled"`
	Health  UnitHealth `json:"health"`
}

type Possession struct {
	Parent UnitIndex `json:"parent"`
	From   UnitIndex `json:"from"`
}

type UnitClass int

const (
	CLASS_UNKNOWN       UnitClass = 0
	CLASS_SCOUT         UnitClass = 1
	CLASS_SOLDIER       UnitClass = 2
	CLASS_MEDIC         UnitClass = 3
	CLASS_SNIPER        UnitClass = 4
	CLASS_HEAVY         UnitClass = 5
	CLASS_SCRAMBLER     UnitClass = 6
	CLASS_MOBI          UnitClass = 7
	CLASS_ARTILLERY     UnitClass = 8
	CLASS_BRAMBLE       UnitClass = 9
	CLASS_BRAMBLE_THORN UnitClass = 10
)

func (class UnitClass) String() string {
	if class > CLASS_BRAMBLE_THORN || class < 0 {
		class = CLASS_UNKNOWN
	}
	return []string{
		"UNKNOWN_CLASS",
		"SCOUT",
		"SOLDIER",
		"MEDIC",
		"SNIPER",
		"HEAVY",
		"SCRAMBLER",
		"MOBI",
		"ARTILLERY",
		"BRAMBLE",
		"BRAMBLE_THORN",
	}[int(class)]
}

type UnitIndex int

type UnitRace int

const (
	RACE_UNKNOWN     UnitRace = 0
	RACE_FEEDBACK    UnitRace = 1
	RACE_ADORABLES   UnitRace = 2
	RACE_SCALLYWAGS  UnitRace = 3
	RACE_VEGGIENAUTS UnitRace = 4
)

func (race UnitRace) String() string {
	if race > RACE_VEGGIENAUTS || race < 0 {
		race = RACE_UNKNOWN
	}
	return []string{
		"UNKNOWN_RACE",
		"FEEDBACK",
		"ADORABLES",
		"SCALLYWAGS",
		"VEGGIENAUTS",
	}[int(race)]
}

func ParseRace(name string) UnitRace {
	return map[string]UnitRace{
		"UNKNOWN_RACE": RACE_UNKNOWN,
		"FEEDBACK":     RACE_FEEDBACK,
		"ADORABLES":    RACE_ADORABLES,
		"SCALLYWAGS":   RACE_SCALLYWAGS,
		"VEGGIENAUTS":  RACE_VEGGIENAUTS,
	}[name]
}

func (race *UnitRace) UnmarshalJSON(encoded []byte) error {
	var intVal int
	if err := json.Unmarshal(encoded, &intVal); err != nil {
		*race = UnitRace(intVal)
	} else {
		var strVal string
		if err := json.Unmarshal(encoded, &strVal); err != nil {
			*race = ParseRace(strVal)
		} else {
			return err
		}
	}
	return nil
}

func (race UnitRace) MarshalJSON() ([]byte, error) {
	return json.Marshal(int(race))
}

type UnitHealth int
