package gen

var order = &Template{
	Name: "Order",
	Text: `
func (m {{.Name}}) Order(column, order string) *{{.Name}}Relation {
	return m.newRelation().Order(column, order)
}

func (r *{{.Name}}Relation) Order(column, order string) *{{.Name}}Relation {
        r.Relation.OrderBy(column, order)
        return r
}
`}
