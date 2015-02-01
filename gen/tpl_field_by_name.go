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
`}
