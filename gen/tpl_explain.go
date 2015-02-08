package gen

var explain = &Template{
	Name: "Explain",
	Text: `
func (r *{{.Name}}Relation) Explain() *{{.Name}}Relation {
        r.Relation.Explain()
        return r
}
`}
