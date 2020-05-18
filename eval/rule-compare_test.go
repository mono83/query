package eval

import (
	"fmt"
	"testing"

	"github.com/mono83/query/match"
	"github.com/stretchr/testify/assert"
)

var compDataProvider = []struct {
	Strict, NonStrict bool
	Operator          match.Type
	Left, Right       interface{}
}{
	{true, true, match.GreaterThan, 5, 4},
	{false, false, match.GreaterThan, 4, 4},
	{true, true, match.GreaterThanEquals, 10, 2},
	{true, true, match.GreaterThanEquals, 11, 11},
	{true, true, match.LesserThan, -2, 1},
	{false, false, match.LesserThan, 0, 0},
	{true, true, match.LesserThanEquals, 99, 122},
	{true, true, match.LesserThanEquals, 100, 100},

	{false, true, match.GreaterThan, 5., 4},
}

func TestEvaluator_comp(t *testing.T) {
	strict := Evaluator{Strict: true}
	nonstrict := Evaluator{}
	for _, data := range compDataProvider {
		t.Run(fmt.Sprintf("[%s] %v == %v -> %v / %v", data.Operator.String(), data.Left, data.Right, data.Strict, data.NonStrict), func(t *testing.T) {
			if data.Strict {
				assert.True(t, strict.comp(data.Left, data.Right, data.Operator))
			} else {
				assert.False(t, strict.comp(data.Left, data.Right, data.Operator))
			}
			if data.NonStrict {
				assert.True(t, nonstrict.comp(data.Left, data.Right, data.Operator))
			} else {
				assert.False(t, nonstrict.comp(data.Left, data.Right, data.Operator))
			}
		})
	}
}
