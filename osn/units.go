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
// github:kevindamm/wits-go/osn/units.go

package osn

import (
	"github.com/kevindamm/wits-go/witsjson"
)

type UnitStatus struct { // also see schema.UnitPlacement, schema.UnitTurnState
	Index UnitIndex         `json:"identifier"`
	Owner PlayerIndex       `json:"owner"`
	Team  TeamIndex         `json:"team"`
	Color PlayerColor       `json:"color"`
	Class UnitClassEnum     `json:"class"`
	Race  witsjson.UnitRace `json:"race"`

	HexCoord
	Attacked    Boolish `json:"hasAttacked"`
	Moved       Boolish `json:"hasMoved"`
	Transformed Boolish `json:"hasTransformed"`

	Health    UnitHealth `json:"health"`
	AltStatus Boolish    `json:"isAlt"`
	HealthAlt UnitHealth `json:"altHealth"`

	Parent      UnitIndex `json:"parent"`
	SpawnedFrom UnitIndex `json:"spawnedFrom"`
}

type UnitIndex int

type AlternateForm struct {
	Enabled Boolish    `json:"enabled"`
	Health  UnitHealth `json:"health"`
}

type Possession struct {
	Parent UnitIndex `json:"parent"`
	From   UnitIndex `json:"from"`
}

type UnitHealth int
