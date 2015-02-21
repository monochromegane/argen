package query

import "testing"

func TestDelete(t *testing.T) {
	del := Delete{}
	del.Table("table")
	del.Where("columnA", "value1")

	q, b := del.Build()

	assertQuery(t, "DELETE FROM table WHERE columnA = ?;", q)
	assertBinds(t, []interface{}{"value1", "value2", "value1"}, b)
}
