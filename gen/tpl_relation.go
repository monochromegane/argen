package gen

var relation = &Template{
	Name: "Relation",
	Text: `
type {{.Name}}Relation struct {
	src *{{.Name}}
	*ar.Relation
}

func (m *{{.Name}}) newRelation() *{{.Name}}Relation {
	r := &{{.Name}}Relation{
		m,
		ar.NewRelation(db, logger).Table("{{.TableName}}"),
	}
	r.Select({{range .Fields}}
		"{{.ColumnName}}",{{end}}
	)
	{{if .DefaultScope}}m.defaultScope(ar.Scope{r.Relation, nil}){{end}}
	return r
}
`}
