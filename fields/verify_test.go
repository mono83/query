package fields

import (
	"fmt"
	"github.com/mono83/query"
	"github.com/mono83/query/conditions"
	"github.com/mono83/query/filters"
	"github.com/mono83/query/rules"
	"github.com/mono83/query/sorting"
	"github.com/stretchr/testify/assert"
	"testing"
)

var testVerifyData = []struct {
	Error   string
	Rules   []query.Rule
	Sorting []query.Sorting
}{
	{
		Error: ``,
	},
	{
		Error:   ``,
		Sorting: []query.Sorting{sorting.Field(Indexed("id"), query.Asc)},
	},
	{
		Error:   `field "id" does not support sorting`,
		Sorting: []query.Sorting{sorting.Field(New("id", true, false), query.Asc)},
	},
	{
		Error:   `no field found in sorting by "id"`,
		Sorting: []query.Sorting{sorting.String("id", query.Asc)},
	},
	{
		Error: ``,
		Rules: []query.Rule{
			rules.Eq(Indexed("id"), 10),
			rules.IsNull(New("nextAt", true, false)),
		},
	},
	{
		Error: `field "nextAt" used in {Rule {{"nextAt",nofilter,sort} (fields.field)} Is_Null}} does not support filtering`,
		Rules: []query.Rule{
			rules.IsNull(New("nextAt", false, true)),
		},
	},
	{
		Error: `no field found in rule {Rule {nextAt (string)} Is_Null}}`,
		Rules: []query.Rule{
			rules.IsNull("nextAt"),
		},
	},
}

func TestVerifyFilter(t *testing.T) {

	for _, datum := range testVerifyData {
		t.Run(fmt.Sprint(datum), func(t *testing.T) {
			filter := filters.FromCondition(
				conditions.ForAllRules(datum.Rules...),
				datum.Sorting,
				1,
				0,
			)

			if len(datum.Error) > 0 {
				err := VerifyFilter(filter)
				if assert.Error(t, err) {
					assert.Equal(t, datum.Error, err.Error())
				}
			} else {
				assert.NoError(t, VerifyFilter(filter))
			}
		})
	}
}
