package query

import "testing"

func TestInnerJoin(t *testing.T) {
	j := innerJoin("table", "columnA", "value")
	q, b := j.build()

	assertQuery(t, " INNER JOIN table ON columnA = ?", q)
	assertBinds(t, []interface{}{"value"}, b)
}
