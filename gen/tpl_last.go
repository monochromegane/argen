package gen

var last = &Template{
	Name: "Last",
	Text: `
func (m {{.Name}}) Last() (*{{.Name}}, error) {
        return m.newRelation().Last()
}

func (r *{{.Name}}Relation) Last() (*{{.Name}}, error) {
        return r.Order("{{.PrimaryKeyColumn}}", "DESC").Limit(1).QueryRow()
}
`}
