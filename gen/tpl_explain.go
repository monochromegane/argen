package gen

var explain = &Template{
	Name: "Explain",
	Text: `
func (r *{{.Name}}Relation) Explain() error {
        rows, err := r.Relation.Explain().Query()
        if err != nil {
                return err
        }
        defer rows.Close()

        columns, _ := rows.Columns()
        var values [][]string
        for rows.Next() {
                vals := make([]string, len(columns))
                ptrs := make([]interface{}, len(columns))
                for i, _ := range vals {
                        ptrs[i] = &vals[i]
                }
                rows.Scan(ptrs...)
                values = append(values, vals)
        }

        goban.Render(columns, values)
        return nil
}
`}
