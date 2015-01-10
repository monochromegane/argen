package goar

import (
	"fmt"
	"io"
	"text/template"
)

func Write(w io.Writer, st *Struct) {

	tplText := `
package {{.Package}}

import "github.com/monochromegane/goar"

type {{.Name}}Relation struct {
	db *goar.DB
	*goar.Select
}
`
	tpl := template.Must(template.New("t").Parse(tplText))
	if err := tpl.Execute(w, st); err != nil {
		fmt.Println(err)
	}

}
