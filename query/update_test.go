package query

import "testing"

func TestUpdate(t *testing.T) {
	update := Update{}
	update.Table("table")
	update.Params(map[string]interface{}{
		"columnA": "value1",
	})
	update.Where("columnA", "value1")

	q, b := update.Build()

	assertQuery(t, "UPDATE table SET columnA = ? WHERE columnA = ?", q)
	assertBinds(t, []interface{}{"value1", "value1"}, b)
}
