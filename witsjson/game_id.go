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
// github:kevindamm/wits-go/witsjson/game_id.go

package witsjson

import "encoding/json"

// Game IDs read in by JSON Unmarshaling will have already been shortened.
type OsnGameID string

func (gameID OsnGameID) ShortID() string {
	return string(gameID)
}

func (gameID OsnGameID) MarshalJSON() ([]byte, error) {
	return json.Marshal(string(gameID))
}
