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
	"fmt"
)

type OsnPlayerTurn struct {
	State   GameState `json:"state"`
	Actions []OsnPlayerAction
}

type OsnPlayerAction interface {
	Name() string
	AsDict() map[string]interface{}
}

func coerce[T OsnPlayerAction](encoded []byte) (data T, err error) {
	var output struct {
		Data T `json:"action"`
	}
	err = json.Unmarshal(encoded, &output)
	if err != nil {
		return data, err
	}
	return output.Data, nil
}

func ParseGenericAction(typename string, encoded []byte) (OsnPlayerAction, error) {
	switch typename {
	case "StartTurnAction":
		return coerce[StartTurnAction](encoded)
	case "SelectUnitAction":
		return coerce[SelectUnitAction](encoded)
	case "MoveUnitAction":
		return coerce[MoveUnitAction](encoded)
	case "ActiveHealAction":
		return coerce[ActiveHealAction](encoded)
	case "SelectSpawnTileAction":
		return coerce[SelectSpawnTileAction](encoded)
	case "EndTurnAction":
		return coerce[EndTurnAction](encoded)
	case "SpawnUnitAction":
		return coerce[SpawnUnitAction](encoded)
	case "RangeAttackAction":
		return coerce[RangeAttackAction](encoded)
	case "ScramblerSpellAction":
		return coerce[ScramblerSpellAction](encoded)
	case "RootBrambleAction":
		return coerce[RootBrambleAction](encoded)
	case "SpawnThornAction":
		return coerce[SpawnThornAction](encoded)
	case "RetractThornAction":
		return coerce[RetractThornAction](encoded)
	case "ToggleBombShellModeAction":
		return coerce[ToggleBombShellModeAction](encoded)
	case "ToggleAction":
		return coerce[ToggleAction](encoded)
	case "EatAction":
		return coerce[EatAction](encoded)
	case "SpitAction":
		return coerce[SpitAction](encoded)
	case "TeleportAction":
		return coerce[TeleportAction](encoded)
	}

	return nil, fmt.Errorf("Unrecognized typename " + typename)
}

type basePlayerAction struct {
	ActionName string `json:"name"`
}

func (actionBase basePlayerAction) Name() string {
	return actionBase.ActionName
}

func (actionBase basePlayerAction) String() string {
	return fmt.Sprintf("PlayerAction{\"%s\")", actionBase.Name())
}

type StartTurnAction struct {
	basePlayerAction
}

func (action StartTurnAction) String() string {
	return "StartTurnAction{}"
}

func (action StartTurnAction) AsDict() map[string]interface{} {
	return map[string]interface{}{
		"name": action.Name(),
	}
}

type EndTurnAction struct {
	basePlayerAction
}

func (action EndTurnAction) String() string {
	return "EndTurnAction{}"
}

func (action EndTurnAction) AsDict() map[string]interface{} {
	return map[string]interface{}{
		"name": action.Name(),
	}
}

type SelectUnitAction struct {
	basePlayerAction
	Index UnitIndex `json:"pawnID"`
}

func (action SelectUnitAction) String() string {
	return fmt.Sprintf("SelectUnitAction{#%d}", action.Index)
}

func (action SelectUnitAction) AsDict() map[string]interface{} {
	return map[string]interface{}{
		"name":  action.Name(),
		"index": action.Index,
	}
}

type MoveUnitAction struct {
	basePlayerAction
	Index   UnitIndex `json:"pawnID"`
	MovetoI int       `json:"desti"`
	MovetoJ int       `json:"destj"`

	From HexCoord `json:"from,omitempty"`
	To   HexCoord `json:"to,omitempty"`
}

func (action MoveUnitAction) String() string {
	return fmt.Sprintf("MoveUnitAction{#%d, (%d, %d)}",
		action.Index, action.MovetoI, action.MovetoJ)
}

func (action MoveUnitAction) AsDict() map[string]interface{} {
	if action.From == action.To {
		return map[string]interface{}{
			"name":   action.Name(),
			"pawnID": action.Index,
			"desti":  action.To.Column,
			"destj":  action.To.Row,
		}
	}

	return map[string]interface{}{
		"type": "MoveUnit",
		"from": action.From,
		"to":   action.To,
	}
}

