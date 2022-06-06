package eval

import (
	"fmt"
	"testing"

	"github.com/mono83/query"
	"github.com/mono83/query/match"
	"github.com/mono83/query/rules"
	"github.com/stretchr/testify/assert"
)

func TestAllMatchOperators(t *testing.T) {
	if supported, all := 11, match.Count(); supported != all {
		t.Errorf("Seems like not every of %d match operators are supported by eval, that handles %d", all, supported)
	}
}

var testEvalRuleDataProvider = []struct {
	Strict, NonStrict bool
	Rule              query.Rule
}{
	{true, true, rules.New("foo", match.IsNull, nil)},
	{false, false, rules.New("foo", match.NotIsNull, nil)},
	{true, true, rules.New("foo", match.NotIsNull, "")},
	{true, true, rules.New("foo", match.NotIsNull, 0)},
	{false, false, rules.New("foo", match.NotIsNull, nil)},

	{true, true, rules.New("foo", match.Eq, "foo")},
	{false, false, rules.New("foo", match.Eq, "bar")},
	{false, true, rules.New(42, match.Eq, int8(42))},
	{false, true, rules.New(42, match.Eq, int16(42))},
	{false, true, rules.New(42, match.Eq, int32(42))},
	{false, true, rules.New(42, match.Eq, int64(42))},
	{false, true, rules.New(42, match.Eq, uint8(42))},
	{false, true, rules.New(42, match.Eq, uint16(42))},
	{false, true, rules.New(42, match.Eq, uint32(42))},
	{false, true, rules.New(42, match.Eq, uint64(42))},
	{false, true, rules.New(42, match.Eq, float32(42))},
	{false, true, rules.New(42, match.Eq, float64(42))},
	{true, true, rules.New("foo", match.Neq, "bar")},

	{true, true, rules.New(33, match.Gt, 1)},
	{false, false, rules.New(4, match.Gt, 4)},
	{false, true, rules.New(31, match.Gt, -2.)},

	{true, true, rules.New(31, match.Gte, 2)},
	{true, true, rules.New(4, match.Gte, 4)},
	{false, true, rules.New(31, match.Gte, -2.)},

	{true, true, rules.New(4, match.Lt, 18)},
	{false, false, rules.New(-2, match.Lt, -2)},
	{false, true, rules.New(3, match.Lt, 18.)},

	{true, true, rules.New(-1, match.Lte, 1)},
	{true, true, rules.New(9, match.Lte, 9)},
	{false, true, rules.New(1, match.Lte, 2.)},

	{true, true, rules.New(4, match.In, []int{2, 3, 4})},
	{true, true, rules.New(4, match.In, []interface{}{2, 3, 4})},
	{false, true, rules.New(4, match.In, []interface{}{2, 3, 4.})},
	{false, false, rules.New(4, match.In, []int{2, 3})},

	{true, true, rules.New(4, match.NotIn, []int{1, 2, 3})},
	{false, false, rules.New(4, match.NotIn, []int{1, 2, 3, 4})},
}

func TestEvaluator_Rule(t *testing.T) {
	strict := Evaluator{Strict: true}
	nonstrict := Evaluator{}
	for _, data := range testEvalRuleDataProvider {
		t.Run(fmt.Sprintf("%v -> %v / %v", data.Rule, data.Strict, data.NonStrict), func(t *testing.T) {
			if data.Strict {
				assert.True(t, strict.Rule(data.Rule), "rule %s must be true in strict mode", data.Rule)
			} else {
				assert.False(t, strict.Rule(data.Rule), "rule %s must be false in strict mode", data.Rule)
			}
			if data.NonStrict {
				assert.True(t, nonstrict.Rule(data.Rule), "rule %s must be true in strict mode", data.Rule)
			} else {
				assert.False(t, nonstrict.Rule(data.Rule), "rule %s must be false in strict mode", data.Rule)
			}
		})
	}
}
