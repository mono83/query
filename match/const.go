package match

import (
	"strconv"
)

// List of supported rule operations
const (
	Unknown Type = 0

	IsNil     Type = 1 // Alias
	IsNull    Type = 1
	NotIsNil  Type = 2 // Alias
	NotIsNull Type = 2
	IsNotNull Type = 2 // Alias
	IsNotNil  Type = 2 // Alias

	Eq        Type = 3 // Alias
	Equals    Type = 3
	Neq       Type = 4 // Alias
	NotEquals Type = 4

	In    Type = 5
	NotIn Type = 6

	Gt               Type = 7 // Alias
	GreaterThan      Type = 7
	Lte              Type = 8 // Alias
	LesserThanEquals Type = 8 // Alias
	LowerThanEquals  Type = 8

	Gte               Type = 9 // Alias
	GreaterThanEquals Type = 9
	Lt                Type = 10 // Alias
	LesserThan        Type = 10 // Alias
	LowerThan         Type = 10

	Contains Type = 11
	Like     Type = 11 // Alias
)

func (t Type) String() string {
	if t.IsCustom() {
		return "Custom #" + strconv.Itoa(int(t))
	}
	if t == Unknown {
		return "Unknown"
	}

	d, ok := Definitions[t]
	if !ok {
		return "Unsupported #" + strconv.Itoa(int(t))
	}

	return d.Name
}

// Names returns full list of string names (including math symbols), that can be
// used to identify rule type
func (t Type) Names() []string {
	def, ok := Definitions[t]
	if !ok {
		return nil
	}

	return def.Names()
}

// These constants describes top and bottom boundaries for registered
// matcher type constants
var (
	lower byte = 1
	upper byte = 11
)

// Count return total count of match operators
func Count() int { return int(upper - lower + 1) }

// All returns full list of rule operations, except Unknown
// Used in tests primarily
func All() []Type {
	var all []Type
	for i := lower; i <= upper; i++ {
		all = append(all, Type(i))
	}

	return all
}
