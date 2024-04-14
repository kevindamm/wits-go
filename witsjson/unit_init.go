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

	"github.com/kevindamm/wits-go"
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
	Coord  HexCoordJSON     `json:"coord"`
	Team_  FriendlyEnumJSON `json:"team"`
	Class_ UnitClassJSON    `json:"class"`
}

func (init UnitInitJSON) Position() wits.HexCoord {
	return init.Coord
}

func (init UnitInitJSON) Class() wits.UnitClassEnum {
	return wits.UnitClassEnum(init.Class_)
}

func (init UnitInitJSON) IsSpecial() bool {
	return wits.UnitClassEnum(init.Class_) == wits.CLASS_SPECIAL
}

func (init UnitInitJSON) Race() wits.UnitRaceEnum {
	return wits.RACE_UNKNOWN
}

func (init UnitInitJSON) Team() wits.FriendlyEnum {
	return wits.FriendlyEnum(init.Team_)
}

func (init UnitInitJSON) Cost() wits.ActionPoints {
	// Retrieves the cost for this unit.
	return wits.CostForUnit(wits.UnitClassEnum(init.Class_))
}

func (init UnitInitJSON) Health() wits.UnitHealth {
	return 0 // receiver will use unit's default health.
}

func (init UnitInitJSON) Strength() wits.UnitHealth {
	// Retrieves the strength for this unit
	// (inaccuratelly generalized to 3 for specials)
	return wits.StrengthForUnit(wits.UnitClassEnum(init.Class_))
}

func (init UnitInitJSON) Distance() wits.TileDistance {
	// Retrieves the distance (visible ant motive) for this unit.
	return wits.DistanceForUnit(wits.UnitClassEnum(init.Class_))
}

func (unit *UnitInitJSON) UnmarshalJSON(encoded []byte) error {
	// Only parse unit init as JSON.
	var initial struct {
		Coord HexCoordJSON `json:"coord"`
		Team  string       `json:"team"`
		Class string       `json:"class"`
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

type UnitRaceJSON wits.UnitRaceEnum

func ParseRace(name string) UnitRaceJSON {
	return map[string]UnitRaceJSON{
		"UNKNOWN":     UnitRaceJSON(wits.RACE_UNKNOWN),
		"FEEDBACK":    UnitRaceJSON(wits.RACE_FEEDBACK),
		"ADORABLES":   UnitRaceJSON(wits.RACE_ADORABLES),
		"SCALLYWAGS":  UnitRaceJSON(wits.RACE_SCALLYWAGS),
		"VEGGIENAUTS": UnitRaceJSON(wits.RACE_VEGGIENAUTS),
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
	encoded := wits.UnitRaceEnum(race).String()
	if encoded == "UNKNOWN" {
		return []byte{}, fmt.Errorf("unknown race %d", byte(race))
	}
	return json.Marshal(encoded)
}

// UNIT CLASS

type UnitClassJSON wits.UnitClassEnum

func ParseClass(name string) wits.UnitClassEnum {
	return map[string]wits.UnitClassEnum{
		"UNKNOWN": wits.CLASS_UNKNOWN,
		"RUNNER":  wits.CLASS_RUNNER,
		"SOLDIER": wits.CLASS_SOLDIER,
		"MEDIC":   wits.CLASS_MEDIC,
		"SNIPER":  wits.CLASS_SNIPER,
		"HEAVY":   wits.CLASS_HEAVY,
		"THORN":   wits.CLASS_THORN,
		"SPECIAL": wits.CLASS_SPECIAL,
	}[name]
}

func (class *UnitClassJSON) UnmarshalJSON(encoded []byte) error {
	var strVal string
	if err := json.Unmarshal(encoded, &strVal); err == nil {
		*class = UnitClassJSON(ParseClass(strVal))
		if *class == UnitClassJSON(wits.CLASS_UNKNOWN) {
			return fmt.Errorf("unknown class in JSON [%s]", strVal)
		}
		return nil
	}
	var intVal uint
	if err := json.Unmarshal(encoded, &intVal); err != nil {
		return err
	}
	if intVal == uint(wits.CLASS_UNKNOWN) {
		return fmt.Errorf("unknown class in JSON [%d]", intVal)
	}
	if intVal > uint(wits.CLASS_SPECIAL) {
		return fmt.Errorf("invalid class value [%d]", intVal)
	}
	*class = UnitClassJSON(intVal)
	return nil
}

func (class UnitClassJSON) MarshalJSON() ([]byte, error) {
	encoded := wits.UnitClassEnum(class).String()
	if encoded == "UNKNOWN" {
		return []byte{}, fmt.Errorf("marshaling unknown class %d", byte(class))
	}
	return json.Marshal(encoded)
}
