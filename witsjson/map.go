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

	"github.com/kevindamm/wits-go"
)

// Game maps may be loaded as a complete definition or only referenced by name.
// The game logic must load the complete map definition in order to simulate
// the gameplay in its turns.
type GameMapJSON struct {
	id   wits.GameMapID
	defn *MapDefinition
}

// Whether the map has been loaded or not, its identifier is known.
func (gamemap GameMapJSON) MapID() wits.GameMapID {
	return wits.GameMapID(gamemap.id)
}

// Returns true if the full map definition has been loaded from disk.
func (m GameMapJSON) IsLoaded() bool {
	// Any valid map must have at least one traversible tile.
	return m.defn != nil && len(m.defn.Terrain.Floor_) > 0
}

func (gamemap GameMapJSON) MapName() wits.GameMapName {
	if !gamemap.IsLoaded() {
		return ""
	}
	return wits.GameMapName(gamemap.defn.Name)
}

func (gamemap GameMapJSON) Terrain() wits.TerrainDefinition {
	if !gamemap.IsLoaded() {
		return TerrainDefinition{}
	}
	return gamemap.defn.Terrain
}

func (gamemap GameMapJSON) Units() []wits.UnitInit {
	if !gamemap.IsLoaded() {
		return nil
	}
	units := make([]wits.UnitInit, len(gamemap.defn.Init.Units))
	for i, unit := range gamemap.defn.Init.Units {
		units[i] = wits.UnitInit(unit)
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
	var mapID wits.GameMapID
	if err := json.Unmarshal(encoded, &mapID); err == nil {
		m.id = wits.GameMapID(mapID)
		return nil
	}

	var data MapDefinition
	if err := json.Unmarshal(encoded, &data); err != nil {
		return err
	}
	m.id = wits.GameMapID(data.MapID)
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
	Floor_ FloorList `json:"floor"`
	Wall_  WallList  `json:"wall"`
	Bonus_ BonusList `json:"bonus"`
	Spawn_ SpawnList `json:"spawn"`
	Base_  BaseList  `json:"base"`
}

func NewTile(terrain string, i, j int) wits.TileDefinition {
	return map[string]TerrainType{
		"FLOOR": FloorList_New,
		"WALL":  WallList_New,
		"BONUS": BonusList_New,
	}[terrain](NewHexCoord(i, j))
}
func NewSpawn(i, j int, team wits.FriendlyEnum) wits.TileDefinition {
	return spawn{NewHexCoord(i, j), team}
}
func NewBase(i, j int, team wits.FriendlyEnum) wits.TileDefinition {
	return base{NewHexCoord(i, j), team}
}

func (terrain TerrainDefinition) Floor() []wits.TileDefinition {
	return coerceTileDefs(terrain.Floor_)
}

func (terrain TerrainDefinition) Wall() []wits.TileDefinition {
	return coerceTileDefs(terrain.Wall_)
}

func (terrain TerrainDefinition) Bonus() []wits.TileDefinition {
	return coerceTileDefs(terrain.Bonus_)
}

func (terrain TerrainDefinition) Spawn() []wits.TileDefinition {
	return coerceTileDefs(terrain.Spawn_)
}

func (terrain TerrainDefinition) Base() []wits.TileDefinition {
	return coerceTileDefs(terrain.Base_)
}

// A generic approach to unmarshaling each list of coordinates into their
// representative terrain type.  This provides the benefit of typed unmarshaling
// of the JSON representation without the tedium of repetitive code.
func UnmarshalTerrain[T ~[]wits.TileDefinition](
	encoded []byte, enum string, defs *T) error {

	list := make([]wits.TileDefinition, 0)
	var coords [][]int
	if err := json.Unmarshal(encoded, &coords); err != nil {
		return err
	}

	for _, coord := range coords {
		if len(coord) != 2 {
			return fmt.Errorf("terrain coordinate with incorrect dimensions %v", coord)
		}
		list = append(list, NewTile(enum, coord[0], coord[1]))
	}

	*defs = T(list)
	return nil
}

// Marshaling the types back into a list can also be done in a generic way.
// The list of tile definitions may contain any terrain type, this method will
// both filter and encode while keeping to the signature of JSON marshaling.
// When filtering, it handles common classes (base, spawn, bonus) together.
func MarshalTerrain[T ~[]wits.TileDefinition](
	enum string, data *T) ([]byte, error) {
	coords := make([][]int, 0)

	switch enum {
	case "FLOOR":
		// match all the floor tiles
		for _, item := range *data {
			if item.IsFloor() {
				pos := item.Position()
				coords = append(coords, []int{pos.I(), pos.J()})
			}
		}
	case "WALL":
		// match all the wall tiles
		for _, item := range *data {
			if item.IsWall() {
				pos := item.Position()
				coords = append(coords, []int{pos.I(), pos.J()})
			}
		}
	case "BONUS":
		for _, item := range *data {
			if item.IsBonus() {
				pos := item.Position()
				coords = append(coords, []int{pos.I(), pos.J()})
			}
		}
	case "SPAWN":
		// match any type of 'spawn' and collate by team.
		teamspawns := make([][][]int, 2)
		teamspawns[0] = make([][]int, 0)
		teamspawns[1] = make([][]int, 0)
		for _, item := range *data {
			if item.IsSpawn() {
				pos := item.Position()
				team := item.Team()
				if (team == wits.FR_ALLY || team == wits.FR_ENEMY2) && len(teamspawns) == 2 {
					teamspawns = grow_spawnslist(teamspawns)
				}
				index := int(team - wits.FR_SELF)
				teamspawns[index] = append(teamspawns[index], []int{pos.I(), pos.J()})
			}
		}
	case "BASE":
		// match any type of 'base' and index by team
		teambase := make([][]int, 2)
		teambase[0] = make([]int, 0)
		teambase[1] = make([]int, 0)
		for _, item := range *data {
			if item.IsBase() {
				pos := item.Position()
				team := item.Team()
				if (team == wits.FR_ALLY || team == wits.FR_ENEMY2) && len(teambase) == 2 {
					teambase = grow_baselist(teambase)
				}
				index := int(team - wits.FR_SELF)
				teambase[index] = []int{pos.I(), pos.J()}
			}
		}
	}

	return json.Marshal(coords)
}

func coerceTileDefs[T ~[]wits.TileDefinition](tiles T) []wits.TileDefinition {
	casted := make([]wits.TileDefinition, len(tiles))
	for i, tile := range tiles {
		casted[i] = wits.TileDefinition(tile)
	}
	return casted
}

// Functor for a curried constructor of a specific tile type.
type TerrainType func(HexCoordJSON) wits.TileDefinition

// Unpacked from a JSON of []HexCoord into a []TileDefinition of TERRAIN_TYPE_FLOOR.
type FloorList []wits.TileDefinition
type floor HexCoordJSON

func FloorList_New(coord HexCoordJSON) wits.TileDefinition {
	return floor(coord)
}

func (tile floor) Position() wits.HexCoord {
	return HexCoordJSON(tile)
}
func (tile floor) CanWalk() bool           { return true }
func (tile floor) IsFloor() bool           { return true }
func (tile floor) IsWall() bool            { return false }
func (tile floor) IsSpawn() bool           { return false }
func (tile floor) IsBase() bool            { return false }
func (tile floor) IsBonus() bool           { return false }
func (tile floor) Team() wits.FriendlyEnum { return wits.FR_UNKNOWN }
func (tile floor) Typename() string        { return "FLOOR" }
func (tile floor) Equals(other wits.TileDefinition) bool {
	return (other.IsFloor() &&
		other.Position().I() == tile.Position().I() &&
		other.Position().J() == tile.Position().J())
}

// Unmarshals the list of coordinates for floor positions.
func (defs *FloorList) UnmarshalJSON(encoded []byte) error {
	return UnmarshalTerrain(encoded, "FLOOR", defs)
}

func (defs *FloorList) MarshalJSON() ([]byte, error) {
	return MarshalTerrain("FLOOR", defs)
}

// Unpacked from a JSON of []HexCoord into a []TileDefinition of TERRAIN_TYPE_WALL.
type WallList []wits.TileDefinition
type wall HexCoordJSON

func WallList_New(coord HexCoordJSON) wits.TileDefinition {
	return wall(coord)
}

func (tile wall) Position() wits.HexCoord {
	return HexCoordJSON(tile)
}
func (tile wall) CanWalk() bool           { return false }
func (tile wall) IsFloor() bool           { return false }
func (tile wall) IsWall() bool            { return true }
func (tile wall) IsSpawn() bool           { return false }
func (tile wall) IsBase() bool            { return false }
func (tile wall) IsBonus() bool           { return false }
func (tile wall) Team() wits.FriendlyEnum { return wits.FR_UNKNOWN }
func (tile wall) Typename() string        { return "WALL" }
func (tile wall) Equals(other wits.TileDefinition) bool {
	return (other.IsFloor() &&
		other.Position().I() == tile.Position().I() &&
		other.Position().J() == tile.Position().J())
}

// Unmarshals the list of coordinates for wall positions.
func (defs *WallList) UnmarshalJSON(encoded []byte) error {
	return UnmarshalTerrain(encoded, "WALL", defs)
}

func (defs *WallList) MarshalJSON() ([]byte, error) {
	return MarshalTerrain("WALL", defs)
}

// Unpacked from a JSON of []HexCoord into a []TileDefinition of TERRAIN_TYPE_SPAWN.
type SpawnList []wits.TileDefinition

type spawn struct {
	wits.HexCoord
	wits.FriendlyEnum
}

func (tile spawn) Position() wits.HexCoord {
	return tile.HexCoord
}
func (tile spawn) CanWalk() bool           { return true }
func (tile spawn) IsFloor() bool           { return false }
func (tile spawn) IsWall() bool            { return false }
func (tile spawn) IsSpawn() bool           { return true }
func (tile spawn) IsBase() bool            { return false }
func (tile spawn) IsBonus() bool           { return false }
func (tile spawn) Team() wits.FriendlyEnum { return tile.FriendlyEnum }
func (tile spawn) Typename() string        { return "SPAWN" }
func (tile spawn) Equals(other wits.TileDefinition) bool {
	return (other.IsSpawn() &&
		other.Position().I() == tile.Position().I() &&
		other.Position().J() == tile.Position().J())
}

// Grow a two-team (solo) list of list of positions into a four-team (duos) list of lists.
func grow_spawnslist(teamspawns_solo [][][]int) [][][]int {
	teamspawns_duos := make([][][]int, 4)
	teamspawns_duos[0] = teamspawns_solo[0]
	teamspawns_duos[1] = teamspawns_solo[1]
	teamspawns_duos[2] = make([][]int, 0)
	teamspawns_duos[3] = make([][]int, 0)
	return teamspawns_duos
}

// Unmarshals the list of coordinates for spawn positions.
func (defs *SpawnList) UnmarshalJSON(encoded []byte) error {
	var all_spawns [][]HexCoordJSON
	if err := json.Unmarshal(encoded, &all_spawns); err != nil {
		return err
	}

	tiles := make([]wits.TileDefinition, 0)
	for i, spawnlist := range all_spawns {
		for _, coord := range spawnlist {
			tiles = append(tiles,
				spawn{coord, wits.FR_SELF + wits.FriendlyEnum(i)})
		}
	}

	*defs = tiles
	return nil
}

func (defs *SpawnList) MarshalJSON() ([]byte, error) {
	return MarshalTerrain("SPAWN", defs)
}

// Unpacked from a JSON of []HexCoord into a []TileDefinition of TERRAIN_TYPE_BASE.
type BaseList []wits.TileDefinition
type base struct {
	wits.HexCoord
	wits.FriendlyEnum
}

func (tile base) Position() wits.HexCoord {
	return tile.HexCoord
}
func (tile base) CanWalk() bool           { return false }
func (tile base) IsFloor() bool           { return false }
func (tile base) IsWall() bool            { return false }
func (tile base) IsSpawn() bool           { return false }
func (tile base) IsBase() bool            { return true }
func (tile base) IsBonus() bool           { return false }
func (tile base) Team() wits.FriendlyEnum { return tile.FriendlyEnum }
func (tile base) Typename() string        { return "BASE" }
func (tile base) Equals(other wits.TileDefinition) bool {
	return (other.IsBase() &&
		other.Position().I() == tile.Position().I() &&
		other.Position().J() == tile.Position().J())
}

// Grow a two-team (solo) list of list of positions into a four-team (duos) list of lists.
func grow_baselist(teambase_solo [][]int) [][]int {
	teambase_duos := make([][]int, 4)
	teambase_duos[0] = teambase_solo[0]
	teambase_duos[1] = teambase_solo[1]
	teambase_duos[2] = make([]int, 0)
	teambase_duos[3] = make([]int, 0)
	return teambase_duos
}

// Unmarshals the list of coordinates for base positions.
func (defs *BaseList) UnmarshalJSON(encoded []byte) error {
	var bases []HexCoordJSON
	if err := json.Unmarshal(encoded, &bases); err != nil {
		return err
	}

	tiles := make([]wits.TileDefinition, 0)
	for i, coord := range bases {
		tiles = append(tiles,
			base{coord, wits.FR_SELF + wits.FriendlyEnum(i)})
	}

	*defs = tiles
	return nil
}

func (defs *BaseList) MarshalJSON() ([]byte, error) {
	return MarshalTerrain("BASE", defs)
}

// Unpacked from a JSON of []HexCoord into a []TileDefinition of TERRAIN_TYPE_BONUS.
type BonusList []wits.TileDefinition
type bonus HexCoordJSON

func BonusList_New(coord HexCoordJSON) wits.TileDefinition {
	return bonus(coord)
}

func (tile bonus) Position() wits.HexCoord {
	return HexCoordJSON(tile)
}
func (tile bonus) CanWalk() bool           { return true }
func (tile bonus) IsFloor() bool           { return false }
func (tile bonus) IsWall() bool            { return false }
func (tile bonus) IsSpawn() bool           { return false }
func (tile bonus) IsBase() bool            { return false }
func (tile bonus) IsBonus() bool           { return true }
func (tile bonus) Team() wits.FriendlyEnum { return wits.FR_UNKNOWN }
func (tile bonus) Typename() string        { return "BONUS" }
func (tile bonus) Equals(other wits.TileDefinition) bool {
	return (other.IsBonus() &&
		other.Position().I() == tile.Position().I() &&
		other.Position().J() == tile.Position().J())
}

// Unmarshals the list of coordinates for bonus positions.
func (defs *BonusList) UnmarshalJSON(encoded []byte) error {
	return UnmarshalTerrain(encoded, "BONUS", defs)
}

func (defs *BonusList) MarshalJSON() ([]byte, error) {
	return MarshalTerrain("BONUS", defs)
}

// Initialization of map-related game state that is not terrain related.
type MapInit struct {
	Units []UnitInitJSON `json:"units"`
}

type Rotation struct {
	Position HexCoordJSON `json:"position"`
	Center   bool         `json:"center"`
}

type Reflection struct {
	Axis   int            `json:"axis"`
	Flip   ReflectionType `json:"flip"`
	Cetner bool           `json:"center"`
}

type ReflectionType string

type TileDistance int
