package generator

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"
	"text/template"
	"unicode"

	"github.com/gedex/inflector"
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

import (
	"database/sql"

	"github.com/monochromegane/goar"
)

var db *sql.DB

func Use(DB *sql.DB) {
	db = DB
}
{{range .}}
func (m {{.Name}}) Find(id {{.PrimaryKeyType}}) (*{{.Name}}, error) {
	r := m.newRelation()
        q, b := r.Select.Where("{{.PrimaryKeyColumn}}", id).Build()
        row := &{{.Name}}{}
        if err := db.QueryRow(q, b...).Scan({{.FieldNames "&row."| joinField}}); err != nil {
                return nil, err
        }
        return row, nil
}

func (m {{.Name}}) Where(cond string, args ...interface{}) *{{.Name}}Relation {
	r := m.newRelation()
        return r.Where(cond, args...)
}

func (m *{{.Name}}) newRelation() *{{.Name}}Relation {
	sel := &goar.Select{}
	sel.Table("{{.Name}}").Columns({{.ColumnNames | joinColumn}})
	return &{{.Name}}Relation{sel}
}

type {{.Name}}Relation struct {
	*goar.Select
}

func (r *{{.Name}}Relation) Query() ([]*{{.Name}}, error) {
        q, b := r.Build()
        rows, err := db.Query(q, b...)
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

func (r *{{.Name}}Relation) First() (*{{.Name}}, error) {
	q, b := r.OrderBy("{{.PrimaryKeyColumn}}", goar.ASC).Limit(1).Build()
        row := &{{.Name}}{}
        if err := db.QueryRow(q, b...).Scan({{.FieldNames "&row."| joinField}}); err != nil {
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
{{$model := .}}
{{range .Anotations}}
{{if .BelongsTo}}
func (m *{{$model.Name}}) {{.Arg | capitalize}}() (*{{.Arg | capitalize}}, error) {
	return {{.Arg | capitalize}}{}.Where("{{$model.PrimaryKeyColumn}}", m.{{.Arg | capitalize}}ID).First()
}
{{else if .HasOne}}
func (m *{{$model.Name}}) {{.Arg | capitalize}}() ([]*{{.Arg | capitalize}}, error) {
	return {{.Arg | capitalize | singularize}}{}.Where("{{$model.TableName}}_id", m.{{$model.PrimaryKeyField}}).First()
}
{{else if .HasMany}}
func (m *{{$model.Name}}) {{.Arg | capitalize}}() ([]*{{.Arg | capitalize | singularize}}, error) {
	return {{.Arg | capitalize | singularize}}{}.Where("{{$model.TableName}}_id", m.{{$model.PrimaryKeyField}}).Query()
}
{{end}}
{{end}}
{{end}}
`
	t := template.New("t")
	t.Funcs(template.FuncMap{
		"capitalize":  capitalize,
		"joinColumn":  joinColumn,
		"joinField":   joinField,
		"singularize": inflector.Singularize,
		"pluralize":   inflector.Pluralize,
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
