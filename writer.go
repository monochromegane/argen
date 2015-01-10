package goar

import (
	"bufio"
	"io"
	"os"
	"text/template"
)

func writeToFile(file string, structs structs) error {
	f, err := os.Create(file)
	if err != nil {
		return err
	}

	w := bufio.NewWriter(f)
	defer w.Flush()

	return write(w, structs)
}

func write(w io.Writer, structs structs) error {

	const tplText = `package {{.pkg}}
`
	tpl := template.Must(template.New("t").Parse(tplText))
	if err := tpl.Execute(w, structs); err != nil {
		return err
	}
	return nil
}
