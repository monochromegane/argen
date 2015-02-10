package query

import "testing"

func TestOffset(t *testing.T) {
	o := offset{}
	o.setOffset(1)
	q, b := o.build()

	assertQuery(t, " OFFSET ?", q)
	assertBinds(t, []interface{}{1}, b)
}
