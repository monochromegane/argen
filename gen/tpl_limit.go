package gen

var limit = &Template{
	Name: "Limit",
	Text: `
func (m {{.Name}}) Limit(limit int) *{{.Name}}Relation {
	return m.newRelation().Limit(limit)
}

func (r *{{.Name}}Relation) Limit(limit int) *{{.Name}}Relation {
        r.Relation.Limit(limit)
        return r
}
`}
