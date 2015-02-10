package query

import "testing"

func TestOneCondition(t *testing.T) {
	c := condition{phrase: "WHERE"}
	c.addExpression("columnA = 'const'")

	q, b := c.build()

	assertQuery(t, " WHERE columnA = 'const'", q)
	assertEmptyBinds(t, b)
}

func TestTwoCondition(t *testing.T) {
	c := condition{phrase: "WHERE"}
	c.addExpression("columnA", "value")

	q, b := c.build()

	assertQuery(t, " WHERE columnA = ?", q)
	assertBinds(t, []interface{}{"value"}, b)
}

func TestThreeCondition(t *testing.T) {
	c := condition{phrase: "WHERE"}
	c.addExpression("columnA", "<>", "value")

	q, b := c.build()

	assertQuery(t, " WHERE columnA <> ?", q)
	assertBinds(t, []interface{}{"value"}, b)
}

func TestConditionWithMultiExpression(t *testing.T) {
	c := condition{phrase: "WHERE"}
	c.addExpression("columnA", "value1")
	c.addExpression("columnB", "value2")

	q, b := c.build()

	assertQuery(t, " WHERE columnA = ? AND columnB = ?", q)
	assertBinds(t, []interface{}{"value1", "value2"}, b)
}
