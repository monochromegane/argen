package gen

var where = &Template{
	Name: "Where",
	Text: `
func (r *{{.Name}}Relation) Where(cond string, args ...interface{}) *{{.Name}}Relation {
        r.Relation.Where(cond, args...)
        return r
}
`}
