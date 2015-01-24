package gen

var offset = &Template{
	Name: "Offset",
	Text: `
func (r *{{.Name}}Relation) Offset(offset int) *{{.Name}}Relation {
        r.Relation.Offset(offset)
        return r
}
`}
