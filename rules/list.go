package rules

import "github.com/mono83/query"

// List is a handy alias to create slice of rules
func List(r ...query.Rule) []query.Rule {
	return r
}