type ActiveHealAction struct {
	basePlayerAction
	Index   UnitIndex `json:"pawnID"`
	TargetI int       `json:"desti"`
	TargetJ int       `json:"destj"`

	Caster HexCoord `json:"caster"`
}

func (action ActiveHealAction) String() string {
	return fmt.Sprintf("ActiveHealAction{#%d, (%d, %d)}",
		action.Index, action.TargetI, action.TargetJ)
}

func (action ActiveHealAction) AsDict() map[string]interface{} {
	return map[string]interface{}{
		"name":   action.Name(),
		"index":  action.Index,
		"target": []int{action.TargetI, action.TargetJ},
	}
}

type SelectSpawnTileAction struct {
	basePlayerAction
	SelectX int `json:"ix"`
	SelectY int `json:"iy"`
}

func (action SelectSpawnTileAction) String() string {
	return fmt.Sprintf("SelectSpawnTileAction{(%d, %d)}",
		action.SelectX, action.SelectY)
}

func (action SelectSpawnTileAction) AsDict() map[string]interface{} {
	return map[string]interface{}{
		"name":     action.Name(),
		"position": []int{action.SelectX, action.SelectY},
	}
}

type SpawnUnitAction struct {
	basePlayerAction
	Position []int         `json:"position,omitempty"`
	Color    PlayerColor   `json:"color"`
	Class    UnitClassEnum `json:"role"`
}

func (action SpawnUnitAction) String() string {
	return fmt.Sprintf("SpawnUnitAction{%s, %s}",
		action.Color, action.Class)
}

func (action SpawnUnitAction) AsDict() map[string]interface{} {
	actionDict := map[string]interface{}{
		"name":  action.Name(),
		"color": action.Color,
		"class": action.Class,
	}
	if len(action.Position) > 0 {
		actionDict["position"] = action.Position
	}
	return actionDict
}

type RangeAttackAction struct {
	basePlayerAction
	Index   UnitIndex `json:"pawnID"`
	TargetI int       `json:"desti"`
	TargetJ int       `json:"destj"`
}

func (action RangeAttackAction) String() string {
	return fmt.Sprintf("RangeAttackAction{#%d, (%d, %d)}",
		action.Index, action.TargetI, action.TargetJ)
}

func (action RangeAttackAction) AsDict() map[string]interface{} {
	return map[string]interface{}{
		"name":   action.Name(),
		"index":  action.Index,
		"target": []int{action.TargetI, action.TargetJ},
	}
}

type ScramblerSpellAction struct {
	basePlayerAction
	Index   UnitIndex `json:"pawnID"`
	TargetI int       `json:"desti"`
	TargetJ int       `json:"destj"`
}

func (action ScramblerSpellAction) String() string {
	return fmt.Sprintf("ScramblerSpellAction{#%d, (%d, %d)}",
		action.Index, action.TargetI, action.TargetJ)
}

func (action ScramblerSpellAction) AsDict() map[string]interface{} {
	return map[string]interface{}{
		"name":     action.Name(),
		"index":    action.Index,
		"position": []int{action.TargetI, action.TargetJ},
	}
}

type SpawnThornAction struct {
	basePlayerAction
	Index   UnitIndex `json:"pawnID"`
	TargetI int       `json:"desti"`
	TargetJ int       `json:"destj"`
}

func (action SpawnThornAction) String() string {
	return fmt.Sprintf("SpawnThornAction{#%d, (%d, %d)}",
		action.Index, action.TargetI, action.TargetJ)
}

func (action SpawnThornAction) AsDict() map[string]interface{} {
	return map[string]interface{}{
		"name":     action.Name(),
		"index":    action.Index,
		"position": []int{action.TargetI, action.TargetJ},
	}
}

type RetractThornAction struct {
	basePlayerAction
	Index UnitIndex `json:"pawnID"`
	UnitI int       `json:"desti"`
	UnitJ int       `json:"destj"`
}

func (action RetractThornAction) String() string {
	return fmt.Sprintf("RetractThornAction{#%d, (%d, %d)}",
		action.Index, action.UnitI, action.UnitJ)
}

