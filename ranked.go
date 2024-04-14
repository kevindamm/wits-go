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
// github:kevindamm/wits-go/ranked.go

package wits

// The league & rank of each player at the beginning of the match.
type PlayerStandings interface {
	Tier() LeagueTier
	Rank() LeagueRank
}

// Ranked competitions are divided into skill levels, advancing within a
// local (top-100) ranking leads to advancing to the next higher tier.
type LeagueTier string

// This is the rank out of a group of 100 players (an integer between 1..100).
type LeagueRank int

const (
	LEAGUE_TIER_UNRANKED     LeagueTier = "Unranked"
	LEAGUE_TIER_NOVICE       LeagueTier = "Novice"
	LEAGUE_TIER_INTERMEDIATE LeagueTier = "Intermediate"
	LEAGUE_TIER_ADVANCED     LeagueTier = "Advanced"
	LEAGUE_TIER_EXPERT       LeagueTier = "Expert"
)

// The league & rank of each player as a result of the match outcome.
type PlayerStandingsUpdate interface {
	PlayerStandings
	Delta() int
}
