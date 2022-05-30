package conditions

import "github.com/mono83/query"

// IsEmpty returns true if given condition nil or contains
// no rules or nested conditions
func IsEmpty(c query.Condition) bool {
	return c == nil || (len(c.Rules()) == 0 && len(c.Conditions()) == 0)
}
