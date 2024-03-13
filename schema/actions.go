package schema

type PlayerTurn interface {
	TurnCount() uint
	Actions() []PlayerAction

	// Temporarily here so that we can validate the simulation against the intermediate states.
	State() GameState
}

type PlayerAction interface {
	ActionName() string
	RelVarEncoding() string
	Visit(*GameState) error
}

// Non-negative integer, the amount of "wits" (action points) available, or cost.
type ActionPoints byte

//
// UNKNOWN ACTION
//

type UnknownActionError struct {
	Name string
}

func (e UnknownActionError) Error() string { return "unknown action name " + e.Name }

//
// PASS (retain remaining wits)
//

// While not absolutely necessary to have the encoding work out, as the player
// turns have dedicated lists they are in, this "no-op" action is crucial.  It
// allows for simplifying assumptions later about empty sets/subsets vis-a-vis
// the recurrence relation of reordering action subgroups (when canonicalizing).
type PassAction struct{}

func (PassAction) ActionName() string     { return "Pass" }
func (PassAction) RelVarEncoding() string { return `["pass"]` }
func (PassAction) Visit(*GameState) error { return nil }
