package rules

import (
	"testing"

	"github.com/mono83/query"
	"github.com/mono83/query/match"
	"github.com/stretchr/testify/assert"
)

func TestEq(t *testing.T) {
	r := Eq("foo", "bar")
	assert.Equal(t, "foo", r.GetLeft())
	assert.Equal(t, "bar", r.GetRight())
	assert.Equal(t, match.Eq, r.GetType())
}

func TestIsNull(t *testing.T) {
	r := IsNull("foo")
	assert.Equal(t, "foo", r.GetLeft())
	assert.Nil(t, r.GetRight())
	assert.Equal(t, match.IsNull, r.GetType())
}

func TestIsNotNull(t *testing.T) {
	r := IsNotNull("foo")
	assert.Equal(t, "foo", r.GetLeft())
	assert.Nil(t, r.GetRight())
	assert.Equal(t, match.NotIsNull, r.GetType())
}

func TestMatchID64(t *testing.T) {
	var r query.Rule

	r = MatchID64()
	assert.IsType(t, False{}, r)

	r = MatchID64(10)
	assert.Equal(t, match.Eq, r.GetType())
	assert.Equal(t, "id", r.GetLeft())
	assert.Equal(t, int64(10), r.GetRight())

	r = MatchID64(10, 11, 12)
	assert.Equal(t, match.In, r.GetType())
	assert.Equal(t, "id", r.GetLeft())
	v, ok := r.GetRight().([]int64)
	if assert.True(t, ok) {
		assert.Len(t, r.GetRight(), 3)
		assert.Equal(t, int64(10), v[0])
		assert.Equal(t, int64(11), v[1])
		assert.Equal(t, int64(12), v[2])
	}
}
