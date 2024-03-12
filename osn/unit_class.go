package osn

import "github.com/kevindamm/wits-go/schema"

type UnitClassEnum byte

const (
	CLASS_UNKNOWN UnitClassEnum = iota
	CLASS_RUNNER
	CLASS_SOLDIER
	CLASS_MEDIC
	CLASS_SNIPER
	CLASS_HEAVY
	CLASS_SCRAMBLER
	CLASS_MOBI
	CLASS_BOMBSHELL
	CLASS_BRAMBLE
	CLASS_BRAMBLE_THORN
)

func (class UnitClassEnum) String() string {
	if class > CLASS_BRAMBLE_THORN || class < 0 {
		class = CLASS_UNKNOWN
	}
	return []string{
		"UNKNOWN_CLASS",
		"RUNNER",
		"SOLDIER",
		"MEDIC",
		"SNIPER",
		"HEAVY",
		"SCRAMBLER",
		"MOBI",
		"BOMBSHELL",
		"BRAMBLE",
		"BRAMBLE_THORN",
	}[int(class)]
}

func (class UnitClassEnum) AsWitsEnum() schema.UnitClass {
	if class > CLASS_BRAMBLE_THORN || class < 0 {
		class = CLASS_UNKNOWN
	}
	return []schema.UnitClass{
		schema.CLASS_UNKNOWN,
		schema.CLASS_RUNNER,
		schema.CLASS_SOLDIER,
		schema.CLASS_MEDIC,
		schema.CLASS_SNIPER,
		schema.CLASS_HEAVY,
		schema.CLASS_SPECIAL,
		schema.CLASS_SPECIAL,
		schema.CLASS_SPECIAL,
		schema.CLASS_SPECIAL,
		schema.CLASS_THORN,
	}[int(class)]
}
