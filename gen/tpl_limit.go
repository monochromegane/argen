package gen

var limit = &Template{
	Name: "Limit",
	Text: `
func (r *{{.Name}}Relation) Limit(limit int) *{{.Name}}Relation {
        r.Relation.Limit(limit)
        return r
}
`}
