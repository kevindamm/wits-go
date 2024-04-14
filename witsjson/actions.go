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
// github:kevindamm/wits-go/witsjson/actions.go

package witsjson

import (
	"encoding/json"
	"fmt"

	"github.com/kevindamm/wits-go"
)

type PlayerTurnJSON struct {
	Turn_    uint                `json:"turn"`
	Actions_ []wits.PlayerAction `json:"actions"`

	// Temporarily here so that we can validate the simulation against the intermediate states.
	State_ wits.GameState `json:"state"`
}

func (turn PlayerTurnJSON) TurnCount() uint {
	return turn.Turn_
}

// On odd turns it is team 1's turn, on even turns it is team 2's turn.
func (turn PlayerTurnJSON) Team() wits.FriendlyEnum {
	// The subtraction from 2 makes the mod give {1, 2} values instead of {0, 1}.
	return 2 - wits.FriendlyEnum(turn.Turn_%2)
}

// Works for both 1v1 and multiplayer.
func (turn PlayerTurnJSON) Opponent() wits.FriendlyEnum {
	if turn.Team() == wits.FR_ENEMY {
		return wits.FR_SELF
	} else {
		return wits.FR_ENEMY
	}
}

// Give the list of actions performed for the current turn.
func (turn PlayerTurnJSON) Actions() []wits.PlayerAction {
	return turn.Actions_
}

// DEPRECATED: included for sanity checks over existing replays until proper testing is in place.
func (turn PlayerTurnJSON) State() wits.GameState {
	return turn.State_
}

type playerActionJSON struct {
	Name_ ActionNameJSON `json:"name"`
}

type ActionNameJSON string

const (
	PASS_PLAY     ActionNameJSON = "Pass"
	MOVE_UNIT     ActionNameJSON = "MoveUnit"
	HEAL_UNIT     ActionNameJSON = "HealUnit"
	SPAWN_UNIT    ActionNameJSON = "SpawnUnit"
	ATTACK        ActionNameJSON = "Attack"
	CHARM_UNIT    ActionNameJSON = "CharmUnit"
	TOGGLE_ALT    ActionNameJSON = "ToggleAlt"
	TELEPORT_UNIT ActionNameJSON = "Teleport"
)

func ParseGenericAction(typename string, encoded []byte) (wits.PlayerAction, error) {
	if len(typename) == 0 {
		return nil, wits.UnknownActionError{Name: `""`}
	}
	switch ActionNameJSON(typename) {
	case MOVE_UNIT:
		// Most units can move, except toggled artillery, bramble & its thorns.
		return coerce[MoveUnitAction](encoded)
	case HEAL_UNIT:
		// The "medic" special action.
		return coerce[HealUnitAction](encoded)
	case SPAWN_UNIT:
		// Includes SpawnThorn action.
		return coerce[SpawnUnitAction](encoded)
	case ATTACK:
		// Includes soldier, heavy, sniper, scout and bramble-thorn attacks.
		return coerce[AttackAction](encoded)
	case CHARM_UNIT:
		// The "scrambler" special action.
		return coerce[CharmUnitAction](encoded)
	case TOGGLE_ALT:
		// Includes retract thorn (which affects it and all its children).
		return coerce[ToggleAltAction](encoded)
	case TELEPORT_UNIT:
		// The "mobi" special action.
		return coerce[TeleportUnitAction](encoded)

	case PASS_PLAY:
		return wits.PassAction{}, nil
	}
	return nil, wits.UnknownActionError{Name: typename}
}

// Generic JSON-decoding routine for all PlayerAction implementing types.
func coerce[T wits.PlayerAction](encoded []byte) (T, error) {
	var output struct {
		Data T `json:"action"`
	}
	if err := json.Unmarshal(encoded, &output); err != nil {
		return output.Data, err
	}
	return output.Data, nil
}

//                            //
// ACTIONS /// CONCRETE TYPES //
//                            //

// Moves a unit from a HexCoord position to a (different) HexCoord position.
type MoveUnitAction struct {
	playerActionJSON
	From wits.HexCoord `json:"from"`
	To   wits.HexCoord `json:"to"`
}

func (action MoveUnitAction) ActionName() string { return string(MOVE_UNIT) }

func (action MoveUnitAction) RelVarEncoding() string {
	return fmt.Sprintf(`["move", ["ij", %d, %d], ["ij", %d, %d]]`,
		action.From.I(), action.From.J(), action.To.I(), action.To.J())
}

