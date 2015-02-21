package gen

var fieldByName = &Template{
	Name: "FieldByName",
	Text: `
{{$tableName := .TableName}}
func (m *{{.Name}}) fieldValueByName(name string) interface{} {
        switch name { {{range .Fields}}
	case "{{.ColumnName}}", "{{$tableName}}.{{.ColumnName}}":
		return m.{{.Name}}{{end}}
        default:
                return ""
        }
}

func (m *{{.Name}}) fieldPtrByName(name string) interface{} {
        switch name { {{range .Fields}}
	case "{{.ColumnName}}", "{{$tableName}}.{{.ColumnName}}":
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

func (m *{{.Name}}) isColumnName(name string) bool {
	for _, c := range m.columnNames() {
		if c == name {
			return true
		}
	}
	return false
}

func (m *{{.Name}}) columnNames() []string {
        return []string{ {{range .Fields}}
		"{{.ColumnName}}",{{end}}
        }
}
`}
