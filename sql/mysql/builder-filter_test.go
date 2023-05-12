package mysql

import (
	"fmt"
	"testing"

	"github.com/mono83/query"
	"github.com/mono83/query/filters"
	"github.com/mono83/query/names"
	"github.com/mono83/query/rules"
	"github.com/mono83/query/sorting"
	"github.com/stretchr/testify/assert"
)

var filterToSQLDataProvider = []struct {
	SQL    string
	Filter query.Filter
}{
	// No sorting
	{"`id` NOT IS NULL", filters.New(query.And, []query.Rule{rules.IsNotNull(names.String("id"))}, nil, nil, 0, 0)},
	{"`id` NOT IS NULL LIMIT 5", filters.New(query.And, []query.Rule{rules.IsNotNull(names.String("id"))}, nil, nil, 5, 0)},
	{"`id` NOT IS NULL", filters.New(query.And, []query.Rule{rules.IsNotNull(names.String("id"))}, nil, nil, 0, 5)},
	{"`id` NOT IS NULL LIMIT 5,2", filters.New(query.And, []query.Rule{rules.IsNotNull(names.String("id"))}, nil, nil, 2, 5)},
	{" LIMIT 5", filters.New(query.And, nil, nil, nil, 5, 0)},
	{" ORDER BY `id` DESC LIMIT 5", filters.New(query.And, nil, nil, []query.Sorting{sorting.Desc("id")}, 5, 0)},
	{" ORDER BY `id` DESC", filters.New(query.And, nil, nil, []query.Sorting{sorting.Desc("id")}, 0, 0)},

	// With sorting
	{
		"`id` NOT IS NULL ORDER BY `id` ASC",
		filters.New(
			query.And,
			[]query.Rule{rules.IsNotNull(names.String("id"))},
			nil,
			[]query.Sorting{sorting.Asc("id")},
			0,
			0,
		),
	},
	{
		"`id` NOT IS NULL ORDER BY `id` ASC,`name` DESC",
		filters.New(
			query.And,
			[]query.Rule{rules.IsNotNull(names.String("id"))},
			nil,
			[]query.Sorting{sorting.Asc("id"), sorting.Desc("name")},
			0,
			0,
		),
	},
	{
		"`id` NOT IS NULL ORDER BY `id` ASC LIMIT 3",
		filters.New(
			query.And,
			[]query.Rule{rules.IsNotNull(names.String("id"))},
			nil,
			[]query.Sorting{sorting.Asc("id")},
			3,
			0,
		),
	},
}

func TestFilterToSQL(t *testing.T) {
	for _, d := range filterToSQLDataProvider {
		t.Run(fmt.Sprint(d.Filter), func(t *testing.T) {
			b := NewStatementBuilder()
			if assert.NoError(t, b.WriteFilter(d.Filter)) {
				stmt := b.Build()
				assert.Equal(t, d.SQL, stmt.Query())
			}
		})
	}
}
