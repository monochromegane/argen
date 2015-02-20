package gen

var joinsHasAny = &Template{
	Name: "JoinsHasAny",
	Text: `
func (m {{.Recv.Name}}) Joins{{.Func}}() *{{.Recv.Name}}Relation {
        return m.newRelation().Joins{{.Func}}()
}

func (r *{{.Recv.Name}}Relation) Joins{{.Func}}() *{{.Recv.Name}}Relation {
	asc := r.src.{{.FuncName}}()
	fk := "{{.ForeignKey}}"
	if asc != nil && asc.ForeignKey != "" {
		fk = asc.ForeignKey
	}
	r.Relation.InnerJoin("{{.TableName}}", fmt.Sprintf("{{.TableName}}.%s = {{.Recv.TableName}}.{{.Recv.PrimaryKeyColumn}}", fk))
        return r
}
`}
