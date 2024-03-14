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
// github:kevindamm/wits-go/schema/hex_coord_test.go

package schema_test

import (
	"encoding/json"
	"reflect"
	"testing"

	"github.com/kevindamm/wits-go/schema"
)

func TestHexCoord_UnmarshalJSON(t *testing.T) {
	type args struct {
	}
	tests := []struct {
		name    string
		coord   schema.HexCoord
		encoded string
		wantErr bool
	}{
		{"basic", schema.NewHexCoord(4, -2), "[4, -2]", false},
		{"asdict", schema.NewHexCoord(4, -2), `{ "i": 4, "j": -2 }`, false},
		{"zero", schema.NewHexCoord(0, 0), "[0, 0]", false},
		{"invalid", schema.HexCoord{}, "120", true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var got schema.HexCoord
			if err := json.Unmarshal([]byte(tt.encoded), &got); (err != nil) != tt.wantErr {
				t.Errorf("HexCoord unmarshal (JSON) error = %v, wantErr %v", err, tt.wantErr)
			}
			if tt.coord.I() != got.I() || tt.coord.J() != got.J() {
				t.Errorf("HexCoord unmarshalled to incorrect values. %v != %v", got, tt.coord)
			}
		})
	}
}

func TestHexCoord_MarshalJSON(t *testing.T) {
	tests := []struct {
		name    string
		coord   schema.HexCoord
		want    string
		wantErr bool
	}{
		{"basic", schema.NewHexCoord(1, 3), "[1,3]", false},
		{"more", schema.NewHexCoord(4, 2), "[4,2]", false},
		{"big", schema.NewHexCoord(1<<10, 9000), "[1024,9000]", false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := json.Marshal(tt.coord)
			if (err != nil) != tt.wantErr {
				t.Errorf("HexCoord.MarshalJSON() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, []byte(tt.want)) {
				t.Errorf("HexCoord.MarshalJSON() = %v, want %v", string(got), string(tt.want))
			}
		})
	}
}
