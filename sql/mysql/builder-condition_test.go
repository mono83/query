package mysql

import (
	"fmt"
	"testing"

	"github.com/mono83/query"
	"github.com/mono83/query/conditions"
	"github.com/mono83/query/match"
	"github.com/mono83/query/names"
	"github.com/mono83/query/rules"
	"github.com/stretchr/testify/assert"
)

var conditionToSQLDataProvider = []struct {
	SQL          string
	Placeholders []interface{}
	Condition    query.Condition
}{
	{"`name` = ?", []interface{}{"bar"}, conditions.ForAllRules(rules.Eq(names.String("name"), "bar"))},
	{
		"(`id` > ? AND `name` = ?)",
		[]interface{}{10, "bar"},
		conditions.ForAllRules(
			rules.New(names.String("id"), match.GreaterThan, 10),
			rules.Eq(names.String("name"), "bar"),
		),
	},
	{
		"(`id` <= ? OR `name` = ?)",
		[]interface{}{3, "bar"},
		conditions.ForAnyRule(
			rules.New(names.String("id"), match.LesserThanEquals, 3),
			rules.Eq(names.String("name"), "bar"),
		),
	},
	{
		"(`type` = ? OR (`lastLoginAt` = `firstLoginAt` OR `blockedAt` IS NULL) OR (`scope` = ? AND `type` = ?))",
		[]interface{}{"admin", "current", "user"},
		conditions.New(
			query.Or,
			[]query.Rule{rules.Eq(names.String("type"), "admin")},
			[]query.Condition{
				conditions.ForAnyRule(
					rules.Eq(names.String("lastLoginAt"), names.String("firstLoginAt")),
					rules.IsNull(names.String("blockedAt")),
				),
				conditions.ForAllRules(
					rules.Eq(names.String("scope"), "current"),
					rules.Eq(names.String("type"), "user"),
				),
			},
		),
	},
}

func TestConditionToSQL(t *testing.T) {
	for _, d := range conditionToSQLDataProvider {
		t.Run(fmt.Sprint(d.Condition), func(t *testing.T) {
			b := NewStatementBuilder()
			if assert.NoError(t, b.WriteCondition(d.Condition)) {
				stmt := b.Build()
				if assert.Equal(t, d.SQL, stmt.Query()) {
					if assert.Equal(t, len(d.Placeholders), len(stmt.Args()), "parameters count don't match") {
						for i, a := range stmt.Args() {
							assert.Equal(t, d.Placeholders[i], a)
						}
					}
				}
			}
		})
	}
}
