package gen

var queryRow = &Template{
	Name: "QueryRow",
	Text: `
func (r *{{.Name}}Relation) QueryRow() (*{{.Name}}, error) {
	row := &{{.Name}}{}
	err := r.Relation.QueryRow(row.fieldPtrsByName(r.Relation.GetColumns())...)
	if err != nil {
		return nil, err
	}
	return row, nil
}
`}
