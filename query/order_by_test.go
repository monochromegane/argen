package query

import "testing"

func TestOrderBy(t *testing.T) {
	o := orderBy{}
	o.addOrder("columnA", ASC)
	o.addOrder("columnB", DESC)

	q := o.build()

	assertQuery(t, " ORDER BY columnA ASC, columnB DESC", q)
}
