package queries

import "github.com/mono83/query"

// New constructs new queyr object with requested data
func New(columns []query.Named, schema query.Named, logic query.Logic, rules []query.Rule, conditions []query.Condition, sorting []query.Sorting, limit, offset int) query.Query {
	return impl{
		logic:      logic,
		rules:      rules,
		conditions: conditions,
		sorting:    sorting,
		offset:     offset,
		limit:      limit,
		schema:     schema,
		columns:    columns,
	}
}

type impl struct {
	logic      query.Logic
	rules      []query.Rule
	conditions []query.Condition
	sorting    []query.Sorting
	offset     int
	limit      int
	schema     query.Named
	columns    []query.Named
}

func (i impl) Type() query.Logic             { return i.logic }
func (i impl) Rules() []query.Rule           { return i.rules }
func (i impl) Conditions() []query.Condition { return i.conditions }
func (i impl) Sorting() []query.Sorting      { return i.sorting }
func (i impl) Limit() int                    { return i.limit }
func (i impl) Offset() int                   { return i.offset }
func (i impl) Schema() query.Named           { return i.schema }
func (i impl) Columns() []query.Named        { return i.columns }
