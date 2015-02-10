package query

import "testing"

func TestOneCondition(t *testing.T) {
	c := condition{phrase: "WHERE"}
	c.addExpression("columnA = 'const'")

	q, b := c.build()

	expectedQuery := " WHERE columnA = 'const'"

	if q != expectedQuery {
		t.Errorf("query should be %s, but %s", expectedQuery, q)
	}

	if len(b) != 0 {
		t.Errorf("binds length should be 0, but %d", len(b))
	}
}

func TestTwoCondition(t *testing.T) {
	c := condition{phrase: "WHERE"}
	c.addExpression("columnA", "value")

	q, b := c.build()

	expectedQuery := " WHERE columnA = ?"
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

func TestThreeCondition(t *testing.T) {
	c := condition{phrase: "WHERE"}
	c.addExpression("columnA", "<>", "value")

	q, b := c.build()

	expectedQuery := " WHERE columnA <> ?"
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

func TestConditionWithMultiExpression(t *testing.T) {
	c := condition{phrase: "WHERE"}
	c.addExpression("columnA", "value1")
	c.addExpression("columnB", "value2")

	q, b := c.build()

	expectedQuery := " WHERE columnA = ? AND columnB = ?"
	expectedBinds := []interface{}{"value1", "value2"}

	if q != expectedQuery {
		t.Errorf("query should be %s, but %s", expectedQuery, q)
	}

	for i, v := range b {
		if v != expectedBinds[i] {
			t.Errorf("binds should be %v, but %v", expectedBinds[i], v)
		}
	}
}
