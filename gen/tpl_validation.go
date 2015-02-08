package gen

var validation = &Template{
	Name: "Validation",
	Text: `
func (m {{.Name}}) IsValid() (bool, *ar.Errors) {
        result := true
	errors := &ar.Errors{}
        rules := map[string]*ar.Validation{ {{range .Validation}}
		"{{.ColumnName}}": m.{{.FuncName}}().Rule(),{{end}}
        }
        for name, rule:= range rules {
                if ok, errs := ar.NewValidator(rule).IsValid(m.fieldByName(name)); !ok {
                        result = false
			errors.Set(name, errs)
                }
        }
        return result, errors
}
`}
