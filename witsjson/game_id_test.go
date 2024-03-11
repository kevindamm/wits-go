package witsjson_test

import (
	"encoding/json"
	"reflect"
	"testing"

	"github.com/kevindamm/wits-go/witsjson"
)

func TestOsnGameID_MarshalJSON(t *testing.T) {
	tests := []struct {
		name   string
		gameID witsjson.OsnGameID
		want   []byte
	}{
		{"basic quoted", witsjson.OsnGameID("gameidgameid"), []byte(`"gameidgameid"`)},
		{"with-hyphens", witsjson.OsnGameID("game-id-game-gg"), []byte(`"game-id-game-gg"`)},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var got []byte
			got, err := json.Marshal(tt.gameID)
			if err != nil {
				t.Error(err)
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("OsnGameID.MarshalJSON() = %v, want %v", got, tt.want)
			}

			var got2 witsjson.OsnGameID
			err = json.Unmarshal(got, &got2)
			if err != nil {
				t.Error(err)
			}
			if !reflect.DeepEqual(got2, tt.gameID) {
				t.Errorf("OsnGameID.UnmarshalJSON() = %v, want %v", got2, tt.gameID)
			}
		})
	}
}
