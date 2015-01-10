package main

import (
	"fmt"
	"os"

	"github.com/jessevdk/go-flags"
	"github.com/monochromegane/goar/generator"
)

var opts generator.Option

func main() {

	args, err := flags.Parse(&opts)
	if err != nil {
		os.Exit(1)
	}

	from := os.Getenv("GOFILE")
	if from == "" && len(args) > 0 {
		from = args[0]
	} else {
		os.Exit(1)
	}

	err = generator.Generate(from, "gen.txt")
	if err != nil {
		fmt.Printf("%v\n", err)
	}

}
