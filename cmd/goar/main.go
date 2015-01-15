package main

import (
	"fmt"
	"os"

	"github.com/jessevdk/go-flags"
	"github.com/monochromegane/goar/gen"
)

var opts gen.Option

func main() {

	args, err := flags.Parse(&opts)
	if err != nil {
		os.Exit(1)
	}

	from := os.Getenv("GOFILE")
	if from == "" {
		if len(args) > 0 {
			from = args[0]
		} else {
			os.Exit(1)
		}
	}

	err = gen.Generate(from, opts)
	if err != nil {
		fmt.Printf("%v\n", err)
	}

}
