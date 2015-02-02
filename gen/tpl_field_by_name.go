package gen

var fieldByName = &Template{
	Name: "FieldByName",
	Text: `
func (m {{.Name}}) fieldByName(name string) interface{} {
        switch name { {{range .Fields}}
	case "{{.ColumnName}}":
		return m.{{.Name}}{{end}}
        default:
                return ""
        }
}

func (m {{.Name}}) fieldsByName(names []string) []interface{} {
        fields := []interface{}{}
        for _, n := range names {
                f := m.fieldByName(n)
                fields = append(fields, &f)
        }
        return fields
}
`}
