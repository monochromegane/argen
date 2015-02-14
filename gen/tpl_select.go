package gen

var sel = &Template{
	Name: "Select",
	Text: `
func (m {{.Name}}) Select(columns ...string) *{{.Name}}Relation {
	return m.newRelation().Select(columns...)
}

func (r *{{.Name}}Relation) Select(columns ...string) *{{.Name}}Relation {
	r.Relation.Columns(columns...)
	return r
}
`}
