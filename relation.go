package ar

import (
	"database/sql"
	"reflect"
	"strings"
	"time"
	"unicode"

	"github.com/monochromegane/argen/query"
	"github.com/monochromegane/goban"
)

type Relation struct {
	*query.Select
	db     *sql.DB
	logger *Logger
}

func NewRelation(db *sql.DB, logger *Logger) *Relation {
	return &Relation{
		Select: &query.Select{},
		db:     db,
		logger: logger,
	}
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

func (r *Relation) Build() (string, []interface{}) {
	return r.Select.Build()
}

func (r *Relation) Explain() error {
	r.Select.Explain()
	rows, err := r.Query()
	if err != nil {
		return err
	}
	defer rows.Close()

	columns, _ := rows.Columns()
	var values [][]string
	for rows.Next() {
		vals := make([]string, len(columns))
		ptrs := make([]interface{}, len(columns))
		for i, _ := range vals {
			ptrs[i] = &vals[i]
		}
		rows.Scan(ptrs...)
		values = append(values, vals)
	}

	goban.Render(columns, values)
	return nil
}

func (r *Relation) Query() (*sql.Rows, error) {
	q, b := r.Build()
	defer r.log(time.Now(), q, b...)
	return r.db.Query(q, b...)
}

func (r *Relation) QueryRow(dest ...interface{}) error {
	q, b := r.Build()
	defer r.log(time.Now(), q, b...)
	return r.db.QueryRow(q, b...).Scan(dest...)
}

func (r *Relation) log(t time.Time, sql string, args ...interface{}) {
	r.logger.Print(time.Now().Sub(t), sql, args)
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
