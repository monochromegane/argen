package gen

var first = &Template{
	Name: "First",
	Text: `
func (r *{{.Name}}Relation) First() (*{{.Name}}, error) {
        return r.Order("{{.PrimaryKeyColumn}}", "ASC").Limit(1).QueryRow()
}
`}
