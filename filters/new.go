package filters

import "github.com/mono83/query"

// New constructs new filter object with requested data
func New(logic query.Logic, rules []query.Rule, conditions []query.Condition, sorting []query.Sorting, limit, offset int) query.Filter {
	return filter{
		logic:      logic,
		rules:      rules,
		conditions: conditions,
		sorting:    sorting,
		limit:      limit,
		offset:     offset,
	}
}

type filter struct {
	logic      query.Logic
	rules      []query.Rule
	conditions []query.Condition
	sorting    []query.Sorting
	offset     int
	limit      int
}

func (f filter) Type() query.Logic             { return f.logic }
func (f filter) Rules() []query.Rule           { return f.rules }
func (f filter) Conditions() []query.Condition { return f.conditions }
func (f filter) Sorting() []query.Sorting      { return f.sorting }
func (f filter) Offset() int                   { return f.offset }
func (f filter) Limit() int                    { return f.limit }
