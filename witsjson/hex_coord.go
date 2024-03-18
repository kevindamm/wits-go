package witsjson

import (
	"encoding/json"
	"fmt"
)

// Satisfies schema.HexCoord and serializes to a 2D-array.  When deseriailzing,
// asserts that the list/slice has exactly two elements, returning an error.
type HexCoordJSON struct {
	i int // first element in the list
	j int // second element in the list
}

func (coord HexCoordJSON) I() int {
	return coord.i
}

func (coord HexCoordJSON) J() int {
	return coord.j
}

// Trivial constructor function, exposes a one-time setter for the coordinates.
func NewHexCoord(i, j int) HexCoordJSON {
	return HexCoordJSON{i, j}
}

// Marshals the coordinate as its 2D list representation.
func (coord HexCoordJSON) MarshalJSON() ([]byte, error) {
	asArray := []int{coord.i, coord.j}
	return json.Marshal(asArray)
}

// The deserialization process first tries to parse it as an int array, checking
// that it fits the dimension types.  Failing that, it will read the value as an
// (i, j)-keyed jsondict representation.  Otherwise returns an error.
func (coord *HexCoordJSON) UnmarshalJSON(encoded []byte) error {
	// First, attempt to parse it as a list of two integers.
	array2D := make([]int, 2)
	err := json.Unmarshal(encoded, &array2D)
	if err == nil {
		if len(array2D) != 2 {
			return fmt.Errorf("a HexCoord{} should have exactly 2 dimensions (found %d)", len(array2D))
		}
		coord.i = array2D[0]
		coord.j = array2D[1]
	} else {
		// Otherwise try to unmarshal as a struct representation.
		var jsondict struct {
			I int `json:"i"`
			J int `json:"j"`
		}
		if err2 := json.Unmarshal(encoded, &jsondict); err2 != nil {
			return fmt.Errorf("could not unmarshal HexCoord,\n  %s\n,\n  %s\n---\n", err, err2)
		}
		coord.i = jsondict.I
		coord.j = jsondict.J
	}

	return nil
}

// Simple string representation that formats it like a vector.
// Not meant to be its parseable form.
func (coord HexCoordJSON) String() string {
	return fmt.Sprintf("<%d, %d>", coord.i, coord.j)
}
