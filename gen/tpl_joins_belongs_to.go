package gen

var joinsBelongsTo = &Template{
	Name: "JoinsBelongsTo",
	Text: `
func (m {{.Recv.Name}}) Joins{{.Func}}() *{{.Recv.Name}}Relation {
        return m.newRelation().Joins{{.Func}}()
}

func (r *{{.Recv.Name}}Relation) Joins{{.Func}}() *{{.Recv.Name}}Relation {
	asc := r.src.{{.FuncName}}()
	pk := "{{.PrimaryKey}}"
	fk := "{{.ForeignKey}}"
	if asc != nil && asc.PrimaryKey != "" {
		pk = asc.PrimaryKey
	}
	if asc != nil && asc.ForeignKey != "" {
		fk = asc.ForeignKey
	}
	r.Relation.InnerJoin("{{.TableName}}", fmt.Sprintf("{{.TableName}}.%s = {{.Recv.TableName}}.%s", pk, fk))
        return r
}
`}
