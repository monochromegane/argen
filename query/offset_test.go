package query

import "testing"

func TestOffset(t *testing.T) {
	o := offset{}
	o.setOffset(1)
	q, b := o.build()

	expectedQuery := " OFFSET ?"
	expectedBinds := []interface{}{1}

	if q != expectedQuery {
		t.Errorf("query should be %s, but %s", expectedQuery, q)
	}
	for i, v := range b {
		if v != expectedBinds[i] {
			t.Errorf("binds should be %v, but %v", expectedBinds[i], v)
		}
	}
}
