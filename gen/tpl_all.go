package gen

var all = &Template{
	Name: "All",
	Text: `
func (m {{.Name}}) All() *{{.Name}}Relation {
	return m.newRelation().All()
}

func (r *{{.Name}}Relation) All() *{{.Name}}Relation {
	return r
}
`}
