package ar

import (
	"database/sql"

	"github.com/monochromegane/argen/query"
)

type Update struct {
	*query.Update
	exec *Executer
}

func NewUpdate(db *sql.DB, logger *Logger) *Update {
	return &Update{
		Update: &query.Update{},
		exec:   &Executer{db, logger},
	}
}

func (u *Update) Table(table string) *Update {
	u.Update.Table(table)
	return u
}

func (u *Update) Where(cond string, args ...interface{}) *Update {
	u.Update.Where(cond, args...)
	return u
}

func (u *Update) And(cond string, args ...interface{}) *Update {
	return u.Where(cond, args...)
}

func (u *Update) Params(params map[string]interface{}) *Update {
	u.Update.Params(params)
	return u
}

func (u *Update) Exec() (sql.Result, error) {
	q, b := u.Update.Build()
	return u.exec.Exec(q, b...)
}
