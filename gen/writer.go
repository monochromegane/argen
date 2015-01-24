package gen

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

	tplText := `package {{.Package}}

import (
	"database/sql"

	"github.com/monochromegane/argen"
)

var db *sql.DB

func Use(DB *sql.DB) {
	db = DB
}
{{range .}}
{{template "Relation" .}}
{{end}}
` + templates.ToString()

	t := template.New("t")
	t.Funcs(template.FuncMap{})
	tpl := template.Must(t.Parse(tplText))
	if err := tpl.Execute(w, structs); err != nil {
		return err
	}

	return nil
}
