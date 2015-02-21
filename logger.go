package ar

import (
	"fmt"
	"log"
	"time"
)

type Logger struct {
	*log.Logger
	LogMode bool
}

func (l Logger) Print(values ...interface{}) {
	if !l.LogMode {
		return
	}

	log := []interface{}{}
	// currentTime
	currentTime := "\033[33m[" + time.Now().Format("2006-01-02 15:04:05") + "]\033[0m"
	// duration
	duration := fmt.Sprintf(" \033[36;1m[%.2fms]\033[0m ", float64(values[0].(time.Duration).Nanoseconds()/1e4)/100.0)
	// sql
	sql := values[1]

	log = append(log, currentTime, duration, sql)

	// args
	if a, ok := values[2].([]interface{}); ok && len(a) > 0 {
		log = append(log, a)
	}

	l.Println(log...)
}
