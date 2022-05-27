package rules

import (
	"github.com/mono83/query"
	"github.com/mono83/query/match"
)

var False query.Rule = falseRule{}

// False is special rule instance, that is always false
type falseRule struct{}

// GetLeft is query.Rule interface implementation
func (falseRule) Left() interface{} { return 1 }

// GetRight is query.Rule interface implementation
func (falseRule) Right() interface{} { return 2 }

// GetType is query.Rule interface implementation
func (falseRule) Type() match.Type { return match.Eq }
