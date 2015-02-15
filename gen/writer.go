package gen

import (
	"bufio"
	"bytes"
	"io"
	"os"
	"text/template"

	"golang.org/x/tools/imports"
)

func writeToFile(file, template string, structs structs) error {
	f, err := os.Create(file)
	if err != nil {
		return err
	}

	w := bufio.NewWriter(f)
	defer w.Flush()

	b, err := writeWithFormat(file, template, structs)
	if err != nil {
		return err
	}

	w.Write(b)
	return nil
}

func writeWithFormat(file, template string, structs structs) ([]byte, error) {
	var b bytes.Buffer
	w := bufio.NewWriter(&b)

	write(w, template, structs)
	w.Flush()

	formatted, err := imports.Process(file, b.Bytes(), nil)
	if err != nil {
		return nil, err
	}
	return formatted, nil
}

func write(w io.Writer, tplText string, structs structs) error {
	t := template.New("t")
	t.Funcs(template.FuncMap{})
	tpl := template.Must(t.Parse(tplText))
	if err := tpl.Execute(w, structs); err != nil {
		return err
	}

	return nil
}
