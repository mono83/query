package conditions

import (
	"testing"

	"github.com/mono83/query"
	"github.com/mono83/query/rules"
	"github.com/stretchr/testify/assert"
)

func TestIsEmpty(t *testing.T) {
	assert.True(t, IsEmpty(nil))
	assert.True(t, IsEmpty(New(query.And, nil, nil)))
	assert.True(t, IsEmpty(New(query.And, []query.Rule{}, []query.Condition{})))
	assert.False(t, IsEmpty(New(query.And, []query.Rule{rules.IsNull("foo")}, nil)))
	assert.False(t, IsEmpty(New(query.And, nil, []query.Condition{New(query.And, nil, nil)})))
}
