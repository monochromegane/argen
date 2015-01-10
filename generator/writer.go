package generator

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"
	"text/template"
	"unicode"
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

	const tplText = `package {{.Package}}

import "github.com/monochromegane/goar"
{{range .}}
func {{.Name | capitalize}}(db *goar.DB) *{{.Name}}Relation {
	sel := &goar.Select{}
	sel.Table("{{.Name}}").Columns({{.FieldNames | join}})
	return &{{.Name}}Relation{db, sel}
}
{{end}}
`
	t := template.New("t")
	t.Funcs(template.FuncMap{
		"capitalize": capitalize,
		"join":       join,
	})
	tpl := template.Must(t.Parse(tplText))
	if err := tpl.Execute(w, structs); err != nil {
		return err
	}
	return nil
}

func capitalize(s string) string {
	c := []rune(s)
	c[0] = unicode.ToUpper(c[0])
	return string(c)
}

func join(s []string) string {
	return fmt.Sprintf("\"%s\"", strings.Join(s, "\", \""))
}
