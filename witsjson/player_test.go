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

	"github.com/kevindamm/wits-go/schema"
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
				witsjson.UnitRaceJSON(schema.RACE_FEEDBACK),
				witsjson.FriendlyEnumJSON(schema.FR_SELF),
				witsjson.TerminalStatusJSON(schema.VICTORY_DESTRUCTION),
				witsjson.PlayerStandingsJSON{Tier_: "Gifted", Rank_: 25},
				witsjson.StandingsAfterJSON{Tier_: "Gifted", Rank_: 23, Delta_: 4},
				witsjson.BaseHealth(5), 0},
		},
		//		{"basic", []byte(`{
		//  	  "name": "Player TWO",
		//  		"gcID": "G:135798642",
		//  		"race": 2,
		//  		"team": 0
		//  	}`),
		//			witsjson.PlayerRoleJSON{
		//				witsjson.PlayerID{"G:135798642"},
		//				"Player TWO",
		//				witsjson.UnitRaceJSON(schema.RACE_ADORABLES),
		//				witsjson.FriendlyEnumJSON(schema.FR_UNKNOWN)},
		//		},
		//		{"basic", []byte(`{
		//  	  "name": "Player THREE",
		//  		"gcID": "G:24601111",
		//  		"race": 3,
		//  		"team": 1
		//  	}`),
		//			witsjson.PlayerRoleJSON{
		//				witsjson.PlayerID{"G:24601111"},
		//				"Player THREE",
		//				witsjson.UnitRaceJSON(schema.RACE_SCALLYWAGS),
		//				witsjson.FriendlyEnumJSON(schema.FR_SELF)},
		//		},
		//		{"basic", []byte(`{
		//  	  "name": "Player FOUR",
		//  		"gcID": "G:135798",
		//  		"race": 4,
		//  		"team": 2
		//  	}`),
		//			witsjson.PlayerRoleJSON{
		//				witsjson.PlayerID{"G:135798"},
		//				"Player FOUR",
		//				witsjson.UnitRaceJSON(schema.RACE_VEGGIENAUTS),
		//				witsjson.FriendlyEnumJSON(schema.FR_ENEMY)},
		//		},
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
