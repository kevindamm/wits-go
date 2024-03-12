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
// github:kevindamm/wits-go/osn/boolish.go

package osn

import "encoding/json"

// Able to unmarshal from both int and bool representations,
type Boolish struct {
	Value bool
}

// Decoding from JSON format tries both boolean and integer representations.
func (b *Boolish) UnmarshalJSON(encoded []byte) error {
	var boolVal bool
	if err := json.Unmarshal(encoded, &boolVal); err != nil {
		b.Value = boolVal
	} else {
		var intVal int
		if err := json.Unmarshal(encoded, &boolVal); err != nil {
			b.Value = (intVal != 0)
		} else {
			return err
		}
	}
	return nil
}

// Always marshals as a bool value.
func (b Boolish) MarshalJSON() ([]byte, error) {
	return json.Marshal(b.Value)
}
