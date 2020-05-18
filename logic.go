package query

// Logic describes inner relations inside conditions
// Can be AND or OR
type Logic byte

// List of defined condition relations
const (
	None Logic = 0
	And  Logic = 1
	Or   Logic = 2
)

// String returns string representation of logic operator
func (l Logic) String() string {
	switch l {
	case And:
		return "AND"
	case Or:
		return "OR"
	default:
		return ""
	}
}
