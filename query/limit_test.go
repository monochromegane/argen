package query

import "testing"

func TestLimit(t *testing.T) {
	l := limit{}
	l.setLimit(1)
	q, b := l.build()

	expectedQuery := " LIMIT ?"
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
