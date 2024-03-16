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
// github:kevindamm/wits-go/witsjson/map.go

package witsjson

import (
	"encoding/json"
	"fmt"

	"github.com/kevindamm/wits-go/schema"
)

// Game maps may be loaded as a complete definition or only referenced by name.
// The game logic must load the complete map definition in order to simulate
// the gameplay in its turns.
type GameMapJSON struct {
	id   schema.GameMapID
	defn *MapDefinition
}

// Whether the map has been loaded or not, its identifier is known.
func (gamemap GameMapJSON) MapID() schema.GameMapID {
	return schema.GameMapID(gamemap.id)
}

// Returns true if the full map definition has been loaded from disk.
func (m GameMapJSON) IsLoaded() bool {
	// Any valid map must have at least one traversible tile.
	return m.defn != nil && len(m.defn.Terrain.Floor) > 0
}

func (gamemap GameMapJSON) MapName() schema.GameMapName {
	if !gamemap.IsLoaded() {
		return ""
	}
	return schema.GameMapName(gamemap.defn.Name)
}

func (gamemap GameMapJSON) Terrain() TerrainDefinition {
	if !gamemap.IsLoaded() {
		return TerrainDefinition{}
	}
	return gamemap.defn.Terrain
}

func (gamemap GameMapJSON) Units() []schema.UnitInit {
	if gamemap.defn == nil {
		return nil
	}
	units := make([]schema.UnitInit, len(gamemap.defn.Init.Units))
	for i, unit := range gamemap.defn.Init.Units {
		units[i] = schema.UnitInit(unit)
	}
	return units
}

func (gamemap GameMapJSON) Legacy() bool {
	if gamemap.defn == nil || gamemap.defn.Legacy == nil {
		return false
	}
	return *gamemap.defn.Legacy
}

// We only need to marshal the game's name,
// its definition is held (once) in a separate file.
// To marshal the file's content, encode the MapDefinition type.
func (m GameMapJSON) MarshalJSON() ([]byte, error) {
	return json.Marshal(m.id)
}

// When unmarshaling, we first attempt to read it as a string (as it would be
// when found in a game replay encoding).  If it is not a string, we unmarshal
// the map definition (as it appears in the map file representation).
func (m *GameMapJSON) UnmarshalJSON(encoded []byte) error {
	var mapID schema.GameMapID
	if err := json.Unmarshal(encoded, &mapID); err == nil {
		m.id = schema.GameMapID(mapID)
		return nil
	}

	var data MapDefinition
	if err := json.Unmarshal(encoded, &data); err != nil {
		return err
	}
	m.id = schema.GameMapID(data.MapID)
	m.defn = &data
	return nil
}

type MapDefinition struct {
	MapID   string            `json:"map_id"`
	Name    string            `json:"name"`
	Terrain TerrainDefinition `json:"terrain"`
	Init    MapInit           `json:"init"`
	Rotate  *Rotation         `json:"rotate,omitempty"`
	Mirror  *Reflection       `json:"mirror,omitempty"`
	Legacy  *bool             `json:"legacy,omitempty"`
}

type TerrainDefinition struct {
	Floor FloorList `json:"floor"`
	Wall  WallList  `json:"wall"`
	Bonus BonusList `json:"bonus"`
	Spawn SpawnList `json:"spawn"`
	Base  BaseList  `json:"base"`
}

func Tile(terrain schema.MapTerrain, i, j int) schema.TileDefinition {
	return TileDefinition{terrain, schema.NewHexCoord(i, j)}
}

type TileDefinition struct {
	schema.MapTerrain
	schema.HexCoord
}

func (def TileDefinition) Position() schema.HexCoord {
	return def.HexCoord
}

func (def TileDefinition) Terrain() schema.MapTerrain {
	return def.MapTerrain
}

// A generic approach to unmarshaling each list of coordinates into their
// representative terrain type.  This provides the benefit of typed unmarshaling
// of the JSON representation without the tedium of repetitive code.
func UnmarshalTerrain[T ~[]schema.TileDefinition](
	encoded []byte, enum schema.MapTerrain, defs *T) error {

	list := make([]schema.TileDefinition, 0)
	var coords [][]int
	if err := json.Unmarshal(encoded, &coords); err != nil {
		return err
	}

	for _, coord := range coords {
		if len(coord) != 2 {
			return fmt.Errorf("floor coordinate with incorrect dimensions %v", coord)
		}
		list = append(list, Tile(enum, coord[0], coord[1]))
	}

	*defs = T(list)
	return nil
}

