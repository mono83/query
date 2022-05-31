package mysql

import (
	"testing"

	"github.com/mono83/query"
	"github.com/mono83/query/names"
	"github.com/stretchr/testify/assert"
)

var keysDataProvider = []struct {
	Expected string
	Provided string
}{
	{"`id`", "id"},
	{"`id`", "`id`"},
	{"``id`", "`id"}, // Invalid behaviour
}

func TestWriteKey(t *testing.T) {
	for _, d := range keysDataProvider {
		t.Run(d.Provided, func(t *testing.T) {
			assert.Equal(t, d.Expected, NewStatementBuilder().WriteKey(d.Provided).Build().Query())
		})
	}
}

var namedsDataProvider = []struct {
	ExpectedColumn  string
	ExpectedSchema  string
	ExpectedSorting string
	Source          query.Named
}{
	{"`foo`", "`foo`", "`foo`", names.String("foo")},
	{"`foo`", "`foo`", "`foo`", names.String("`foo`")},
	{"`bar` as `zzz`", "`bar`", "`bar`", names.Alias("bar", "zzz")},
	{"`bar` as `zzz`", "`bar`", "`bar`", names.Alias("`bar`", "`zzz`")},
}

func TestWriteNameds(t *testing.T) {
	var b *StatementBuilder
	for _, d := range namedsDataProvider {
		t.Run(d.ExpectedColumn, func(t *testing.T) {
			b = NewStatementBuilder()
			if assert.NoError(t, b.WriteColumn(d.Source)) {
				assert.Equal(t, d.ExpectedColumn, b.Build().Query())
			}

			b = NewStatementBuilder()
			if assert.NoError(t, b.WriteSchema(d.Source)) {
				assert.Equal(t, d.ExpectedSchema, b.Build().Query())
			}

			b = NewStatementBuilder()
			b.WriteNamed(d.Source)
			assert.Equal(t, d.ExpectedSorting, b.Build().Query())
		})
	}
}
