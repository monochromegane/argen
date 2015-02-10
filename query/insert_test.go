package query

import "testing"

func testInsert(t *testing.T) {
	insert := Insert{}
	insert.Table("table")
	insert.Params(map[string]interface{}{
		"columnA": "value1",
		"columnB": "value2",
	})
	q, b := insert.Build()

	expectedQuery := "INSERT INTO table (columnA, columnB) VALUES (?, ?);"
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
