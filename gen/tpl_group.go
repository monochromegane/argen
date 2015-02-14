package gen

var group = &Template{
	Name: "Group",
	Text: `
func (m {{.Name}}) Group(group string, groups ...string) *{{.Name}}Relation {
	return m.newRelation().Group(group, groups...)
}

func (r *{{.Name}}Relation) Group(group string, groups ...string) *{{.Name}}Relation {
        r.Relation.GroupBy(group, groups...)
        return r
}
`}
