package schema_test

import (
	"encoding/json"
	"reflect"
	"testing"

	"github.com/kevindamm/wits-go/schema"
)

func TestHexCoord_UnmarshalJSON(t *testing.T) {
	type args struct {
		encoded string
	}
	tests := []struct {
		name    string
		coord   schema.HexCoord
		args    args
		wantErr bool
	}{
		{"basic", schema.NewHexCoord(4, -2), args{"[4, -2]"}, false},
		{"asdict", schema.NewHexCoord(4, -2), args{`{ "i": 4, "j": -2 }`}, false},
		{"zero", schema.NewHexCoord(0, 0), args{"[0, 0]"}, false},
		{"invalid", schema.HexCoord{}, args{"120"}, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var got schema.HexCoord
			if err := json.Unmarshal([]byte(tt.args.encoded), &got); (err != nil) != tt.wantErr {
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
