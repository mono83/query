package conditions

import (
	"github.com/mono83/query"
	"github.com/mono83/query/rules"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCountRules(t *testing.T) {
	assert.Equal(t, 0, CountRules(nil))
	assert.Equal(t, 0, CountRules(ForAllRules()))
	assert.Equal(t, 1, CountRules(ForAllRules(rules.IsNull("foo"))))
	assert.Equal(
		t,
		3,
		CountRules(New(query.And, []query.Rule{rules.IsNotNull("bar")}, []query.Condition{ForAllRules(rules.IsNull("foo"), rules.IsNull("bar"))})),
	)
}
