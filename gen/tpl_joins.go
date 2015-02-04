package gen

var joins = &Template{
	Name: "Joins",
	Text: `
func (m {{.Recv.Name}}) Joins{{.Func}}() *{{.Recv.Name}}Relation {
        r := m.newRelation()
	r.Relation.InnerJoin("{{.TableName}}", "{{.TableName}}.{{.PrimaryKey}}", "{{.ForeignKey}}")
        return r
}

func (r *{{.Recv.Name}}Relation) Joins{{.Func}}() *{{.Recv.Name}}Relation {
	r.Relation.InnerJoin("{{.TableName}}", "{{.TableName}}.{{.PrimaryKey}}", "{{.ForeignKey}}")
        return r
}
`}
