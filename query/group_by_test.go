package query

import "testing"

func TestGroupBy(t *testing.T) {
	g := groupBy{}
	g.setGroups("columnA", "columnB")

	q := g.build()

	expectedQuery := " GROUP BY columnA, columnB"

	if q != expectedQuery {
		t.Errorf("query should be %s, but %s", expectedQuery, q)
	}
}
