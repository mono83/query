package rules

import "github.com/mono83/query"

// WithRight constructs new rule with requested right value.
func WithRight(r query.Rule, v interface{}) query.Rule {
	if r == nil || r.Right() == v {
		return r
	}

	return New(r.Left(), r.Type(), v)
}

// WithLeft constructs new rule with requested left value.
func WithLeft(r query.Rule, v interface{}) query.Rule {
	if r == nil || r.Left() == v {
		return r
	}

	return New(v, r.Type(), r.Right())
}
