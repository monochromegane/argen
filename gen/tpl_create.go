package gen

var create = &Template{
	Name: "Create",
	Text: `
func (m {{.Name}}) Create(p {{.Name}}Params) (*{{.Name}}, error) {
	n := m.Build(p)
        _, errs := n.Save()
        return n, errs
}
`}
