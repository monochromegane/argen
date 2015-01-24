package gen

var order = &Template{
	Name: "Order",
	Text: `
func (r *{{.Name}}Relation) Order(column, order string) *{{.Name}}Relation {
        r.Relation.OrderBy(column, order)
        return r
}
`}
