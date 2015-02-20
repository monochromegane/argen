package tests

import (
	"log"
	"os"
	"time"

	"github.com/monochromegane/argen"
)

var logMode bool
var logger = ar.Logger{log.New(os.Stdout, "", 0)}

func LogMode(mode bool) {
	logMode = mode
}

func Log(t time.Time, sql string, args ...interface{}) {
	if logMode {
		logger.Print(time.Now().Sub(t), sql, args)
	}
}
