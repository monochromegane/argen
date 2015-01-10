package goar

import (
	"go/parser"
	"go/token"
)

func Generate(from, to string) error {
	structs, err := parse(from)
	if err != nil {
		return err
	}

	err = writeToFile(to, structs)
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
	return AnotatedStructs(f, "+goar"), nil
}