// Marshaling the types back into a list can also be done in a generic way.
// The list of tile definitions may contain any terrain type, this method will
// both filter and encode while keeping to the signature of JSON marshaling.
// When filtering, it handles common classes (base, spawn, bonus) together.
func MarshalTerrain[T ~[]schema.TileDefinition](
	enum schema.MapTerrain, data *T) ([]byte, error) {
	coords := make([][]int, 0)

	switch enum & schema.TERRAIN_TYPE_MASK {
	case schema.TERRAIN_TYPE_BONUS:
		// bonus tiles, avoid the 'UNKNOWN' representation
		for _, item := range *data {
			terraintype := item.Terrain() & schema.TERRAIN_TYPE_MASK
			if terraintype == schema.TERRAIN_TYPE_BONUS && enum != schema.TERRAIN_UNKNOWN {
				pos := item.Position()
				coords = append(coords, []int{pos.I(), pos.J()})
			}
		}
	case schema.TERRAIN_TYPE_SPAWN:
		// match any type of 'spawn'
		for _, item := range *data {
			terraintype := item.Terrain() & schema.TERRAIN_TYPE_MASK
			if terraintype == schema.TERRAIN_TYPE_SPAWN {
				pos := item.Position()
				coords = append(coords, []int{pos.I(), pos.J()})
			}
		}
	case schema.TERRAIN_TYPE_BASE:
		// match any type of 'base'
		for _, item := range *data {
			terraintype := item.Terrain() & schema.TERRAIN_TYPE_MASK
			if terraintype == schema.TERRAIN_TYPE_BASE {
				pos := item.Position()
				coords = append(coords, []int{pos.I(), pos.J()})
			}
		}
	case schema.TERRAIN_TYPE_OTHER:
		// 'other' type needs an exact match
		for _, item := range *data {
			if item.Terrain() == enum {
				pos := item.Position()
				coords = append(coords, []int{pos.I(), pos.J()})
			}
		}
	}

	return json.Marshal(coords)
}

// Unpacked from a JSON of []HexCoord into a []TileDefinition of TERRAIN_TYPE_FLOOR.
type FloorList []schema.TileDefinition

// Unmarshals the list of coordinates for floor positions.
func (defs *FloorList) UnmarshalJSON(encoded []byte) error {
	return UnmarshalTerrain(encoded, schema.TERRAIN_FLOOR, defs)
}

func (defs *FloorList) MarshalJSON() ([]byte, error) {
	return MarshalTerrain(schema.TERRAIN_FLOOR, defs)
}

// Unpacked from a JSON of []HexCoord into a []TileDefinition of TERRAIN_TYPE_WALL.
type WallList []schema.TileDefinition

// Unmarshals the list of coordinates for wall positions.
func (defs *WallList) UnmarshalJSON(encoded []byte) error {
	return UnmarshalTerrain(encoded, schema.TERRAIN_WALL, defs)
}

func (defs *WallList) MarshalJSON() ([]byte, error) {
	return MarshalTerrain(schema.TERRAIN_WALL, defs)
}

// Unpacked from a JSON of []HexCoord into a []TileDefinition of TERRAIN_TYPE_BONUS.
type BonusList []schema.TileDefinition

// Unmarshals the list of coordinates for bonus positions.
func (defs *BonusList) UnmarshalJSON(encoded []byte) error {
	return UnmarshalTerrain(encoded, schema.TERRAIN_BONUS_NEUTRAL, defs)
}

func (defs *BonusList) MarshalJSON() ([]byte, error) {
	return MarshalTerrain(schema.TERRAIN_TYPE_BONUS, defs)
}

// Unpacked from a JSON of []HexCoord into a []TileDefinition of TERRAIN_TYPE_SPAWN.
type SpawnList []schema.TileDefinition

// Unmarshals the list of coordinates for spawn positions.
func (defs *SpawnList) UnmarshalJSON(encoded []byte) error {
	var all_teams [][]schema.HexCoord
	if err := json.Unmarshal(encoded, &all_teams); err != nil {
		return err
	}
	colorSpawns := func(spawns []schema.HexCoord, terrain schema.MapTerrain) {
		for _, spawn := range spawns {
			*defs = append(*defs, Tile(
				terrain,
				spawn.I(),
				spawn.J()))
		}
	}

	colorSpawns(all_teams[0], schema.TERRAIN_SPAWN_RED)
	colorSpawns(all_teams[1], schema.TERRAIN_SPAWN_BLUE)
	if len(all_teams) > 2 {
		colorSpawns(all_teams[2], schema.TERRAIN_SPAWN_GOLD)
		colorSpawns(all_teams[3], schema.TERRAIN_SPAWN_GREEN)
	}
	return nil
}

func (defs *SpawnList) MarshalJSON() ([]byte, error) {
	return MarshalTerrain(schema.TERRAIN_TYPE_SPAWN, defs)
}

// Unpacked from a JSON of []HexCoord into a []TileDefinition of TERRAIN_TYPE_BASE.
type BaseList []schema.TileDefinition

// Unmarshals the list of coordinates for base positions.
func (defs *BaseList) UnmarshalJSON(encoded []byte) error {
	if err := UnmarshalTerrain(encoded, schema.TERRAIN_TYPE_BASE, defs); err != nil {
		return err
	}
	for i, def := range *defs {
		base_type := schema.TERRAIN_BASE_RED + schema.MapTerrain(i)
		(*defs)[i] = Tile(base_type, def.Position().I(), def.Position().J())
	}

	return nil
}

func (defs *BaseList) MarshalJSON() ([]byte, error) {
	return MarshalTerrain(schema.TERRAIN_TYPE_BASE, defs)
}

// Initialization of map-related game state that is not terrain related.
type MapInit struct {
	Units []UnitInitJSON `json:"units"`
}

type Rotation struct {
	Position schema.HexCoord `json:"position"`
	Center   bool            `json:"center"`
}

type Reflection struct {
	Axis   int            `json:"axis"`
	Flip   ReflectionType `json:"flip"`
	Cetner bool           `json:"center"`
}

type ReflectionType string

type TileDistance int
