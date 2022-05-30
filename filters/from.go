package filters

import "github.com/mono83/query"

// FromCondition constructs new filter from condition with
// appended sorting, limit and offset.
func FromCondition(c query.Condition, s []query.Sorting, limit, offset int) query.Filter {
	if c == nil {
		return nil
	}

	return New(c.Type(), c.Rules(), c.Conditions(), s, limit, offset)
}
