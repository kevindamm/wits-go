package schema

type GameReplay interface {
	GameID() OsnGameID
	Map() GameMap
	MapTheme() string
	Players() []PlayerRole

	InitState() GameInit
	MatchReplay() []PlayerTurn
	MatchResult() TerminalStatus
}

type GameState interface {
	BaseHP(player FriendlyEnum) BaseHealth
	BonusWits() []HexCoord
	Units() []UnitPlacement
}

type GameInit interface {
	Units() []UnitInit
}

// Outcome enum relative to the player (player 1 if without context)
// Values include win (destruction), win (extinction), ...lose, forfeit
// These
type TerminalStatus byte

const (
	STATUS_UNKNOWN TerminalStatus = iota
	VICTORY_DESTRUCTION
	VICTORY_EXTINCTION
	VICTORY_RESIGNATION
	LOSS_DELAY_OF_GAME
	LOSS_DESTRUCTION
	LOSS_EXTINCTION
	LOSS_RESIGNATION
)

type BaseHealth byte
