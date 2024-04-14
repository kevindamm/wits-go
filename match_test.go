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
// github:kevindamm/wits-go/match_test.go

package wits_test

import (
	"reflect"
	"testing"

	wits "github.com/kevindamm/wits-go"
)

func TestTerminalStatus_Opposing(t *testing.T) {
	tests := []struct {
		name   string
		status wits.TerminalStatus
		want   wits.TerminalStatus
	}{
		{"unk", wits.STATUS_UNKNOWN, wits.STATUS_UNKNOWN},
		{"delay", wits.DELAY_OF_GAME, wits.DELAY_OF_GAME},
		{"destroy", wits.VICTORY_DESTRUCTION, wits.LOSS_DESTRUCTION},
		{"extinguish", wits.VICTORY_EXTINCTION, wits.LOSS_EXTINCTION},
		{"uncle", wits.VICTORY_RESIGNATION, wits.LOSS_RESIGNATION},
		{"destroyed", wits.LOSS_DESTRUCTION, wits.VICTORY_DESTRUCTION},
		{"extinct", wits.LOSS_EXTINCTION, wits.VICTORY_EXTINCTION},
		{"resigning", wits.LOSS_RESIGNATION, wits.VICTORY_RESIGNATION},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.status.Opposing(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("TerminalStatus.Opposing() = %v, want %v", got, tt.want)
			}
		})
	}
}
