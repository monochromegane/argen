package gen

var relation = &Template{
	Name: "Relation",
	Text: `
type {{.Name}}Relation struct {
	src *{{.Name}}
	*ar.Relation
}

func (m *{{.Name}}) newRelation() *{{.Name}}Relation {
        r := &ar.Relation{}
        r.Table("{{.TableName}}").Columns({{range .Fields}}
		"{{.ColumnName}}",{{end}}
	)
	{{if .DefaultScope}}m.defaultScope(ar.Scope{r, nil}){{end}}
        return &{{.Name}}Relation{m, r}
}
`}
