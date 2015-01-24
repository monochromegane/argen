package gen

var group = &Template{
	Name: "Group",
	Text: `
func (r *{{.Name}}Relation) Group(group string, groups ...string) *{{.Name}}Relation {
        r.Relation.GroupBy(group, groups...)
        return r
}
`}
