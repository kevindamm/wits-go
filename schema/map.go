package schema

type GameMap interface {
	MapName() GameMapName
	GameID() GameMapID
	Terrain() map[HexCoordIndex]TileDefinition
	Units() []UnitInit
}

type GameMapName string
type GameMapID string

type TileDefinition interface {
	Positional
	Terrain() MapTerrain
}

type TileDistance uint8

type TileState interface {
	Placement
	UnitStateExtended
}

type MapTerrain byte

const (
	TERRAIN_UNKNOWN MapTerrain = iota
	TERRAIN_BONUS_NEUTRAL
	TERRAIN_BONUS_RED
	TERRAIN_BONUS_BLUE
	TERRAIN_FLOOR_1
	TERRAIN_FLOOR_2
	TERRAIN_FLOOR_3
	TERRAIN_FLOOR_4
	TERRAIN_WALL_1
	TERRAIN_WALL_2
	TERRAIN_WALL_3
	TERRAIN_WALL_4
	TERRAIN_BASE_RED
	TERRAIN_BASE_BLUE
	TERRAIN_SPAWN_RED
	TERRAIN_SPAWN_BLUE
)

func ParseTerrain(terrainEnum string) MapTerrain {
	return map[string]MapTerrain{
		"UNKNOWN":    TERRAIN_UNKNOWN,
		"BONUS":      TERRAIN_BONUS_NEUTRAL,
		"FLOOR_1":    TERRAIN_FLOOR_1,
		"FLOOR_2":    TERRAIN_FLOOR_2,
		"FLOOR_3":    TERRAIN_FLOOR_3,
		"FLOOR_4":    TERRAIN_FLOOR_4,
		"WALL_1":     TERRAIN_WALL_1,
		"WALL_2":     TERRAIN_WALL_2,
		"WALL_3":     TERRAIN_WALL_3,
		"WALL_4":     TERRAIN_WALL_4,
		"BASE_RED":   TERRAIN_BASE_RED,
		"BASE_BLUE":  TERRAIN_BASE_BLUE,
		"SPAWN_RED":  TERRAIN_SPAWN_RED,
		"SPAWN_BLUE": TERRAIN_SPAWN_BLUE,
	}[terrainEnum]
}
