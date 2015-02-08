package ar

import (
	"reflect"
	"strings"
	"unicode"

	"github.com/monochromegane/argen/query"
)

type Relation struct {
	*query.Select
}

type Insert struct {
	*query.Insert
}

type Update struct {
	*query.Update
}

func (r *Relation) Table(table string) *Relation {
	r.Select.Table(table)
	return r
}

func (r *Relation) Columns(columns ...string) *Relation {
	r.Select.Columns(columns...)
	return r
}

func (r *Relation) GetColumns() []string {
	return r.Select.GetColumns()
}

func (r *Relation) Where(cond string, args ...interface{}) *Relation {
	r.Select.Where(cond, args...)
	return r
}

func (r *Relation) And(cond string, args ...interface{}) *Relation {
	return r.Where(cond, args...)
}

func (r *Relation) OrderBy(column, order string) *Relation {
	r.Select.OrderBy(column, order)
	return r
}

func (r *Relation) Limit(limit int) *Relation {
	r.Select.Limit(limit)
	return r
}

func (r *Relation) Offset(offset int) *Relation {
	r.Select.Offset(offset)
	return r
}

func (r *Relation) GroupBy(group string, groups ...string) *Relation {
	r.Select.GroupBy(group, groups...)
	return r
}

func (r *Relation) Having(cond string, args ...interface{}) *Relation {
	r.Select.Having(cond, args...)
	return r
}

func (r *Relation) Explain() *Relation {
	r.Select.Explain()
	return r
}

func (r *Relation) Build() (string, []interface{}) {
	return r.Select.Build()
}

func IsZero(v interface{}) bool {
	return reflect.ValueOf(v).Interface() == reflect.Zero(reflect.TypeOf(v)).Interface()
}

func ToCamelCase(s string) string {
	var camel string
	for _, split := range strings.Split(s, "_") {
		c := []rune(split)
		c[0] = unicode.ToUpper(c[0])
		camel += string(c)
	}
	return camel
}
