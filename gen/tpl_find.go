package gen

var find = &Template{
	Name: "Find",
	Text: `
func (m {{.Name}}) Find(id {{.PrimaryKeyType}}) (*{{.Name}}, error) {
        return m.newRelation().Find(id)
}

func (r *{{.Name}}Relation) Find(id {{.PrimaryKeyType}}) (*{{.Name}}, error) {
        return r.FindBy("{{.PrimaryKeyColumn}}", id)
}
`}
