package gen

var explain = &Template{
	Name: "Explain",
	Text: `
func (r *{{.Name}}Relation) Explain() error {
        r.Relation.Explain()
        q, b := r.Build()
	defer Log(time.Now(), q, b...)
        rows, err := db.Query(q, b...)
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

        fmt.Printf("%s %v\n", q, b)
        goban.Render(columns, values)
        return nil
}
`}
