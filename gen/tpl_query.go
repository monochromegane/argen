package gen

var query = &Template{
	Name: "Query",
	Text: `
func (r *{{.Name}}Relation) Query() ([]*{{.Name}}, error) {
        q, b := r.Build()
	defer Log(time.Now(), q, b...)
        rows, err := db.Query(q, b...)
        if err != nil {
                return nil, err
        }
        defer rows.Close()

        results := []*{{.Name}}{}
        for rows.Next() {
                row := &{{.Name}}{}
		err := rows.Scan(row.fieldPtrsByName(r.Relation.GetColumns())...)
                if err != nil {
                        return nil, err
                }
                results = append(results, row)
        }
        return results, nil
}
`}
