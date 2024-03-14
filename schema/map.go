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
// github:kevindamm/wits-go/schema/map.go

package schema

type GameMap interface {
	MapName() GameMapName
	MapID() GameMapID
	Terrain() map[HexCoordIndex]TileDefinition
	Units() []UnitInit
	Legacy() bool
}

// The human-readable (and translatable) name for the map.  Shown in UI/views.
type GameMapName string

// The identifying string, restricted to hyphenated-alphanumeric.  Used as a
// primary key or file name for locating the map definition.
type GameMapID string

// The position (HexCoord) and
type TileDefinition interface {
	Positional
	Terrain() MapTerrain
}

type TileDistance uint8

// The hex-index and unit's essential state (all information per a coordinate).
type TileState interface {
	Placement
	UnitStateExtended
}

// Each valid tile, traversible or blocked, is represented by a common TERRAIN_*
// enum.  They are partitioned and stored by type, coordinate when serialized into
// JSON whereby the enum evaporates into
type MapTerrain byte

// Conveniently, these fit within just four bits, even accommodating for esoteric
// values such as a hole in the floor or a stand-in for "occupied by unit" that is
// distinct from wall (a dynamic terrain, but not expected to be seen in persisted
// instances).  Ordering them in this way provides a convenient bitmask for certain
// properties:
//
//  [_ _][_][_]
//     |  |`-`--- alliance
//     |  `------ visibility/duos
//     `--------- tile class
//
// This loses the information about variations on floor textures and base designs.
// Those details are not essential to the simulation or in deciding strategy, and
// so do not belong with this representation.  A frontend can choose to alter these
// aspects without interfering with this enumeration or the overall map definition.

const (
	TERRAIN_UNKNOWN       MapTerrain = iota
	TERRAIN_BONUS_NEUTRAL            // .. 0b01
	TERRAIN_BONUS_RED                // .. 0b10
	TERRAIN_BONUS_BLUE               // .. 0b11
	TERRAIN_WALL                     // 0b_0100
	TERRAIN_OCCUPIED                 // 0b_0101
	TERRAIN_HOLE                     // 0b_0110
	TERRAIN_FLOOR                    // 0b_0111
	TERRAIN_BASE_RED                 // 0b_1000
	TERRAIN_BASE_BLUE                // 0b_1001
	TERRAIN_BASE_GOLD                // 0b_1010
	TERRAIN_BASE_GREEN               // 0b_1011
	TERRAIN_SPAWN_RED                // 0b_1100
	TERRAIN_SPAWN_BLUE               // 0b_1101
	TERRAIN_SPAWN_GOLD               // 0b_1110
	TERRAIN_SPAWN_GREEN              // 0b_1111
)

const (
	TERRAIN_TYPE_MASK  MapTerrain = 0b1100
	TERRAIN_TYPE_SPAWN MapTerrain = 3 << 2
	TERRAIN_TYPE_BASE  MapTerrain = 2 << 2
	TERRAIN_TYPE_OTHER MapTerrain = 1 << 2
	TERRAIN_TYPE_BONUS MapTerrain = 0 // except UNKNOWN
)

func ParseTerrain(terrainEnum string) MapTerrain {
	return map[string]MapTerrain{
		"UNKNOWN":     TERRAIN_UNKNOWN,
		"BONUS":       TERRAIN_BONUS_NEUTRAL,
		"BONUS_RED":   TERRAIN_BONUS_RED,
		"BONUS_BLUE":  TERRAIN_BONUS_BLUE,
		"WALL":        TERRAIN_WALL,
		"OCCUPIED":    TERRAIN_OCCUPIED,
		"HOLE":        TERRAIN_HOLE,
		"FLOOR":       TERRAIN_FLOOR,
		"BASE_RED":    TERRAIN_BASE_RED,
		"BASE_BLUE":   TERRAIN_BASE_BLUE,
		"BASE_GOLD":   TERRAIN_BASE_GOLD,
		"BASE_GREEN":  TERRAIN_BASE_GREEN,
		"SPAWN_RED":   TERRAIN_SPAWN_RED,
		"SPAWN_BLUE":  TERRAIN_SPAWN_BLUE,
		"SPAWN_GOLD":  TERRAIN_SPAWN_GOLD,
		"SPAWN_GREEN": TERRAIN_SPAWN_GREEN,
	}[terrainEnum]
}

func (terrain MapTerrain) String() string {
	return []string{
		"UNKNOWN",
		"BONUS",
		"BONUS_RED",
		"BONUS_BLUE",
		"WALL",
		"OCCUPIED",
		"HOLE",
		"FLOOR",
		"BASE_RED",
		"BASE_BLUE",
		"BASE_GOLD",
		"BASE_GREEN",
		"SPAWN_RED",
		"SPAWN_BLUE",
		"SPAWN_GOLD",
		"SPAWN_GREEN",
	}[int(terrain)]
}
