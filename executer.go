package ar

import (
	"database/sql"
	"time"
)

type Executer struct {
	db     *sql.DB
	logger *Logger
}

func (e *Executer) Exec(q string, b ...interface{}) (sql.Result, error) {
	defer e.log(time.Now(), q, b...)
	return e.db.Exec(q, b...)
}

func (e *Executer) log(t time.Time, sql string, args ...interface{}) {
	e.logger.Print(time.Now().Sub(t), sql, args)
}