func (action RetractThornAction) AsDict() map[string]interface{} {
	return map[string]interface{}{
		"name":     action.Name(),
		"index":    action.Index,
		"position": []int{action.UnitI, action.UnitJ},
	}
}

type ToggleBombShellModeAction struct {
	basePlayerAction
	Index UnitIndex `json:"pawnID"`
	UnitI int       `json:"desti"`
	UnitJ int       `json:"destj"`
}

func (action ToggleBombShellModeAction) String() string {
	return fmt.Sprintf("ToggleArtilleryModeAction{#%d, (%d, %d)}",
		action.Index, action.UnitI, action.UnitJ)
}

func (action ToggleBombShellModeAction) AsDict() map[string]interface{} {
	return map[string]interface{}{
		"name":     "ToggleBombShellModeAction",
		"index":    action.Index,
		"position": []int{action.UnitI, action.UnitJ},
	}
}

// Buries (or unburies) the Bramble into the ground, needed before spawning brambles.
type RootBrambleAction struct {
	basePlayerAction
	Index   UnitIndex `json:"pawnID"`
	TargetI int       `json:"desti"`
	TargetJ int       `json:"destj"`
}

func (action RootBrambleAction) String() string {
	return fmt.Sprintf("RootBrambleAction{#%d, (%d, %d)}",
		action.Index, action.TargetI, action.TargetJ)
}

func (action RootBrambleAction) AsDict() map[string]interface{} {
	return map[string]interface{}{
		"name":     "RootBrambleAction",
		"index":    action.Index,
		"position": []int{action.TargetI, action.TargetJ},
	}
}

// Simplified, combined action for RootBramble and ToggleArtillery actions
type ToggleAction struct {
	basePlayerAction
	Index    UnitIndex `json:"index"`
	Position []int     `json:"position"`
}

func (action ToggleAction) String() string {
	return fmt.Sprintf("ToggleAction{#%d, (%d, %d)}",
		action.Index, action.Position[0], action.Position[1])
}

func (action ToggleAction) AsDict() map[string]interface{} {
	return map[string]interface{}{
		"name":     "ToggleAction",
		"index":    action.Index,
		"position": action.Position,
	}
}

// The first half of the teleport action, selecting a unit to be teleported.
type EatAction struct {
	basePlayerAction
	Index UnitIndex `json:"pawnID"`
	UnitI int       `json:"desti"`
	UnitJ int       `json:"destj"`
}

func (action EatAction) String() string {
	return fmt.Sprintf("EatAction{#%d, (%d, %d)}",
		action.Index, action.UnitI, action.UnitJ)
}

func (action EatAction) AsDict() map[string]interface{} {
	return map[string]interface{}{
		"name":  action.Name(),
		"index": action.Index,
		"from":  []int{action.UnitI, action.UnitJ},
	}
}

// The second half of the teleport action,
// indicating the location for teleporting the unit to.

type SpitAction struct {
	basePlayerAction
	Index   UnitIndex `json:"pawnID"`
	TargetI int       `json:"desti"`
	TargetJ int       `json:"destj"`
}

func (action SpitAction) String() string {
	return fmt.Sprintf("SpitAction{#%d, (%d, %d)}",
		action.Index, action.TargetI, action.TargetJ)
}

func (action SpitAction) AsDict() map[string]interface{} {
	return map[string]interface{}{
		"name":  action.Name(),
		"index": action.Index,
		"to":    []int{action.TargetI, action.TargetJ},
	}
}

// The combination of EatAction+SpitAction.
type TeleportAction struct {
	basePlayerAction
	Caster HexCoord `json:"caster"`
	From   HexCoord `json:"from"`
	To     HexCoord `json:"to"`
}

func (action TeleportAction) String() string {
	return fmt.Sprintf("TeleportAction{(%d, %d): (%d, %d)->(%d, %d)}",
		action.Caster.Column, action.Caster.Row,
		action.From.Column, action.From.Row,
		action.To.Column, action.To.Row)
}

func (action TeleportAction) AsDict() map[string]interface{} {
	return map[string]interface{}{
		"name": action.ActionName,
		"from": action.From,
		"to":   action.To,
	}
}
