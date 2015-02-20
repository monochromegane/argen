package ar

import (
	"fmt"
	"log"
	"time"
)

type Logger struct {
	*log.Logger
}

func (l Logger) Print(values ...interface{}) {
	// currentTime
	currentTime := "\033[33m[" + time.Now().Format("2006-01-02 15:04:05") + "]\033[0m"
	// duration
	duration := fmt.Sprintf(" \033[36;1m[%.2fms]\033[0m ", float64(values[0].(time.Duration).Nanoseconds()/1e4)/100.0)
	// sql
	sql := values[1]
	args := values[2]

	l.Println(currentTime, duration, sql, args)
}
