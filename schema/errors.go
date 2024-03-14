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
// github:kevindamm/wits-go/schema/errors.go

package schema

import "fmt"

//
// Sentinel values for errors when converting a unit or tile instance.
//

// A variant of InvalidUnitState for when having a known position.
type InvalidTileState struct {
	InvalidUnitState
	At HexCoordIndex
}

func (err InvalidTileState) Index() HexCoordIndex { return err.At }
func (err InvalidTileState) Error() string {
	return fmt.Sprintf("%s (%d)\n", err.string, err.unit)
}

func (err InvalidTileState) MoveTo(index HexCoordIndex) TileState {
	return err
}

// An invalid state includes a copy of the unit parameters used to create it.
type InvalidUnitState struct {
	string `json:"errmsg"`
	unit   UnitState
}

func (err InvalidUnitState) Error() string {
	return fmt.Sprintf("%s (%d)\n", err.string, err.unit)
}

// Information about the underlying unit can still be extracted.
func (err InvalidUnitState) Class() UnitClassEnum { return err.unit.Class() }
func (err InvalidUnitState) IsSpecial() bool      { return err.unit.IsSpecial() }
func (err InvalidUnitState) Race() UnitRaceEnum   { return err.unit.Race() }
func (err InvalidUnitState) Team() FriendlyEnum   { return err.unit.Team() }
func (err InvalidUnitState) Init() UnitState      { return err }

// Zero values for properties of an invalid state.
func (InvalidUnitState) Health() UnitHealth     { return 0 }
func (InvalidUnitState) Strength() UnitHealth   { return 0 }
func (InvalidUnitState) Distance() TileDistance { return 0 }
func (InvalidUnitState) Cost() ActionPoints     { return 0 }

// Appropriate zero values for these are where the unit is not allowed to do anything.
func (InvalidUnitState) IsAlternate() bool { return true }
func (InvalidUnitState) HasActed() bool    { return true }
func (InvalidUnitState) HasMoved() bool    { return true }
func (InvalidUnitState) HasAlted() bool    { return true }

func (err InvalidUnitState) HasParent() bool       { return false }
func (err InvalidUnitState) Parent() HexCoordIndex { return 0 }

// The effect-visiting methods of an invalid state are no-ops.
func (err InvalidUnitState) Toggle() UnitState                  { return err }
func (err InvalidUnitState) ReceiveBoost() UnitState            { return err }
func (err InvalidUnitState) ReceiveDamage(other Unit) UnitState { return err }
func (err InvalidUnitState) ReceiveCharm(other Unit) UnitState  { return err }

func (err InvalidUnitState) DoAction(action PlayerAction) (UnitState, error) {
	return err, fmt.Errorf("invalid unit %d cannot perform an action", err.unit)
}

func (err InvalidUnitState) ToTile(index HexCoordIndex) TileState {
	return InvalidTileState{err, index}
}
