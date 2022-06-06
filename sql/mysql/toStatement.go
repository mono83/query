package mysql

import (
	"github.com/mono83/query"
	"github.com/mono83/query/sql"
)

// QueryToStatement builds sql.Statement for provided query.Query
func QueryToStatement(q query.Query) (sql.Statement, error) {
	b := NewStatementBuilder()
	if err := b.WriteQuery(q); err != nil {
		return nil, err
	}

	return b.Build(), nil
}

// FilterToStatement constructs partial statement for provided query.Filter
func FilterToStatement(f query.Filter) (sql.Statement, error) {
	b := NewStatementBuilder()
	if err := b.WriteFilter(f); err != nil {
		return nil, err
	}

	return b.Build(), nil
}

// ConditionToStatement constructs partial statement for provided query.Condition
func ConditionToStatement(c query.Condition) (sql.Statement, error) {
	b := NewStatementBuilder()
	if err := b.WriteCondition(c); err != nil {
		return nil, err
	}

	return b.Build(), nil
}
