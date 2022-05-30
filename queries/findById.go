package queries

import (
	"github.com/mono83/query"
	"github.com/mono83/query/names"
	"github.com/mono83/query/rules"
)

// FindByID64Simple returns query for matching entries by IDs
func FindByID64Simple(table string, id ...int64) query.Query {
	return id64SimpleQuery{
		table: table,
		ids:   id,
	}
}

type id64SimpleQuery struct {
	table string
	ids   []int64
}

func (id64SimpleQuery) Type() query.Logic             { return query.And }
func (id64SimpleQuery) Columns() []query.Named        { return nil }
func (id64SimpleQuery) Conditions() []query.Condition { return nil }
func (id64SimpleQuery) Sorting() []query.Sorting      { return nil }
func (id64SimpleQuery) Limit() int                    { return 0 }
func (id64SimpleQuery) Offset() int                   { return 0 }
func (i id64SimpleQuery) Schema() query.Named         { return names.String(i.table) }
func (i id64SimpleQuery) Rules() []query.Rule         { return []query.Rule{rules.MatchID64(i.ids...)} }
