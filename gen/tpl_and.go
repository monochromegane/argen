package gen

var and = &Template{
	Name: "And",
	Text: `
func (r *{{.Name}}Relation) And(cond string, args ...interface{}) *{{.Name}}Relation {
        r.Relation.And(cond, args...)
        return r
}
`}
