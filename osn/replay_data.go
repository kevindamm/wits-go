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
// github:kevindamm/wits-go/osn/replay_data.go

package osn

import (
	"encoding/json"
	"fmt"
)

type ReplayData struct {
	Turns []OsnPlayerTurn
}

func (data *ReplayData) UnmarshalJSON(encoded []byte) error {
	var frames []frame
	if err := json.Unmarshal(encoded, &frames); err != nil {
		fmt.Println(err)
		return err
	}

	data.Turns = make([]OsnPlayerTurn, 0)
	turnIndex := -1

	for _, frame := range frames {
		if frame.State != nil {
			newTurn := OsnPlayerTurn{Actions: make([]OsnPlayerAction, 0)}
			newTurn.State = *frame.State
			data.Turns = append(data.Turns, newTurn)
			turnIndex += 1
		} else if frame.Action != nil {
			if turnIndex == -1 {
				if (*frame.Action).Name() != "EndTurnAction" {
					return fmt.Errorf("before-state action of type %s", (*frame.Action).Name())
				}
				continue
			}
			data.Turns[turnIndex].Actions = append(
				data.Turns[turnIndex].Actions, *frame.Action)
		}
	}
	return nil
}

// Used to hold the sum type that OSN replay frames have.
// (some actions interleaved with game states as checkpoints)
type frame struct {
	Action *OsnPlayerAction
	State  *GameState
}

type partialFrame struct {
	Action struct {
		Name string `json:"name"`
	} `json:"action"`
	State *GameState `json:"gameState"`
}

func (frame *frame) UnmarshalJSON(encoded []byte) error {
	var partial partialFrame
	err := json.Unmarshal(encoded, &partial)
	if err != nil {
		return err
	}

	if partial.Action.Name == "" {
		frame.State = partial.State
	} else if partial.State == nil {
		action, err := ParseGenericAction(partial.Action.Name, encoded)
		if err != nil {
			return err
		}
		frame.Action = &action
	}
	return nil
}

func (data *ReplayData) String() string {
	return fmt.Sprintf("Replay with %d turns; final state\n%v",
		len(data.Turns),
		data.Turns[len(data.Turns)-1].State)
}
