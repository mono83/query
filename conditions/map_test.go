package conditions

import (
	"github.com/mono83/query/types"
	"testing"

	"github.com/mono83/query"
	"github.com/mono83/query/rules"
	"github.com/stretchr/testify/assert"
)

func TestMap(t *testing.T) {
	assert.Nil(t, Map(nil, nil))
	assert.Nil(t, Map(nil, func(query.Rule) query.Rule { return nil }))

	c := New(query.And, nil, nil)
	assert.Equal(t, c, Map(c, nil))

	c = New(
		query.And,
		[]query.Rule{
			rules.Eq("foo", "1"),
			rules.Eq("bar", "2"),
			rules.Eq("baz", "3"),
		},
		[]query.Condition{
			New(query.Or, []query.Rule{rules.Eq("foo", "1")}, nil),
		},
	)
	c = Map(c, func(r query.Rule) query.Rule {
		if r.Left() == "bar" {
			return nil
		}
		return rules.WithRight(r, "x"+types.ToString(r.Right()))
	})

	assert.Len(t, c.Rules(), 2)
	assert.Equal(t, c.Rules()[0].Right(), "x1")
	assert.Equal(t, c.Rules()[1].Right(), "x3")
	assert.Equal(t, c.Conditions()[0].Rules()[0].Right(), "x1")
}
