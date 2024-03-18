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
// github:kevindamm/wits-go/osn/game_id_test.go

package osn_test

import (
	"encoding/json"
	"reflect"
	"testing"

	"github.com/kevindamm/wits-go/osn"
)

func TestOsnGameID_MarshalJSON(t *testing.T) {
	tests := []struct {
		name   string
		gameID osn.OsnGameID
		want   []byte
	}{
		{"basic quoted", osn.OsnGameID("gameidgameid"), []byte(`"gameidgameid"`)},
		{"with-hyphens", osn.OsnGameID("game-id-game-gg"), []byte(`"game-id-game-gg"`)},
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

			var got2 osn.OsnGameID
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

func TestOsnGameID_MarshalRemoveJunk(t *testing.T) {
	tests := []struct {
		name   string
		gameID osn.OsnGameID
		want   []byte
		want2  osn.OsnGameID
	}{
		{"with-prefix",
			osn.OsnGameID("ahRzfm91dHdpdHRlcnNnYW1lLWhyZHIVCxIIR2FtZVJvb20Ygame-id"),
			[]byte(`"game-id"`),
			osn.OsnGameID("game-id")},
		{"without-prefix",
			osn.OsnGameID("ggame-id"),
			[]byte(`"ggame-id"`),
			osn.OsnGameID("ggame-id")},
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

			var got2 osn.OsnGameID
			err = json.Unmarshal(got, &got2)
			if err != nil {
				t.Error(err)
			}
			if !reflect.DeepEqual(string(got2), string(tt.want2)) {
				t.Errorf("OsnGameID.UnmarshalJSON() = %v, want %v", got2, tt.want2)
			}
		})
	}
}
