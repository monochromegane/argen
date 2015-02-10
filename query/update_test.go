package query

import "testing"

func TestUpdate(t *testing.T) {
	update := Update{}
	update.Table("table")
	update.Params(map[string]interface{}{
		"columnA": "value1",
		"columnB": "value2",
	})
	update.Where("columnA", "value1")

	q, b := update.Build()

	assertQuery(t, "UPDATE table SET columnA = ?, columnB = ? WHERE columnA = ?", q)
	assertBinds(t, []interface{}{"value1", "value2", "value1"}, b)
}
