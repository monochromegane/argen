package gen

var sel = &Template{
	Name: "Select",
	Text: `
func (m {{.Name}}) Select(columns ...string) *{{.Name}}Relation {
        r := m.newRelation()
        r.Relation.Columns(columns...)
        return r
}
`}
