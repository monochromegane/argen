package gen

import (
	"fmt"
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

func (s structType) HasOne() []HasOne {
	var hasOne []HasOne
	for _, f := range s.Funcs {
		if f.HasOne() {
			hasOne = append(hasOne, HasOne{&s, f})
		}
	}
	return hasOne
}

func (s structType) HasMany() []HasMany {
	var hasMany []HasMany
	for _, f := range s.Funcs {
		if f.HasMany() {
			hasMany = append(hasMany, HasMany{&s, f})
		}
	}
	return hasMany
}

func (s structType) BelongsTo() []BelongsTo {
	var belongsTo []BelongsTo
	for _, f := range s.Funcs {
		if f.BelongsTo() {
			belongsTo = append(belongsTo, BelongsTo{&s, f})
		}
	}
	return belongsTo
}

func (s structType) Scope() []Scope {
	var scope []Scope
	for _, f := range s.Funcs {
		if f.scope() {
			scope = append(scope, Scope{&s, f})
		}
	}
	return scope
}

func (s structType) DefaultScope() bool {
	for _, f := range s.Funcs {
		if f.defaultScope() {
			return true
		}
	}
	return false
}

func (s structType) Validation() []Validation {
	var validation []Validation
	for _, f := range s.Funcs {
		if f.validation() {
			validation = append(validation, Validation{&s, f})
		}
	}
	return validation
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

func (f funcType) HasMany() bool {
	return strings.HasPrefix(f.Name, "hasMany")
}

func (f funcType) HasOne() bool {
	return strings.HasPrefix(f.Name, "hasOne")
}

func (f funcType) BelongsTo() bool {
	return strings.HasPrefix(f.Name, "belongsTo")
}

func (f funcType) scope() bool {
	return strings.HasPrefix(f.Name, "scope")
}

func (f funcType) defaultScope() bool {
	return f.Name == "defaultScope"
}

func (f funcType) validation() bool {
	return strings.HasPrefix(f.Name, "validates")
}

type HasOne struct {
	Recv *structType
	funcType
}

func (h HasOne) FuncName() string {
	return h.funcType.Name
}

func (h HasOne) Func() string {
	return strings.Replace(h.funcType.Name, "hasOne", "", 1)
}

func (h HasOne) Model() string {
	return inflector.Singularize(h.Func())
}

func (h HasOne) ForeignKey() string {
	return fmt.Sprintf("%s_id", toSnakeCase(h.funcType.Recv))
}

type HasMany struct {
	Recv *structType
	funcType
}

func (h HasMany) FuncName() string {
	return h.funcType.Name
}

func (h HasMany) Func() string {
	return strings.Replace(h.funcType.Name, "hasMany", "", 1)
}

func (h HasMany) Model() string {
	return inflector.Singularize(h.Func())
}

func (h HasMany) ForeignKey() string {
	return fmt.Sprintf("%s_id", toSnakeCase(h.funcType.Recv))
}

type BelongsTo struct {
	Recv *structType
	funcType
}

func (b BelongsTo) FuncName() string {
	return b.funcType.Name
}

func (b BelongsTo) Func() string {
	return strings.Replace(b.funcType.Name, "belongsTo", "", 1)
}

func (b BelongsTo) Model() string {
	return inflector.Singularize(b.Func())
}

func (b BelongsTo) PrimaryKey() string {
	return "id"
}

func (b BelongsTo) ForeignKey() string {
	return fmt.Sprintf("%s_id", toSnakeCase(b.Model()))
}

type Scope struct {
	Recv *structType
	funcType
}

func (s Scope) FuncName() string {
	return s.funcType.Name
}

func (s Scope) Func() string {
	return strings.Replace(s.funcType.Name, "scope", "", 1)
}

type Validation struct {
	Recv *structType
	funcType
}

func (v Validation) FuncName() string {
	return v.funcType.Name
}

func (v Validation) FieldName() string {
	return strings.Replace(v.funcType.Name, "validates", "", 1)
}

func toSnakeCase(s string) string {
	const snake = "${1}_${2}"
	reg1 := regexp.MustCompile("([A-Z]+)([A-Z][a-z])")
	reg2 := regexp.MustCompile("([a-z])([A-Z])")
	return strings.ToLower(reg2.ReplaceAllString(reg1.ReplaceAllString(s, snake), snake))
}
