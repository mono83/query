package query

import (
	"testing"

	"github.com/mono83/query/match"
	"github.com/stretchr/testify/assert"
)

func TestCommonRule(t *testing.T) {
	r := CommonRule{Left: 1, Right: "foo", Type: match.Equals}
	assert.Equal(t, 1, r.GetLeft())
	assert.Equal(t, "foo", r.GetRight())
	assert.Equal(t, match.Equals, r.GetType())
	assert.Equal(t, "{Rule {1 (int)} Equal {foo (string)}}", r.String())
}

func TestCommonRuleEmpty(t *testing.T) {
	r := CommonRule{}
	assert.Equal(t, nil, r.GetLeft())
	assert.Equal(t, nil, r.GetRight())
	assert.Equal(t, match.Unknown, r.GetType())
	assert.Equal(t, "{Rule {<nil> (<nil>)} Unknown {<nil> (<nil>)}}", r.String())
}
