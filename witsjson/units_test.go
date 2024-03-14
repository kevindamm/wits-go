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
// github:kevindamm/wits-go/witsjson/units_test.go

package witsjson_test

import (
	"encoding/json"
	"reflect"
	"strconv"
	"testing"

	"github.com/kevindamm/wits-go/schema"
	"github.com/kevindamm/wits-go/witsjson"
)

func TestUnitInit_JSON(t *testing.T) {
	tests := []struct {
		name    string
		encoded string
		succeed bool
		expect  witsjson.UnitInitJSON
	}{
		{"basic JSON", `{"class": "HEAVY", "coord": {"i": 1, "j": 2}, "team": "RED"}`,
			true, witsjson.UnitInitJSON{schema.NewHexCoord(1, 2),
				witsjson.FriendlyEnumJSON(schema.FR_SELF),
				witsjson.UnitClassJSON(schema.CLASS_HEAVY)}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var data witsjson.UnitInitJSON
			err := json.Unmarshal([]byte(tt.encoded), &data)
			if (err != nil) == tt.succeed {
				t.Errorf("UnitInit.UnarshalJSON() error = %v, expected? %v.", err, !tt.succeed)
			}
			if !reflect.DeepEqual(data, tt.expect) {
				t.Errorf("unmarshaled = %v, want %v", data, tt.expect)
			}
		})
	}
}

func TestUnitRaceJSON(t *testing.T) {
	tests := []struct {
		name    string
		encoded string
		succeed bool
		expect  witsjson.UnitRaceJSON
	}{
		{"basic1", `"FEEDBACK"`, true, witsjson.UnitRaceJSON(schema.RACE_FEEDBACK)},
		{"basic2", `"ADORABLES"`, true, witsjson.UnitRaceJSON(schema.RACE_ADORABLES)},
		{"basic3", `"SCALLYWAGS"`, true, witsjson.UnitRaceJSON(schema.RACE_SCALLYWAGS)},
		{"basic4", `"VEGGIENAUTS"`, true, witsjson.UnitRaceJSON(schema.RACE_VEGGIENAUTS)},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var race witsjson.UnitRaceJSON
			err := json.Unmarshal([]byte(tt.encoded), &race)
			if (err != nil) == tt.succeed {
				t.Errorf("UnitInit.UnarshalJSON() error = %v, expected? %v.", err, !tt.succeed)
			}
			if !reflect.DeepEqual(race, tt.expect) {
				t.Errorf("unmarshaled = %v, want %v", race, tt.expect)
			}

			if encoded, err := json.Marshal(race); err != nil {
				t.Error(err)
			} else {
				if tt.encoded != string(encoded) {
					t.Errorf("marshaled = %v, want %v", tt.encoded, string(encoded))
				}
			}
		})
	}
}

func TestUnitClassJSON(t *testing.T) {
	tests := []struct {
		name    string
		encoded string
		succeed bool
		expect  witsjson.UnitClassJSON
	}{
		{"standard", `"RUNNER"`, true, witsjson.UnitClassJSON(schema.CLASS_RUNNER)},
		{"standard", `"MEDIC"`, true, witsjson.UnitClassJSON(schema.CLASS_MEDIC)},
		{"standard", `"SOLDIER"`, true, witsjson.UnitClassJSON(schema.CLASS_SOLDIER)},
		{"standard", `"SNIPER"`, true, witsjson.UnitClassJSON(schema.CLASS_SNIPER)},
		{"standard", `"HEAVY"`, true, witsjson.UnitClassJSON(schema.CLASS_HEAVY)},
		{"special", `"SPECIAL"`, true, witsjson.UnitClassJSON(schema.CLASS_SPECIAL)},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var class witsjson.UnitClassJSON
			err := json.Unmarshal([]byte(tt.encoded), &class)
			if (err != nil) == tt.succeed {
				t.Errorf("UnitInit.UnarshalJSON() error = %v, expected? %v.", err, !tt.succeed)
			}
			if !reflect.DeepEqual(class, tt.expect) {
				t.Errorf("unmarshaled = %v, want %v", class, tt.expect)
			}

			if encoded, err := json.Marshal(class); (err != nil) == tt.succeed {
				t.Error(err)
			} else {
				if tt.encoded != string(encoded) {
					t.Errorf("marshaled = %v, want %v", tt.encoded, string(encoded))
				}
			}
		})
	}
}

func TestUnitClassJSON_UnmarshalExtra(t *testing.T) {
	tests := []struct {
		name    string
		encoded string
		succeed bool
		expect  witsjson.UnitClassJSON
	}{
		{"unknwon", `"UNKNOWN"`, false, witsjson.UnitClassJSON(schema.CLASS_UNKNOWN)},
		{"iunk", "0", false, witsjson.UnitClassJSON(schema.CLASS_UNKNOWN)},
		{"istandard", strconv.Itoa(int(schema.CLASS_RUNNER)), true,
			witsjson.UnitClassJSON(schema.CLASS_RUNNER)},
		{"istandard", strconv.Itoa(int(schema.CLASS_MEDIC)), true,
			witsjson.UnitClassJSON(schema.CLASS_MEDIC)},
		{"istandard", strconv.Itoa(int(schema.CLASS_SOLDIER)), true,
			witsjson.UnitClassJSON(schema.CLASS_SOLDIER)},
		{"istandard", strconv.Itoa(int(schema.CLASS_SNIPER)), true,
			witsjson.UnitClassJSON(schema.CLASS_SNIPER)},
		{"istandard", strconv.Itoa(int(schema.CLASS_HEAVY)), true,
			witsjson.UnitClassJSON(schema.CLASS_HEAVY)},
		{"ispecial", strconv.Itoa(int(schema.CLASS_SPECIAL)), true,
			witsjson.UnitClassJSON(schema.CLASS_SPECIAL)},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var class witsjson.UnitClassJSON
			err := json.Unmarshal([]byte(tt.encoded), &class)
			if (err != nil) == tt.succeed {
				t.Errorf("UnitInit.UnarshalJSON() error = %v, expected? %v.", err, !tt.succeed)
			}
			if !reflect.DeepEqual(class, tt.expect) {
				t.Errorf("unmarshaled = %v, want %v", class, tt.expect)
			}
		})
	}
}
