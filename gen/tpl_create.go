package gen

var create = &Template{
	Name: "Create",
	Text: `
type {{.Name}}Params {{.Name}}

func (m {{.Name}}) Create(p {{.Name}}Params) (*{{.Name}}, error) {
        n := &{{.Name}}{ {{range .Fields}}
		{{.Name}}: p.{{.Name}},{{end}}
        }
        err := n.Save()
        return n, err
}
`}
