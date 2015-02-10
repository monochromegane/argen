package query

import "testing"

func TestGroupBy(t *testing.T) {
	g := groupBy{}
	g.setGroups("columnA", "columnB")

	q := g.build()

	assertQuery(t, " GROUP BY columnA, columnB", q)
}
