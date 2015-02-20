package query

import "testing"

func TestSelect(t *testing.T) {
	s := Select{}
	s.Table("table")
	s.Columns("columnA", "columnB")

	q, b := s.Build()

	assertQuery(t, "SELECT table.columnA, table.columnB FROM table;", q)
	assertEmptyBinds(t, b)
}

func TestSelectExplain(t *testing.T) {
	s := Select{}
	s.Table("table")
	s.Columns("columnA", "columnB")

	q, b := s.Explain().Build()

	assertQuery(t, "EXPLAIN SELECT table.columnA, table.columnB FROM table;", q)
	assertEmptyBinds(t, b)
}

func TestSelectInnerJoin(t *testing.T) {
	s := Select{}
	s.Table("table")
	s.Columns("columnA", "columnB")
	s.InnerJoin("tableB", "id = tableB.table_id")

	q, b := s.Build()

	assertQuery(t, "SELECT table.columnA, table.columnB FROM table INNER JOIN tableB ON id = tableB.table_id;", q)
	assertEmptyBinds(t, b)
}

func TestSelectWhere(t *testing.T) {
	s := Select{}
	s.Table("table")
	s.Columns("columnA", "columnB")
	s.Where("columnA", "value")

	q, b := s.Build()

	assertQuery(t, "SELECT table.columnA, table.columnB FROM table WHERE columnA = ?;", q)
	assertBinds(t, []interface{}{"value"}, b)
}

func TestSelectLimitAndOffset(t *testing.T) {
	s := Select{}
	s.Table("table")
	s.Columns("columnA", "columnB")
	s.Limit(1)
	s.Offset(2)

	q, b := s.Build()

	assertQuery(t, "SELECT table.columnA, table.columnB FROM table LIMIT ? OFFSET ?;", q)
	assertBinds(t, []interface{}{1, 2}, b)
}

func TestSelectGroupByAndHaving(t *testing.T) {
	s := Select{}
	s.Table("table")
	s.Columns("columnA", "columnB")
	s.GroupBy("columnA", "columnB")
	s.Having("columnA", "value")

	q, b := s.Build()

	assertQuery(t, "SELECT table.columnA, table.columnB FROM table GROUP BY columnA, columnB HAVING columnA = ?;", q)
	assertBinds(t, []interface{}{"value"}, b)
}

func assertQuery(t *testing.T, expect, actual string) {
	if expect != actual {
		t.Errorf("query should be %s, but %s", expect, actual)
	}
}

func TestSelectOrderBy(t *testing.T) {
	s := Select{}
	s.Table("table")
	s.Columns("columnA", "columnB")
	s.OrderBy("columnA", ASC)

	q, b := s.Build()

	assertQuery(t, "SELECT table.columnA, table.columnB FROM table ORDER BY columnA ASC;", q)
	assertEmptyBinds(t, b)
}

func assertBinds(t *testing.T, expect, actual []interface{}) {
	for i, a := range actual {
		if expect[i] != a {
			t.Errorf("binds should be %v, but %v", expect[i], a)
		}
	}
}

func assertEmptyBinds(t *testing.T, actual []interface{}) {
	if len(actual) != 0 {
		t.Errorf("binds length should be 0, but %d", len(actual))
	}
}
