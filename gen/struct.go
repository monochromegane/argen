package gen

import (
	"regexp"
	"strings"

	"github.com/gedex/inflector"
)

type structs []*structType

func (ss structs) Package() string {
	return ss[0].Package
}

type structType struct {
	Package  string
	Comments comments
	Name     string
	Fields   []field
	Funcs    funcs
}

func (s structType) TableName() string {
	return toSnakeCase(inflector.Pluralize(s.Name))
}

func (s structType) PrimaryKeyField() string {
	f, _, _ := s.primaryKey()
	return f
}

func (s structType) PrimaryKeyColumn() string {
	_, c, _ := s.primaryKey()
	return c
}

func (s structType) PrimaryKeyType() string {
	_, _, t := s.primaryKey()
	return t
}

func (s structType) primaryKey() (string, string, string) {
	for _, f := range s.Fields {
		if f.isPrimaryKey() {
			return f.Name, f.ColumnName(), f.Type
		}
	}
	return "ID", "id", "int"
}

type comments []comment

type comment string

type field struct {
	Type string
	Name string
	Tag  tag
}

func (f field) ColumnName() string {
	return toSnakeCase(f.Name)
}

func (f field) isPrimaryKey() bool {
	if pk := f.Tag.get("db"); pk == "pk" {
		return true
	}
	return false
}

type tag string

func (t tag) tags() []string {
	return strings.Split(string(t), " ")
}

func (t tag) get(key string) string {
	var value string
	for _, tag := range t.tags() {
		tag = strings.Trim(tag, "`")
		if !strings.HasPrefix(tag, key+":") {
			continue
		}
		pair := strings.SplitN(tag, ":", 2)
		if !strings.HasPrefix(pair[1], "\"") || !strings.HasSuffix(pair[1], "\"") {
			continue
		}
		return strings.Trim(pair[1], "\"")
	}
	return value
}

type funcs []funcType

type funcType struct {
	Recv     string
	Comments comments
	Name     string
}

func toSnakeCase(s string) string {
	const snake = "${1}_${2}"
	reg1 := regexp.MustCompile("([A-Z]+)([A-Z][a-z])")
	reg2 := regexp.MustCompile("([a-z])([A-Z])")
	return strings.ToLower(reg2.ReplaceAllString(reg1.ReplaceAllString(s, snake), snake))
}
