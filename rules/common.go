package rules

import (
	"github.com/mono83/query"
	"github.com/mono83/query/match"
)

// IsNull returns rule with IS NULL matcher for provided field
func IsNull(field interface{}) query.Rule {
	return New(field, match.IsNull, nil)
}

// IsNotNull returns rule with IS NOT NULL matcher for provided field
func IsNotNull(field interface{}) query.Rule {
	return New(field, match.IsNotNull, nil)
}

// Eq returns rule built with EQUALS matcher
func Eq(left, right interface{}) query.Rule {
	return New(left, match.Equals, right)
}

// MatchID64 returns rule for matching IDs
func MatchID64(id ...int64) query.Rule {
	switch len(id) {
	case 0:
		return False{}
	case 1:
		return Eq("id", id[0])
	default:
		return New("id", match.In, id)
	}
}
