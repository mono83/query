package rules

import (
	"testing"

	"github.com/mono83/query/match"
	"github.com/stretchr/testify/assert"
)

func TestNewFull(t *testing.T) {
	r := New("foo", match.Neq, "bar")
	assert.Equal(t, "foo", r.Left())
	assert.Equal(t, "bar", r.Right())
	assert.Equal(t, match.Neq, r.Type())
	if x, ok := r.(full); assert.True(t, ok) {
		assert.Equal(t, "{Rule {foo (string)} Not_Equal {bar (string)}}", x.String())
	}
}

func TestNewLeftPartial(t *testing.T) {
	r := New("xyz", match.IsNotNil, nil)
	assert.Equal(t, "xyz", r.Left())
	assert.Nil(t, r.Right())
	assert.Equal(t, match.IsNotNil, r.Type())
	if x, ok := r.(leftPart); assert.True(t, ok) {
		assert.Equal(t, "{Rule {xyz (string)} Is_Not_Null}}", x.String())
	}
}
