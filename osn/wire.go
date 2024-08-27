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
// github:kevindamm/wits-go/osn/actions.go

package osn

import (
	"encoding/json"
	"io"
)

// This is the format as returned by the web service for a single game replay.
//
// It is a shallow wrapper around the actual game replay, containing four
// dictionary entries, one of which is a string representation of the game.
type WireFormat struct {
	Wrapper OuterWrapper `json:"viewResponse"`
}

type OuterWrapper struct {
	Wrapper string `json:"gameState"`    // Wrapper around the turn's frames.
	Named   string `json:"viewResponse"` // Path reference of an `.hxm` file.
	Found   bool   `json:"foundRoom"`    // Always true - a game lobby was found.
	RoomID  string `json:"room"`         // The game (and game replay) identifier.
}

type InnerWrapper struct {
	Wrapper string `json:"gameState"`
}

// Start with a reasonable slice length when reading the serialized bytes.
const BUFFER_SIZE_INIT int = 1 << 14

func ParseReplay(reader io.Reader) (GameReplay, error) {
	var on_wire WireFormat
	var gamestate InnerWrapper
	var replay GameReplay
	contents := make([]byte, 0, BUFFER_SIZE_INIT)
	_, err := reader.Read(contents)
	if err != nil {
		return replay, err
	}
	err = json.Unmarshal(contents, &on_wire)
	if err != nil {
		return replay, err
	}
	err = json.Unmarshal([]byte(on_wire.Wrapper.Wrapper), &gamestate)
	if err != nil {
		return replay, err
	}
	err = json.Unmarshal([]byte(gamestate.Wrapper), &replay)
	return replay, err
}
