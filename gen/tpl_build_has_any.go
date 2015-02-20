package gen

var buildHasAny = &Template{
	Name: "BuildHasAny",
	Text: `
func (m *{{.Recv.Name}}) Build{{.Model}}(p {{.Model}}Params) *{{.Model}} {
	p.{{.ForeignKeyField}} = m.{{.Recv.PrimaryKeyField}}
	return {{.Model}}{}.Build(p)
}
`}
