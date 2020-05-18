package eval

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

var inDataProvider = []struct {
	Strict, NonStrict bool
	Left, Right       interface{}
}{
	{true, true, nil, nil},
	{true, true, 10, []int{5, 10, 20}},
	{false, true, 10., []interface{}{5, "10", 20}},
	{true, true, "foo", []string{"bar", "foo"}},
}

func TestEvaluator_in(t *testing.T) {
	strict := Evaluator{Strict: true}
	nonstrict := Evaluator{}
	for _, data := range inDataProvider {
		t.Run(fmt.Sprintf("%v == %v -> %v / %v", data.Left, data.Right, data.Strict, data.NonStrict), func(t *testing.T) {
			if data.Strict {
				assert.True(t, strict.in(data.Left, data.Right))
			} else {
				assert.False(t, strict.in(data.Left, data.Right))
			}
			if data.NonStrict {
				assert.True(t, nonstrict.in(data.Left, data.Right))
			} else {
				assert.False(t, nonstrict.in(data.Left, data.Right))
			}
		})
	}
}
