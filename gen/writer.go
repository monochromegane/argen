package gen

import (
	"bufio"
	"bytes"
	"io"
	"os"
	"text/template"

	"golang.org/x/tools/imports"
)

func writeToFile(file string, structs structs) error {
	f, err := os.Create(file)
	if err != nil {
		return err
	}

	w := bufio.NewWriter(f)
	defer w.Flush()

	b, err := writeWithFormat(file, structs)
	if err != nil {
		return err
	}

	w.Write(b)
	return nil
}

func writeWithFormat(file string, structs structs) ([]byte, error) {
	var b bytes.Buffer
	w := bufio.NewWriter(&b)

	write(w, structs)
	w.Flush()

	formatted, err := imports.Process(file, b.Bytes(), nil)
	if err != nil {
		return nil, err
	}
	return formatted, nil
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
{{template "Select" .}}
{{template "Find" .}}
{{template "FindBy" .}}
{{template "First" .}}
{{template "Last" .}}
{{template "Where" .}}
{{template "And" .}}
{{template "Order" .}}
{{template "Limit" .}}
{{template "Offset" .}}
{{template "Group" .}}
{{template "Having" .}}
{{template "Explain" .}}
{{template "Validation" .}}
{{range .Scope}}
{{template "Scope" .}}
{{end}}
{{range .HasMany}}
{{template "HasMany" .}}
{{end}}
{{range .HasOne}}
{{template "HasOne" .}}
{{end}}
{{range .BelongsTo}}
{{template "BelongsTo" .}}
{{end}}
{{range .Joins}}
{{template "Joins" .}}
{{end}}
{{template "Create" .}}
{{template "Save" .}}
{{template "Delete" .}}
{{template "Query" .}}
{{template "QueryRow" .}}
{{template "FieldByName" .}}
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
