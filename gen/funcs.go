package gen

import (
	"fmt"
	"regexp"
	"strings"

	"github.com/gedex/inflector"
)

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

func (f funcType) joins() bool {
	return f.HasMany() || f.HasOne() || f.BelongsTo()
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

func (f funcType) customValidation() bool {
	match, _ := regexp.MatchString("^validate[A-Z0-9]", f.Name)
	return match
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

func (h HasOne) TableName() string {
	return inflector.Pluralize(toSnakeCase(h.Func()))
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

func (h HasMany) TableName() string {
	return toSnakeCase(h.Func())
}

func (h HasMany) ForeignKey() string {
	return fmt.Sprintf("%s_id", toSnakeCase(h.funcType.Recv))
}

func (h HasMany) ForeignKeyField() string {
	return fmt.Sprintf("%sId", h.funcType.Recv)
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

func (b BelongsTo) TableName() string {
	return inflector.Pluralize(toSnakeCase(b.Func()))
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

func (v Validation) ColumnName() string {
	return toSnakeCase(v.FieldName())
}
