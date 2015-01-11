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
	sel.Table("{{.Name}}").Columns({{.ColumnNames | joinColumn}})
	return &{{.Name}}Relation{db, sel}
}

type {{.Name}}Relation struct {
	db *goar.DB
	*goar.Select
}

func (r *{{.Name}}Relation) Query() ([]*{{.Name}}, error) {
        q, b := r.Build()
        rows, err := r.db.Query(q, b...)
        if err != nil {
                return nil, err
        }
        defer rows.Close()

	results := []*{{.Name}}{}
        for rows.Next() {
                row := &{{.Name}}{}
		if err := rows.Scan({{.FieldNames "&row."| joinField}}); err != nil {
                        return nil, err
                }
                results = append(results, row)
        }
        return results, nil
}

func (r *{{.Name}}Relation) Find(id {{.PrimaryKeyType}}) (*{{.Name}}, error) {
        q, b := r.Select.Where("{{.PrimaryKey}}", id).Build()
        row := &{{.Name}}{}
        if err := r.db.QueryRow(q, b...).Scan({{.FieldNames "&row."| joinField}}); err != nil {
                return nil, err
        }
        return row, nil
}

func (r *{{.Name}}Relation) Where(cond string, args ...interface{}) *{{.Name}}Relation {
        r.Select.Where(cond, args...)
        return r
}

func (r *{{.Name}}Relation) And(cond string, args ...interface{}) *{{.Name}}Relation {
        r.Select.And(cond, args...)
        return r
}
{{end}}
`
	t := template.New("t")
	t.Funcs(template.FuncMap{
		"capitalize": capitalize,
		"joinColumn": joinColumn,
		"joinField":  joinField,
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

func joinColumn(s []string) string {
	return fmt.Sprintf("\"%s\"", strings.Join(s, "\", \""))
}

func joinField(s []string) string {
	return strings.Join(s, ", ")
}
