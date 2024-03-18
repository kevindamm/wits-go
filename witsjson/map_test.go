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
// github:kevindamm/wits-go/witsjson/map_test.go

package witsjson_test

import (
	"encoding/json"
	"reflect"
	"testing"

	"github.com/kevindamm/wits-go/schema"
	"github.com/kevindamm/wits-go/witsjson"
)

func TestGameMap_Unmarshal(t *testing.T) {
	var glitchEncoded string = `{
		"name": "Sharkfood Island",
		"map_id": "oml/sharkfood-island",
		"terrain": {
			"floor": [
				[0, 4], [0, 5], [0, 6], [0, 7],
				[1, 3], [1, 4], [1, 5], [1, 6], [1, 7],
				[2, 4], [2, 5], [2, 8],
				[3, 3], [3, 5], [3, 7], [3, 8],
				[4, 3], [4, 4], [4, 5], [4, 6], [4, 7], [4, 8],
				[5, 2], [5, 3], [5, 4], [5, 5], [5, 6], [5, 7], [5, 8],
				[6, 2], [6, 3], [6, 5], [6, 6], [6, 8], [6, 9],
				[7, 2], [7, 3], [7, 4], [7, 5], [7, 6], [7, 7], [7, 8],
				[8, 3], [8, 4], [8, 5], [8, 6], [8, 7], [8, 8],
				[9, 3], [9, 5], [9, 7], [9, 8],
				[10, 4], [10, 5], [10, 8],
				[11, 3], [11, 4], [11, 5], [11, 6], [11, 7],
				[12, 4], [12, 5], [12, 6], [12, 7]
			],
			"wall": [
				[0, 2], [0, 3], [2, 6], [3, 6], [4, 2], [4, 9],
        [8, 2], [8, 9], [9, 6], [10, 6], [12, 2], [12, 3]
			],
			"bonus": [[3, 4], [9, 4]],
			"spawn": [[[2, 7]], [[10, 7]]],
			"base": [[2, 2], [10, 2]] 
		},
		"init": {
			"units": [
					{ "team": "RED", "class": "HEAVY", "coord": [1, 5] },
					{ "team": "RED", "class": "MEDIC", "coord": [3, 5] },
					{ "team": "RED", "class": "SOLDIER", "coord": [3, 7] },
					{ "team": "BLUE", "class": "MEDIC", "coord": [9, 5] },
					{ "team": "BLUE", "class": "SOLDIER", "coord": [9, 7] },
					{ "team": "BLUE", "class": "HEAVY", "coord": [11, 5] }
			]
		},
		"mirror": {
			"axis": 6,
			"flip": "VERTICAL",
			"center": true
		},
		"legacy": true
	}
	`

	t.Run("unmarshal glitch map", func(t *testing.T) {
		var gamemap witsjson.GameMapJSON
		if err := json.Unmarshal([]byte(glitchEncoded), &gamemap); err != nil {
			t.Errorf("GameMap JSON decode error = %v", err)
		}
		if !gamemap.IsLoaded() {
			t.Error("GameMap JSON decoded but not loaded?")
		}
		if !gamemap.Legacy() {
			t.Error("GameMap legacy bit not set.")
		}

		// Check all properties of this map.
		if gamemap.MapID() != "oml/sharkfood-island" {
			t.Errorf("GameMap JSON decode error, map ID [%v]", gamemap.MapID())
		}
		if gamemap.MapName() != "Sharkfood Island" {
			t.Errorf("GameMap JSON decode error, map Name [%v]", gamemap.MapName())
		}

		bonus := witsjson.BonusList{
			witsjson.Tile("bonus", 3, 4),
			witsjson.Tile("bonus", 9, 4),
		}
		if !reflect.DeepEqual(gamemap.Terrain().Bonus_, bonus) {
			t.Errorf("GameMap JSON decoding, terrain definitions differ %v != %v",
				gamemap.Terrain().Bonus_, bonus)
		}

		spawn := witsjson.SpawnList{
			witsjson.Spawn(2, 7, schema.FR_SELF),
			witsjson.Spawn(10, 7, schema.FR_ENEMY),
		}
		if !reflect.DeepEqual(gamemap.Terrain().Spawn_, spawn) {
			t.Errorf("GameMap JSON decoding, spawn tiles differ %v != %v",
				gamemap.Terrain().Spawn_, spawn)
		}

		base := witsjson.BaseList{
			witsjson.Base(2, 2, schema.FR_SELF),
			witsjson.Base(10, 2, schema.FR_ENEMY),
		}
		if !reflect.DeepEqual(gamemap.Terrain().Base_, base) {
			t.Errorf("GameMap JSON decoding, spawn tiles differ %v != %v",
				gamemap.Terrain().Base_, base)
		}

		floor := witsjson.FloorList{
			witsjson.Tile("floor", 0, 4),
			witsjson.Tile("floor", 0, 5),
			witsjson.Tile("floor", 0, 6),
			witsjson.Tile("floor", 0, 7),
			witsjson.Tile("floor", 1, 3),
			witsjson.Tile("floor", 1, 4),
			witsjson.Tile("floor", 1, 5),
			witsjson.Tile("floor", 1, 6),
			witsjson.Tile("floor", 1, 7),
			witsjson.Tile("floor", 2, 4),
			witsjson.Tile("floor", 2, 5),
			witsjson.Tile("floor", 2, 8),
			witsjson.Tile("floor", 3, 3),
			witsjson.Tile("floor", 3, 5),
			witsjson.Tile("floor", 3, 7),
			witsjson.Tile("floor", 3, 8),
			witsjson.Tile("floor", 4, 3),
			witsjson.Tile("floor", 4, 4),
			witsjson.Tile("floor", 4, 5),
			witsjson.Tile("floor", 4, 6),
			witsjson.Tile("floor", 4, 7),
			witsjson.Tile("floor", 4, 8),
			witsjson.Tile("floor", 5, 2),
			witsjson.Tile("floor", 5, 3),
			witsjson.Tile("floor", 5, 4),
			witsjson.Tile("floor", 5, 5),
			witsjson.Tile("floor", 5, 6),
			witsjson.Tile("floor", 5, 7),
			witsjson.Tile("floor", 5, 8),
			witsjson.Tile("floor", 6, 2),
			witsjson.Tile("floor", 6, 3),
			witsjson.Tile("floor", 6, 5),
			witsjson.Tile("floor", 6, 6),
			witsjson.Tile("floor", 6, 8),
			witsjson.Tile("floor", 6, 9),
			witsjson.Tile("floor", 7, 2),
			witsjson.Tile("floor", 7, 3),
			witsjson.Tile("floor", 7, 4),
			witsjson.Tile("floor", 7, 5),
			witsjson.Tile("floor", 7, 6),
			witsjson.Tile("floor", 7, 7),
			witsjson.Tile("floor", 7, 8),
			witsjson.Tile("floor", 8, 3),
			witsjson.Tile("floor", 8, 4),
			witsjson.Tile("floor", 8, 5),
			witsjson.Tile("floor", 8, 6),
			witsjson.Tile("floor", 8, 7),
			witsjson.Tile("floor", 8, 8),
			witsjson.Tile("floor", 9, 3),
			witsjson.Tile("floor", 9, 5),
			witsjson.Tile("floor", 9, 7),
			witsjson.Tile("floor", 9, 8),
			witsjson.Tile("floor", 10, 4),
			witsjson.Tile("floor", 10, 5),
			witsjson.Tile("floor", 10, 8),
			witsjson.Tile("floor", 11, 3),
			witsjson.Tile("floor", 11, 4),
			witsjson.Tile("floor", 11, 5),
			witsjson.Tile("floor", 11, 6),
			witsjson.Tile("floor", 11, 7),
			witsjson.Tile("floor", 12, 4),
			witsjson.Tile("floor", 12, 5),
			witsjson.Tile("floor", 12, 6),
			witsjson.Tile("floor", 12, 7),
		}
		if !reflect.DeepEqual(gamemap.Terrain().Floor_, floor) {
			t.Errorf("GameMap JSON decoding, terrain definitions differ %v != %v",
				gamemap.Terrain().Floor_, floor)
		}

		wall := witsjson.WallList{
			witsjson.Tile("wall", 0, 2),
			witsjson.Tile("wall", 0, 3),
			witsjson.Tile("wall", 2, 6),
			witsjson.Tile("wall", 3, 6),
			witsjson.Tile("wall", 4, 2),
			witsjson.Tile("wall", 4, 9),
			witsjson.Tile("wall", 8, 2),
			witsjson.Tile("wall", 8, 9),
			witsjson.Tile("wall", 9, 6),
			witsjson.Tile("wall", 10, 6),
			witsjson.Tile("wall", 12, 2),
			witsjson.Tile("wall", 12, 3),
		}
		if !reflect.DeepEqual(gamemap.Terrain().Wall_, wall) {
			t.Errorf("GameMap JSON decoding, terrain definitions differ %v != %v",
				gamemap.Terrain().Wall_, wall)
		}

		units := []schema.UnitInit{
			witsjson.UnitInitJSON{
				schema.NewHexCoord(1, 5),
				witsjson.FriendlyEnumJSON(schema.FR_SELF),
				witsjson.UnitClassJSON(schema.CLASS_HEAVY)},
			witsjson.UnitInitJSON{
				schema.NewHexCoord(3, 5),
				witsjson.FriendlyEnumJSON(schema.FR_SELF),
				witsjson.UnitClassJSON(schema.CLASS_MEDIC)},
			witsjson.UnitInitJSON{
				schema.NewHexCoord(3, 7),
				witsjson.FriendlyEnumJSON(schema.FR_SELF),
				witsjson.UnitClassJSON(schema.CLASS_SOLDIER)},
			witsjson.UnitInitJSON{
				schema.NewHexCoord(9, 5),
				witsjson.FriendlyEnumJSON(schema.FR_ENEMY),
				witsjson.UnitClassJSON(schema.CLASS_MEDIC)},
			witsjson.UnitInitJSON{
				schema.NewHexCoord(9, 7),
				witsjson.FriendlyEnumJSON(schema.FR_ENEMY),
				witsjson.UnitClassJSON(schema.CLASS_SOLDIER)},
			witsjson.UnitInitJSON{
				schema.NewHexCoord(11, 5),
				witsjson.FriendlyEnumJSON(schema.FR_ENEMY),
				witsjson.UnitClassJSON(schema.CLASS_HEAVY)},
		}
		if !reflect.DeepEqual(gamemap.Units(), units) {
			t.Errorf("GameMap JSON decoding, unit definitions differ: %v != %v",
				gamemap.Units(), units)
		}
	})
}
