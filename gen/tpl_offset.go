package gen

var offset = &Template{
	Name: "Offset",
	Text: `
func (m {{.Name}}) Offset(offset int) *{{.Name}}Relation {
	return m.newRelation().Offset(offset)
}

func (r *{{.Name}}Relation) Offset(offset int) *{{.Name}}Relation {
        r.Relation.Offset(offset)
        return r
}
`}
