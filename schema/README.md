# Abstract ontology for game state

Provides an interchangeable representation that both `witsjson` and `state`
modules use to serialize/deserialize or manipulate the game, without requiring
that one always import from the other (and of course avoiding import cycles).

The common hierarchy has a `GameReplay` responsible for the map (via terrain)
and the units currently positioned on the map (via a units list).  The replay
also identifies metadata such as the players involved, their rankings before
and after the match, and the sequence of turns (and composite actions) that
each player made during the replay.

The interfaces defined here add some function call and interface-ptr overhead
but this overhead is only present when the static type is explicitly calling
for it.  Much of the `state` module does its processing directly on the bytes
representation, whereas the `witsjson` functions deal primarily with the types
from interfaces in `schema`.

## GameReplay (replay.go)

 * `GameReplay` interface represents all details of a single match.  Contains:

   * identifiers (game ID and map ID)

   * player details (see below) including changes in standings for league games

   * initial conditions (unit locations and other variable state)

   * the replay as a series of turns, each a sequence of actions

   * the result (win/loss) of the match, from the perspective of player_1

All atoms (and near-atomic values like strings) are given dedicated types.  This
isn't strictly necessary but it has the added benefit of reducing mismatches (for
example between a map's name and its ID).  In some cases with JSON unmarshaling
it becomes quite useful to have a type alias for custom transformations.

## Map and MapTerrain (`map.go`)

The map definition is unlike every other property listed here in that it is not
included in the replay representation directly.  The terrain and initial state
are the same for every replay on the same map so these details are held in a
separate file and given a separate structured representation.

 * `MapDescription` is the top-level type, representing all map-related details

   * the map's identifier (for locating the definition) and its name

   * terrain definition -- coordinates and terrain type (for floor, walls and other)

   * initial state (unit placement, can be extended for puzzle-like arbitrary starting conditions)

There is a property, one that will soon be deprecated, that indicates whether
the coordinate system is the legacy (column, row) layout or the more fluent
(i, j) base-vector layout centered on the base location.  The goal is to
transition to the base-vector coordinate system and un-mirror/un-rotate to
normalize team positions as part of the conversion from OSN representation.

## Unit properties and capabilities (`units.go`)

The schema defines both a behavioral representation (`Unit` and `UnitState`)
and a descriptive representation (`UnitInit`).  This is done to separate the
properties typically needed in a map- or game- initialization (which need to
point to the unit's position), from the game-state representation (which has
to enforce exclusivity on location).  Perhaps these could be unified but the
JSON serialization becomes noticeably more convoluted or the coordinte state
becomes unnecessarily repeated between the unit representation and the in-game
map representation.

Enum-style types are also defined here for types that are value types across
the modules.  This inclues the unit's class (abilities), their race (special
type), and their team (FriendlyEnum).

## PlayerTurn and UnitAction (`actions.go`)

The number and complexity of different actions justifies representing them
within a separate file.  The details within the `schema` module are short
relative to the details needed in `witsjson` and `state`, primarily just
setting up the double-dispatch visitor pattern for `Unit -> Action (State)`
application as used by the state, while maintaining the type specification
for witsjson to extract the resulting game's state for serialization.

## HexCoord and HexCoordIndex (`hex_coord.go`)

There is a common type for the coordinate positions as a two-dimensional value.
It has a corresponding JSON type for marshaling and unmarshaling it, and that
uses a simple 2D array rather than expanding it into the equivalent JS-object
structure.  The `state` module does not deal in position values, typically
prefering to use an index (HexCoordIndex, a small unsigned value) that compacts
any coordinate system thoroughly, regardless of the map's shape or topology.

HexCoord is designed to be an immutable value type.  Any transformation on its
value should produce a new value.  The finiteness of coordinates makes it easy
to extend this to singular references of coordinate positions, even without a
concrete implementation.  Much of the search and simulation is done without an
explicit coordinate value, instead using indexes and subgraph representations.

## PlayerRole and PlayerHistory (`player.go`)

There is a representation for each player, `PlayerRole`, including their name
and choice of race for the match, which team/role they correspond to and their
standings (tier, rank) before and after the match.  Any game-specific state
(e.g., base HP and amount of remaining wits) is held in the replay's `init`
properties or inferred from the playout of the players' actions.

There is also a `PlayerHistory` representation which collects all the games
on record for the player and some statistics about their previous league games.