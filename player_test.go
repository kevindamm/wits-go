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
// github:kevindamm/wits-go/player_test.go

package wits_test

import (
	"reflect"
	"testing"

	"github.com/kevindamm/wits-go"
)

func TestFriendlyEnum_Opponent(t *testing.T) {
	tests := []struct {
		name string
		role wits.FriendlyEnum
		want wits.FriendlyEnum
	}{
		{"self", wits.FR_SELF, wits.FR_ENEMY},
		{"enemy", wits.FR_ENEMY, wits.FR_SELF},
		{"ally", wits.FR_ALLY, wits.FR_ENEMY},
		{"enemy", wits.FR_ENEMY2, wits.FR_SELF},
		{"unknown", wits.FR_UNKNOWN, wits.FR_UNKNOWN},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.role.Opponent(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("FriendlyEnum.Opponent() = %v, want %v", got, tt.want)
			}
		})
	}
}
