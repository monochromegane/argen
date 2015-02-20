package gen

var build = &Template{
	Name: "Build",
	Text: `
type {{.Name}}Params {{.Name}}

func (m {{.Name}}) Build(p {{.Name}}Params) *{{.Name}} {
        return &{{.Name}}{ {{range .Fields}}
		{{.Name}}: p.{{.Name}},{{end}}
        }
}
`}
