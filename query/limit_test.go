package query

import "testing"

func TestLimit(t *testing.T) {
	l := limit{}
	l.setLimit(1)
	q, b := l.build()

	assertQuery(t, " LIMIT ?", q)
	assertBinds(t, []interface{}{1}, b)
}
