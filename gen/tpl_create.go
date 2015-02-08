package gen

var create = &Template{
	Name: "Create",
	Text: `
type {{.Name}}Params {{.Name}}

func (m {{.Name}}) Create(p {{.Name}}Params) (*{{.Name}}, *ar.Errors) {
        n := &{{.Name}}{ {{range .Fields}}
		{{.Name}}: p.{{.Name}},{{end}}
        }
        _, errs := n.Save()
        return n, errs
}
`}
