package gen

var last = &Template{
	Name: "Last",
	Text: `
func (r *{{.Name}}Relation) Last() (*{{.Name}}, error) {
        return r.Order("{{.PrimaryKeyColumn}}", "DESC").Limit(1).QueryRow()
}
`}
