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
// github:kevindamm/wits-go/map.go

package wits

type MapDescription interface {
	MapName() GameMapName
	MapID() GameMapID
	Terrain() TerrainDefinition
	Units() []UnitInit
	Legacy() bool
}

// The human-readable (and translatable) name for the map.  Shown in UI/views.
type GameMapName string

// The identifying string, restricted to hyphenated-alphanumeric.  Used as a
// primary key or file name for locating the map definition.
type GameMapID string

type TerrainDefinition interface {
	Floor() []TileDefinition
	Wall() []TileDefinition
	Bonus() []TileDefinition
	Spawn() []TileDefinition
	Base() []TileDefinition
}

// The position (HexCoord) and
type TileDefinition interface {
	Positional
	CanWalk() bool

	IsFloor() bool
	IsWall() bool
	IsSpawn() bool
	IsBase() bool
	IsBonus() bool

	Team() FriendlyEnum
	Typename() string
	Equals(other TileDefinition) bool
}

// non-negative integer value for measuring distance between coordinates.
type TileDistance uint8

// The hex-index and unit's essential state (all information per a coordinate).
type TileState interface {
	Placement
	UnitStateExtended
}
