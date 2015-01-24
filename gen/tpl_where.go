package gen

var where = &Template{
	Name: "Where",
	Text: `
func (m {{.Name}}) Where(cond string, args ...interface{}) *{{.Name}}Relation {
        return m.newRelation().Where(cond, args...)
}

func (r *{{.Name}}Relation) Where(cond string, args ...interface{}) *{{.Name}}Relation {
        r.Relation.Where(cond, args...)
        return r
}
`}
