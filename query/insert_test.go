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

	assertQuery(t, "INSERT INTO table (columnA, columnB) VALUES (?, ?);", q)
	assertBinds(t, []interface{}{1}, b)
}
