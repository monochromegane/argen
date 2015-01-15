package gen

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
	Anotations []anotation
	Name       string
	Fields     []field
}

func (s structType) TableName() string {
	return toSnakeCase(s.Name)
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
			return f.Name, f.ColumnName(), f.typ
		}
	}
	return "ID", "id", "int"
}

type anotation string

func (a anotation) Arg() string {
	s := strings.Split(string(a), " ")
	if len(s) < 2 {
		return ""
	}
	return s[1]
}

func (a anotation) Option(key string) string {
	// similar to tag.Get()
	return ""
}

func (a anotation) BelongsTo() bool {
	return strings.HasPrefix(string(a), "+belongs_to")
}

func (a anotation) HasOne() bool {
	return strings.HasPrefix(string(a), "+has_one")
}

func (a anotation) HasMany() bool {
	return strings.HasPrefix(string(a), "+has_many")
}

func (a anotation) Scope() bool {
	return strings.HasPrefix(string(a), "+scope")
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
	return toSnakeCase(f.Name)
}

func (f field) isPrimaryKey() bool {
	if pk := f.tag.get("db"); pk == "pk" {
		return true
	}
	return false
}

func toSnakeCase(s string) string {
	const snake = "${1}_${2}"
	reg1 := regexp.MustCompile("([A-Z]+)([A-Z][a-z])")
	reg2 := regexp.MustCompile("([a-z])([A-Z])")
	return strings.ToLower(reg2.ReplaceAllString(reg1.ReplaceAllString(s, snake), snake))
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
