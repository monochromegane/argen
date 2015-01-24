package gen

var having = &Template{
	Name: "Having",
	Text: `
func (r *{{.Name}}Relation) Having(cond string, args ...interface{}) *{{.Name}}Relation {
        r.Relation.Having(cond, args...)
        return r
}
`}
