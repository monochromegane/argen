package gen

var findBy = &Template{
	Name: "FindBy",
	Text: `
func (m {{.Name}}) FindBy(cond string, args ...interface{}) (*{{.Name}}, error) {
        return m.newRelation().FindBy(cond, args...)
}

func (r *{{.Name}}Relation) FindBy(cond string, args ...interface{}) (*{{.Name}}, error) {
        return r.Where(cond, args...).Limit(1).QueryRow()
}
`}
