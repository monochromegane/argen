package query

import "testing"

func TestOrderBy(t *testing.T) {
	o := orderBy{}
	o.addOrder("columnA", ASC)
	o.addOrder("columnB", DESC)

	q := o.build()

	expectedQuery := " ORDER BY columnA ASC, columnB DESC"

	if q != expectedQuery {
		t.Errorf("query should be %s, but %s", expectedQuery, q)
	}
}
