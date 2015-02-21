package gen

var sel = &Template{
	Name: "Select",
	Text: `
func (m {{.Name}}) Select(columns ...string) *{{.Name}}Relation {
	return m.newRelation().Select(columns...)
}

func (r *{{.Name}}Relation) Select(columns ...string) *{{.Name}}Relation {
	cs := []string{}
	for _, c := range columns {
		if r.src.isColumnName(c) {
			cs = append(cs, fmt.Sprintf("{{.TableName}}.%s", c))
		} else {
			cs = append(cs, c)	
		}
	}
	r.Relation.Columns(cs...)
	return r
}
`}
