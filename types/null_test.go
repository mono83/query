package types

import (
	"database/sql"
	"fmt"
	"github.com/stretchr/testify/assert"
	"reflect"
	"testing"
)

var nilDataProvider = []struct {
	Expected bool
	Value    interface{}
}{
	{false, nil},
	{false, 5},
	{false, 5.},
	{false, "5"},
	{true, refInt(5)},
	{true, refString("5")},
	{true, []string{}},
	{true, map[int]string{}},
	{true, func() {}},

	{true, reflect.Ptr},
	{true, reflect.Slice},
	{true, reflect.Map},
	{true, reflect.Chan},
	{true, reflect.Func},
	{false, reflect.String},
	{false, reflect.Int},

	{false, sql.ColumnType{}},
}

func TestIsNilable(t *testing.T) {
	for _, datum := range nilDataProvider {
		t.Run(fmt.Sprint(datum), func(t *testing.T) {
			assert.Equal(t, datum.Expected, IsNullable(datum.Value))
		})
	}
}

func refInt(i int) *int          { return &i }
func refString(s string) *string { return &s }
