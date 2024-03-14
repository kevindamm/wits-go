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
// github:kevindamm/wits-go/witsjson/units.go

package witsjson

import (
	"encoding/json"
	"fmt"

	"github.com/kevindamm/wits-go/schema"
)

// Initial unit description, satisfies both Unit and UnitInit interfaces.
//
// This format is suitable for storing in the map definition and replay init.
// Typically, however, the coordinate is part of the tile (not unit) position.
// This will also be missing the race association because map definitions are
// player-race agnostic -- the race isn't chosen until match start.
//
// {"i": 4, "j": 2, "team": 1, "class": 1}
//
// When marshaled as JSON it is always as the dict/object representation.
// The shorthand can be obtained via GdlEncoding.  It isn't enough context to
// make a Unit directly, or to implement Unit, because specials are determined
// by race and race hasn't been chosen yet.  But it is enough for GameMatch
// initialization to construct the equivalent via NewUnit and into a TileState,
// after a game match has begun.
type UnitInitJSON struct {
	Coord  schema.HexCoord  `json:"coord"`
	Team_  FriendlyEnumJSON `json:"team"`
	Class_ UnitClassJSON    `json:"class"`
}

func (init UnitInitJSON) Position() schema.HexCoord {
	return init.Coord
}

func (init UnitInitJSON) Class() schema.UnitClassEnum {
	return schema.UnitClassEnum(init.Class_)
}

func (init UnitInitJSON) IsSpecial() bool {
	return schema.UnitClassEnum(init.Class_) == schema.CLASS_SPECIAL
}

func (init UnitInitJSON) Race() schema.UnitRaceEnum {
	return schema.RACE_UNKNOWN
}

func (init UnitInitJSON) Team() schema.FriendlyEnum {
	return schema.FriendlyEnum(init.Team_)
}

func (init UnitInitJSON) Cost() schema.ActionPoints {
	// Retrieves the cost for this unit.
	return schema.CostForUnit(schema.UnitClassEnum(init.Class_))
}

func (init UnitInitJSON) Strength() schema.UnitHealth {
	// Retrieves the strength for this unit
	// (inaccuratelly generalized to 3 for specials)
	return schema.StrengthForUnit(schema.UnitClassEnum(init.Class_))
}

func (init UnitInitJSON) Distance() schema.TileDistance {
	// Retrieves the distance (visible ant motive) for this unit.
	return schema.DistanceForUnit(schema.UnitClassEnum(init.Class_))
}

func (unit *UnitInitJSON) UnmarshalJSON(encoded []byte) error {
	// Only parse unit init as JSON.
	var initial struct {
		Coord schema.HexCoord `json:"coord"`
		Team  string          `json:"team"`
		Class string          `json:"class"`
	}
	if err := json.Unmarshal(encoded, &initial); err != nil {
		return err
	}
	unit.Coord = initial.Coord
	unit.Team_ = FriendlyEnumJSON(ParseTeam(initial.Team))
	unit.Class_ = UnitClassJSON(ParseClass(initial.Class))
	return nil
}

// An s-expression compatible relation defining the unit, serialized as a string.
func (unit UnitInitJSON) GdlEncoding() string {
	return fmt.Sprintf(`["unit", ["class", "%s"], ["ij", %d, %d], ["team", "%s"]]`,
		string(unit.Class_), unit.Coord.I(), unit.Coord.J(), string(unit.Team_))
}

// UNIT RACE

type UnitRaceJSON schema.UnitRaceEnum

func ParseRace(name string) UnitRaceJSON {
	return map[string]UnitRaceJSON{
		"UNKNOWN":     UnitRaceJSON(schema.RACE_UNKNOWN),
		"FEEDBACK":    UnitRaceJSON(schema.RACE_FEEDBACK),
		"ADORABLES":   UnitRaceJSON(schema.RACE_ADORABLES),
		"SCALLYWAGS":  UnitRaceJSON(schema.RACE_SCALLYWAGS),
		"VEGGIENAUTS": UnitRaceJSON(schema.RACE_VEGGIENAUTS),
	}[name]
}

func (race *UnitRaceJSON) UnmarshalJSON(encoded []byte) error {
	var intVal int
	if err := json.Unmarshal(encoded, &intVal); err == nil {
		*race = UnitRaceJSON(intVal)
		return nil
	}
	var strVal string
	if err := json.Unmarshal(encoded, &strVal); err != nil {
		return err
	}
	*race = ParseRace(strVal)
	return nil
}

func (race UnitRaceJSON) MarshalJSON() ([]byte, error) {
	encoded := schema.UnitRaceEnum(race).String()
	if encoded == "UNKNOWN" {
		return []byte{}, fmt.Errorf("unknown race %d", byte(race))
	}
	return json.Marshal(encoded)
}

// UNIT CLASS

type UnitClassJSON schema.UnitClassEnum

func ParseClass(name string) schema.UnitClassEnum {
	return map[string]schema.UnitClassEnum{
		"UNKNOWN": schema.CLASS_UNKNOWN,
		"RUNNER":  schema.CLASS_RUNNER,
		"SOLDIER": schema.CLASS_SOLDIER,
		"MEDIC":   schema.CLASS_MEDIC,
		"SNIPER":  schema.CLASS_SNIPER,
		"HEAVY":   schema.CLASS_HEAVY,
		"THORN":   schema.CLASS_THORN,
		"SPECIAL": schema.CLASS_SPECIAL,
	}[name]
}

func (class *UnitClassJSON) UnmarshalJSON(encoded []byte) error {
	var strVal string
	if err := json.Unmarshal(encoded, &strVal); err == nil {
		*class = UnitClassJSON(ParseClass(strVal))
		if *class == UnitClassJSON(schema.CLASS_UNKNOWN) {
			return fmt.Errorf("unknown class in JSON [%s]", strVal)
		}
		return nil
	}
	var intVal uint
	if err := json.Unmarshal(encoded, &intVal); err != nil {
		return err
	}
	if intVal == uint(schema.CLASS_UNKNOWN) {
		return fmt.Errorf("unknown class in JSON [%d]", intVal)
	}
	if intVal > uint(schema.CLASS_SPECIAL) {
		return fmt.Errorf("invalid class value [%d]", intVal)
	}
	*class = UnitClassJSON(intVal)
	return nil
}

func (class UnitClassJSON) MarshalJSON() ([]byte, error) {
	encoded := schema.UnitClassEnum(class).String()
	if encoded == "UNKNOWN" {
		return []byte{}, fmt.Errorf("marshaling unknown class %d", byte(class))
	}
	return json.Marshal(encoded)
}
