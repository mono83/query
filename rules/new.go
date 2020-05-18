package rules

import (
	"github.com/mono83/query"
	"github.com/mono83/query/match"
)

// New builds new rule
func New(left interface{}, op match.Type, right interface{}) query.Rule {
	return query.CommonRule{Left: left, Right: right, Type: op}
}
