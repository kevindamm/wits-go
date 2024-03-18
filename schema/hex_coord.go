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
// github:kevindamm/wits-go/schema/hex_coord.go

package schema

// A two-dimensional coordinate for map locations.  Third coordinate is implicit.
// There don't need to be format-specific representations for the raw coordinate.
type HexCoord interface {
	I() int
	J() int
}

// This is a reference to a coordinate, the GameMap can convert it into HexCoord,
// representing in one byte a wide variety of (i, j) coordinate choices.  Its
// primary benefit is an independence from unit index (pawnID in OSN) when storing
// and analyzing past replays and future play-theories.
//
// It serves its purpose as an identifier and perfect hash lookup, but at the cost
// of not being able to calculate the neighboring indices.  They could be converted
// (twice, one to lookup the coordinate, then O(n) lookups for reachable tiles).
// This mapping is static for each GameMap so they can be precomputed or memoized.
// For move-to lookups, each position can have its "by-1", "by-2", "by-..." tiles
// accessible by each distance (up to 5, for the runner).  Similarly, a set of "all"
// visible tiles" can be constructed from these index values instead of coordinate
// pairs, easier to serialize in JSON and easier to compare pairs of coordinates.
type HexCoordIndex uint8
