package schema

import (
	"encoding/json"
	"fmt"
)

// A two-dimensional coordinate for map locations.  Third coordinate is implicit.
// There don't need to be format-specific representations for the raw coordinate.
type HexCoord struct {
	i int
	j int
}

func NewHexCoord(i, j int) HexCoord {
	return HexCoord{i, j}
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

func (coord HexCoord) I() int {
	return coord.i
}

func (coord HexCoord) J() int {
	return coord.j
}

// HexCoord is always marshaled as a 2D int array.  If some idiosyncratic
// json handling wants to emit it as an object or properties, use the public
// getters defined above to destructure the coordinate, but avoid doing so.
func (coord HexCoord) MarshalJSON() ([]byte, error) {
	asArray := []int{coord.i, coord.j}
	return json.Marshal(asArray)
}

// The deserialization process first tries to parse it as an int array, checking
// that it fits the dimension types.  Failing that, it will read the value as an
// (i, j)-keyed jsondict representation.
func (coord *HexCoord) UnmarshalJSON(encoded []byte) error {
	// First attempt to parse it as a list of two integers.
	array2D := make([]int, 2)
	err := json.Unmarshal(encoded, &array2D)
	if err == nil {
		if len(array2D) != 2 {
			return fmt.Errorf("HexCoord{} should have exactly 2 dimensions (found %d)", len(array2D))
		}
		coord.i = array2D[0]
		coord.j = array2D[1]
	} else {
		var jsondict struct {
			I int `json:"i"`
			J int `json:"j"`
		}
		if err := json.Unmarshal(encoded, &jsondict); err != nil {
			return err
		}
		coord.i = jsondict.I
		coord.j = jsondict.J
	}

	return nil
}

func (coord HexCoord) String() string {
	return fmt.Sprintf("<%d, %d>", coord.i, coord.j)
}
