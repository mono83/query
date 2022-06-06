package eval

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

var containDataProvider = []struct {
	Strict, NonStrict bool
	Left, Right       interface{}
}{
	{false, false, nil, nil},
	{false, false, "", nil},
	{false, false, nil, ""},
	{true, true, "", ""},
	{true, true, "FOO BAR", "BAR"},
	{false, false, "FOO BAR", "bar"},
	{false, true, "12345", 45},
}

func TestEvaluator_contains(t *testing.T) {
	strict := Evaluator{Strict: true}
	nonstrict := Evaluator{}
	for _, data := range containDataProvider {
		t.Run(fmt.Sprintf("%v LIKE %v -> %v / %v", data.Left, data.Right, data.Strict, data.NonStrict), func(t *testing.T) {
			if data.Strict {
				assert.True(t, strict.contains(data.Left, data.Right))
			} else {
				assert.False(t, strict.contains(data.Left, data.Right))
			}
			if data.NonStrict {
				assert.True(t, nonstrict.contains(data.Left, data.Right))
			} else {
				assert.False(t, nonstrict.contains(data.Left, data.Right))
			}
		})
	}
}
