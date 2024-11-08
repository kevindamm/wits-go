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
// github:kevindamm/wits-go/witsjson/player_test.go

package witsjson_test

import (
	"encoding/json"
	"reflect"
	"testing"

	"github.com/kevindamm/wits-go"
	"github.com/kevindamm/wits-go/witsjson"
)

func TestPlayerRoleJSON_Name(t *testing.T) {
	tests := []struct {
		name     string
		encoding []byte
		role     witsjson.PlayerRoleJSON
	}{
		{"basic", []byte(`{
  	  "name": "Player Name",
  		"gcID": "G:135798642",
  		"race": 1,
  		"team": 1,
			"result": 1,
			"before": {"tier": "Gifted", "rank": 25},
			"after": {"tier": "Gifted", "rank": 23, "delta": 4},
			"base_hp": 5, "wits": 0
  	}`),
			witsjson.PlayerRoleJSON{
				witsjson.PlayerID{"G:135798642"},
				"Player Name",
				witsjson.UnitRaceJSON(wits.RACE_FEEDBACK),
				witsjson.FriendlyEnumJSON(wits.FR_SELF),
				witsjson.TerminalStatusJSON(wits.VICTORY_DESTRUCTION),
				witsjson.PlayerStandingsJSON{wits.LEAGUE_TIER_INTERMEDIATE, 25},
				wits.PlayerUpdate{Tier: wits.LEAGUE_TIER_INTERMEDIATE, Rank: 23, Delta: 4},
				witsjson.BaseHealth(5), 0},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var got witsjson.PlayerRoleJSON
			err := json.Unmarshal(tt.encoding, &got)
			if err != nil {
				t.Error(err)
			}
			if !reflect.DeepEqual(got, tt.role) {
				t.Errorf("PlayerRoleJSON Unmarshal() = %v, want %v", got, tt.role)
			}
		})
	}
}
