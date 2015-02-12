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

func (s structType) FieldsWithoutPrimaryKey() []field {
	fields := []field{}
	for _, f := range s.Fields {
		if f.isPrimaryKey() {
			continue
		}
		fields = append(fields, f)
	}
	return fields
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

func (s structType) Joins() []Joins {
	var joinses []Joins
	for _, f := range s.Funcs {
		if f.joins() {
			joinses = append(joinses, Joins{&s, f})
		}
	}
	return joinses
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

func (s structType) CustomValidation() []Validation {
	var validation []Validation
	for _, f := range s.Funcs {
		if f.customValidation() {
			validation = append(validation, Validation{&s, f})
		}
	}
	return validation
}

func toSnakeCase(s string) string {
	const snake = "${1}_${2}"
	reg1 := regexp.MustCompile("([A-Z]+)([A-Z][a-z])")
	reg2 := regexp.MustCompile("([a-z])([A-Z])")
	return strings.ToLower(reg2.ReplaceAllString(reg1.ReplaceAllString(s, snake), snake))
}
