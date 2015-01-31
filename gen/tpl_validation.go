package gen

var validation = &Template{
	Name: "Validation",
	Text: `
func (m {{.Name}}) IsValid() bool {
        result := true
        rules := map[*ar.Validation]interface{}{ {{range .Validation}}
                m.{{.FuncName}}().Rule(): m.{{.FieldName}},{{end}}
        }
        for rule, value := range rules {
                if !ar.NewValidator(rule).IsValid(value) {
                        result = false
                }
        }
        return result
}
`}
