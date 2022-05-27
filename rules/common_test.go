package rules

import (
	"testing"

	"github.com/mono83/query"
	"github.com/mono83/query/match"
	"github.com/stretchr/testify/assert"
)

func TestEq(t *testing.T) {
	r := Eq("foo", "bar")
	assert.Equal(t, "foo", r.Left())
	assert.Equal(t, "bar", r.Right())
	assert.Equal(t, match.Eq, r.Type())
}

func TestIsNull(t *testing.T) {
	r := IsNull("foo")
	assert.Equal(t, "foo", r.Left())
	assert.Nil(t, r.Right())
	assert.Equal(t, match.IsNull, r.Type())
}

func TestIsNotNull(t *testing.T) {
	r := IsNotNull("foo")
	assert.Equal(t, "foo", r.Left())
	assert.Nil(t, r.Right())
	assert.Equal(t, match.NotIsNull, r.Type())
}

func TestMatchID64(t *testing.T) {
	var r query.Rule

	r = MatchID64()
	assert.IsType(t, False, r)

	r = MatchID64(10)
	assert.Equal(t, match.Eq, r.Type())
	assert.Equal(t, "id", r.Left())
	assert.Equal(t, int64(10), r.Right())

	r = MatchID64(10, 11, 12)
	assert.Equal(t, match.In, r.Type())
	assert.Equal(t, "id", r.Left())
	v, ok := r.Right().([]int64)
	if assert.True(t, ok) {
		assert.Len(t, r.Right(), 3)
		assert.Equal(t, int64(10), v[0])
		assert.Equal(t, int64(11), v[1])
		assert.Equal(t, int64(12), v[2])
	}
}
