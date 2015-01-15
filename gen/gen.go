package gen

import (
	"go/parser"
	"go/token"
	"strings"
)

func Generate(from string, opts Option) error {
	structs, err := parse(from)
	if err != nil {
		return err
	}

	err = writeToFile(toFileName(from, opts.Prefix, opts.Suffix), structs)
	if err != nil {
		return err
	}

	return nil
}

func parse(file string) (structs, error) {
	fset := token.NewFileSet()
	f, err := parser.ParseFile(fset, file, nil, parser.ParseComments)
	if err != nil {
		return nil, err
	}
	return AnotatedStructs(f, "+AR"), nil
}

func toFileName(from, prefix, suffix string) string {
	names := strings.Split(from, ".")
	return prefix + names[0] + suffix + ".go"
}
