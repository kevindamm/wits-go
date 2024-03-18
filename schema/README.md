# Abstract ontology for game state

Provides an interchangeable representation that both `witsjson` and `state`
modules use to serialize/deserialize or manipulate the game, without requiring
that one always import from the other (and of course avoiding import cycles).

The common hierarchy has a `GameReplay` responsible for the map (via terrain)
and the units currently positioned on the map (via a units list).  The replay
also identifies metadata such as the players involved, their rankings before
and after the match, and the sequence of turns (and composite actions) that
each player made during the replay.

## GameReplay (replay.go)

## Map and MapTerrain (map.go)

## Unit properties and capabilities (units.go)

## PlayerTurn and UnitAction (actions.go)

## HexCoord and HexCoordIndex (hex_coord.go)

## PlayerRole (player.go)
