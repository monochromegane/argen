package ar

import (
	"database/sql"

	"github.com/monochromegane/argen/query"
)

type Delete struct {
	*query.Delete
	exec *Executer
}

func NewDelete(db *sql.DB, logger *Logger) *Delete {
	return &Delete{
		Delete: &query.Delete{},
		exec:   &Executer{db, logger},
	}
}

func (d *Delete) Table(table string) *Delete {
	d.Delete.Table(table)
	return d
}

func (d *Delete) Where(cond string, args ...interface{}) *Delete {
	d.Delete.Where(cond, args...)
	return d
}

func (d *Delete) And(cond string, args ...interface{}) *Delete {
	return d.Where(cond, args...)
}

func (d *Delete) Exec() (sql.Result, error) {
	q, b := d.Delete.Build()
	return d.exec.Exec(q, b...)
}
