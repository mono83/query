package eval

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

var equalsDataProvider = []struct {
	Strict, NonStrict bool
	Left, Right       interface{}
}{
	{true, true, nil, nil},
	{true, true, "foo", "foo"},
	{true, true, 1, 1},
	{true, true, int8(5), int8(5)},
	{true, true, float64(-44), float64(-44)},
	{true, true, float32(12), float32(12)},
	{false, true, "1", 1},
	{false, true, "1", int8(1)},
	{false, true, "1", uint64(1)},
	{false, true, 22, float64(22)},
	{false, true, -3, float32(-3)},
}

func TestEvaluator_equals(t *testing.T) {
	strict := Evaluator{Strict: true}
	nonstrict := Evaluator{}
	for _, data := range equalsDataProvider {
		t.Run(fmt.Sprintf("%v == %v -> %v / %v", data.Left, data.Right, data.Strict, data.NonStrict), func(t *testing.T) {
			if data.Strict {
				assert.True(t, strict.equals(data.Left, data.Right))
			} else {
				assert.False(t, strict.equals(data.Left, data.Right))
			}
			if data.NonStrict {
				assert.True(t, nonstrict.equals(data.Left, data.Right))
			} else {
				assert.False(t, nonstrict.equals(data.Left, data.Right))
			}
		})
	}
}
