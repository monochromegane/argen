package tests

import (
	"log"
	"os"

	"github.com/monochromegane/argen"
)

var logger = &ar.Logger{Logger: log.New(os.Stdout, "", 0)}

func LogMode(mode bool) {
	logger.LogMode = mode
}
