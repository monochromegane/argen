package gen

var first = &Template{
	Name: "First",
	Text: `
func (m {{.Name}}) First() (*{{.Name}}, error) {
	return m.newRelation().First()
}

func (r *{{.Name}}Relation) First() (*{{.Name}}, error) {
        return r.Order("{{.PrimaryKeyColumn}}", "ASC").Limit(1).QueryRow()
}
`}
