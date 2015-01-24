package gen

var queryRow = &Template{
	Name: "QueryRow",
	Text: `
func (r *{{.Name}}Relation) QueryRow() (*{{.Name}}, error) {
	q, b := r.Build()
	row := &{{.Name}}{}
	err := db.QueryRow(q, b...).Scan({{range .Fields}}
		&row.{{.Name}},{{end}}
	)
	if err != nil {
		return nil, err
	}
	return row, nil
}
`}
