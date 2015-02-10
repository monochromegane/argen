package query

import "testing"

func TestInnerJoin(t *testing.T) {
	j := innerJoin("table", "columnA", "value")
	q, b := j.build()

	expectedQuery := " INNER JOIN table ON columnA = ?"
	expectedBinds := []interface{}{"value"}

	if q != expectedQuery {
		t.Errorf("query should be %s, but %s", expectedQuery, q)
	}
	for i, v := range b {
		if v != expectedBinds[i] {
			t.Errorf("binds should be %v, but %v", expectedBinds[i], v)
		}
	}
}
