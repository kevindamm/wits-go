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
// github:kevindamm/wits-go/schema/game_id.go

package schema

// String alias for game_id values from OSN, with normalization of identity.
// When generalizing to schema representation, only provide the shortened ID.
type OsnGameID interface {
	// Returns a shortened representation (trimming the zero-entropy prefix).
	ShortID() string
}
