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
// github:kevindamm/wits-go/schema/player_test.go

package schema_test

import (
	"reflect"
	"testing"

	"github.com/kevindamm/wits-go/schema"
)

func TestFriendlyEnum_Opponent(t *testing.T) {
	tests := []struct {
		name string
		role schema.FriendlyEnum
		want schema.FriendlyEnum
	}{
		{"self", schema.FR_SELF, schema.FR_ENEMY},
		{"enemy", schema.FR_ENEMY, schema.FR_SELF},
		{"ally", schema.FR_ALLY, schema.FR_ENEMY},
		{"unknown", schema.FR_UNKNOWN, schema.FR_UNKNOWN},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.role.Opponent(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("FriendlyEnum.Opponent() = %v, want %v", got, tt.want)
			}
		})
	}
}
