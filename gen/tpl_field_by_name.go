package gen

var fieldByName = &Template{
	Name: "FieldByName",
	Text: `
func (m *{{.Name}}) fieldValueByName(name string) interface{} {
        switch name { {{range .Fields}}
	case "{{.ColumnName}}":
		return m.{{.Name}}{{end}}
        default:
                return ""
        }
}

func (m *{{.Name}}) fieldPtrByName(name string) interface{} {
        switch name { {{range .Fields}}
	case "{{.ColumnName}}":
		return &m.{{.Name}}{{end}}
        default:
                return nil
        }
}

func (m *{{.Name}}) fieldPtrsByName(names []string) []interface{} {
        fields := []interface{}{}
        for _, n := range names {
                f := m.fieldPtrByName(n)
                fields = append(fields, f)
        }
        return fields
}
`}
