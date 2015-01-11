package generator

import (
	"regexp"
	"strings"
)

type structs []structType

func (s structs) Package() string {
	if len(s) == 0 {
		return ""
	}
	return s[0].pkg
}

type structType struct {
	pkg        string
	anotations []string
	Name       string
	Fields     []field
}

func (s structType) FieldNames(prefix string) []string {
	names := []string{}
	for _, f := range s.Fields {
		names = append(names, prefix+f.Name)
	}
	return names
}

func (s structType) ColumnNames() []string {
	names := []string{}
	for _, f := range s.Fields {
		names = append(names, f.ColumnName())
	}
	return names
}

func (s structType) PrimaryKey() string {
	k, _ := s.primaryKey()
	return k
}

func (s structType) PrimaryKeyType() string {
	_, t := s.primaryKey()
	return t
}

func (s structType) primaryKey() (string, string) {
	for _, f := range s.Fields {
		if f.isPrimaryKey() {
			return f.Name, f.typ
		}
	}
	return "id", "int"
}

type field struct {
	typ  string
	Name string
	tag  tag
}

func (f field) ColumnName() string {
	if tag := f.tag.get("column"); tag != "" {
		return tag
	}
	return f.toSnakeCase()
}

func (f field) isPrimaryKey() bool {
	if pk := f.tag.get("db"); pk == "pk" {
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
		if !strings.HasPrefix(tag, key+":") {
			continue
		}
		pair := strings.SplitN(tag, ":", 1)
		if !strings.HasPrefix(pair[1], "\"") || !strings.HasSuffix(pair[1], "\"") {
			continue
		}
		return strings.Trim(pair[1], "\"")
	}
	return value
}

func (f field) toSnakeCase() string {
	const snake = "${1}_${2}"
	reg1 := regexp.MustCompile("([A-Z]+)([A-Z][a-z])")
	reg2 := regexp.MustCompile("([a-z])([A-Z])")
	return strings.ToLower(reg2.ReplaceAllString(reg1.ReplaceAllString(f.Name, snake), snake))
}
