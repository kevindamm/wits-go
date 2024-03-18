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
// github:kevindamm/wits-go/schema/status_test.go

package schema

import (
	"reflect"
	"testing"
)

func TestTerminalStatus_Opposing(t *testing.T) {
	tests := []struct {
		name   string
		status TerminalStatus
		want   TerminalStatus
	}{
		{"unk", STATUS_UNKNOWN, STATUS_UNKNOWN},
		{"delay", DELAY_OF_GAME, DELAY_OF_GAME},
		{"destroy", VICTORY_DESTRUCTION, LOSS_DESTRUCTION},
		{"extinguish", VICTORY_EXTINCTION, LOSS_EXTINCTION},
		{"uncle", VICTORY_RESIGNATION, LOSS_RESIGNATION},
		{"destroyed", LOSS_DESTRUCTION, VICTORY_DESTRUCTION},
		{"extinct", LOSS_EXTINCTION, VICTORY_EXTINCTION},
		{"resigning", LOSS_RESIGNATION, VICTORY_RESIGNATION},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.status.Opposing(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("TerminalStatus.Opposing() = %v, want %v", got, tt.want)
			}
		})
	}
}
