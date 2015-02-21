package ar

import (
	"database/sql"

	"github.com/monochromegane/argen/query"
)

type Insert struct {
	*query.Insert
	exec *Executer
}

func NewInsert(db *sql.DB, logger *Logger) *Insert {
	return &Insert{
		Insert: &query.Insert{},
		exec:   &Executer{db, logger},
	}
}

func (i *Insert) Table(table string) *Insert {
	i.Insert.Table(table)
	return i
}

func (i *Insert) Params(params map[string]interface{}) *Insert {
	i.Insert.Params(params)
	return i
}

func (i *Insert) Exec() (sql.Result, error) {
	q, b := i.Insert.Build()
	return i.exec.Exec(q, b...)
}