func (action MoveUnitAction) Visit(state *wits.GameState) error {
	// Locate the unit for this action and move it to the indicated position.
	// TODO
	return nil
}

// Heals a friendly unit to their initial HP + 1.
type HealUnitAction struct {
	playerActionJSON
	Healer wits.HexCoord `json:"healer"`
	Target wits.HexCoord `json:"target"`
}

func (action HealUnitAction) ActionName() string { return string(HEAL_UNIT) }

func (action HealUnitAction) RelVarEncoding() string {
	return fmt.Sprintf(`["heal", ["ij", %d, %d], ["ij", %d, %d]]`,
		action.Healer.I(), action.Healer.J(), action.Target.I(), action.Target.J())
}

func (action HealUnitAction) Visit(state *wits.GameState) error {
	// Locate the target unit and update its health to class.hp+1
	// TODO
	return nil
}

// Units may be spawned only from specific locations on the map.
type SpawnUnitAction struct {
	playerActionJSON
	Spawn wits.HexCoord `json:"spawn"`
	Class UnitClassJSON `json:"class"`
}

func (action SpawnUnitAction) ActionName() string { return string(SPAWN_UNIT) }

func (action SpawnUnitAction) RelVarEncoding() string {
	return fmt.Sprintf(`["spawn", ["ij", %d, %d], %s]`,
		action.Spawn.I(), action.Spawn.J(), string(action.Class))
}

// Parentage is determined by this action but the parent/spawned-from state is
// not stored as part of the unit -- it is instead determined at the game state
// (via turn reconstruction and a specialized game state).  This allows for a
// unified implementation here, whether it was a SpawnTile or a Bramble.
func (action SpawnUnitAction) Visit(state *wits.GameState) error {
	// TODO
	return nil
}

// Some units may attack other units, and their attack strength is dependent
// on the UnitClass of the attacking unit, the details of which are in the game
// state.  The action itself only needs to mention the attacker's location and
// the location of the unit's target (only units may attack).
type AttackAction struct {
	Agent  wits.HexCoord `json:"agent"`
	Target wits.HexCoord `json:"target"`
}

func (action AttackAction) ActionName() string { return string(ATTACK) }

func (action AttackAction) RelVarEncoding() string {
	return fmt.Sprintf(`["pow", ["ij", %d, %d], ["ij", %d, %d]]`,
		action.Agent.I(), action.Agent.J(), action.Target.I(), action.Target.J())
}

func (action AttackAction) Visit(state *wits.GameState) error {
	// TODO
	return nil
}

// This is a special action for the Scrambler unit class.  It converts the unit
// of an opposing team onto the player's team.
type CharmUnitAction struct {
	Agent  wits.HexCoord `json:"agent"`
	Target wits.HexCoord `json:"target"`
}

func (action CharmUnitAction) ActionName() string { return string(CHARM_UNIT) }

func (action CharmUnitAction) RelVarEncoding() string {
	return fmt.Sprintf(`["charm", ["ij", %d, %d], ["ij", %d, %d]]`,
		action.Agent.I(), action.Agent.J(), action.Target.I(), action.Target.J())
}

func (action CharmUnitAction) Visit(state *wits.GameState) error {
	// TODO
	return nil
}

type ToggleAltAction struct {
	wits.HexCoord `json:"position"`
}

func (action ToggleAltAction) ActionName() string { return string(TOGGLE_ALT) }

func (action ToggleAltAction) RelVarEncoding() string {
	return fmt.Sprintf(`["toggle", ["ij", %d, %d]]`, action.I(), action.J())
}

func (action ToggleAltAction) Visit(state *wits.GameState) error {
	// TODO
	return nil
}

type TeleportUnitAction struct {
	wits.HexCoord `json:"mobi"`
	From          wits.HexCoord `json:"from"`
	To            wits.HexCoord `json:"to"`
}

func (action TeleportUnitAction) ActionName() string { return string(TELEPORT_UNIT) }

func (action TeleportUnitAction) RelVarEncoding() string {
	return fmt.Sprintf(`["port", ["ij", %d, %d], ["ij", %d, %d], ["ij", %d, %d]]`,
		action.I(), action.J(),
		action.From.I(), action.From.J(),
		action.To.I(), action.To.J())
}

func (action TeleportUnitAction) Visit(state *wits.GameState) error {
	// TODO
	return nil
}
