package gen

var scope = &Template{
	Name: "Scope",
	Text: `
func (m {{.Recv.Name}}) {{.Func}}(args ...interface{}) *{{.Recv.Name}}Relation {
        r := m.newRelation()
        m.{{.FuncName}}(ar.Scope{r.Relation, args})
        return r
}

func (r *{{.Recv.Name}}Relation) {{.Func}}(args ...interface{}) *{{.Recv.Name}}Relation {
        r.src.{{.FuncName}}(ar.Scope{r.Relation, args})
        return r
}
`}
