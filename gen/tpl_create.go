package gen

var create = &Template{
	Name: "Create",
	Text: `
func (m {{.Name}}) Create(p {{.Name}}Params) (*{{.Name}}, *ar.Errors) {
	n := m.Build(p)
        _, errs := n.Save()
        return n, errs
}
`}
