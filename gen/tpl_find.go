package gen

var find = &Template{
	Name: "Find",
	Text: `
func (m {{.Name}}) Find(id {{.PrimaryKeyType}}) (*{{.Name}}, error) {
        return m.newRelation().Where("{{.PrimaryKeyColumn}}", id).QueryRow()
}
`}
