package gen

var update = &Template{
	Name: "Update",
	Text: `
func (m *{{.Name}}) Update(p {{.Name}}Params) (bool, *ar.Errors) {
{{range .Fields}}
	if !ar.IsZero(p.{{.Name}}) {
                m.{{.Name}} = p.{{.Name}}
        }{{end}}
	return m.Save()
}

func (m *{{.Name}}) UpdateColumns(p {{.Name}}Params) (bool, *ar.Errors) {
{{range .Fields}}
	if !ar.IsZero(p.{{.Name}}) {
                m.{{.Name}} = p.{{.Name}}
        }{{end}}
	return m.Save(false)
}
`}
