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
// github:kevindamm/wits-go/schema/status.go

package schema

// An enum describing the win or loss result, relative to the current player
// (relative to player 1 if without context).
// Values include win (destruction), win (extinction), ...lose, forfeit
type TerminalStatus byte

const (
	STATUS_UNKNOWN TerminalStatus = iota
	VICTORY_DESTRUCTION
	VICTORY_EXTINCTION
	VICTORY_RESIGNATION
	DELAY_OF_GAME
	LOSS_DESTRUCTION
	LOSS_EXTINCTION
	LOSS_RESIGNATION
)

// Differentiating bit for non-unknown/defaulting game results.
const WINLOSS_BIT TerminalStatus = 0b0100

// Converts win or loss status to the equivalent state from the opponent's view.
// An UNKNOWN is understandably still UNKNOWN, and may represent a game not yet
// completed.
//
// Although a DELAY_OF_GAME is considered equivalent to resignation in actual
// play, there is a special exception made here to give it a similar unitary
// representation similar to unknown.  This is useful in situations where reward
// might erroneously be propagated back from a timing issue when it may not have
// been a skill-related issue.
func (status TerminalStatus) Opposing() TerminalStatus {
	if status == STATUS_UNKNOWN || status == DELAY_OF_GAME {
		return status
	}
	return (status ^ WINLOSS_BIT)
}
