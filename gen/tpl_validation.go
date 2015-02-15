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
                if ok, errs := ar.NewValidator(rule).IsValid(m.fieldValueByName(name)); !ok {
                        result = false
			errors.Set(name, errs)
                }
        }
	customs := []ar.CustomValidator{ {{range .CustomValidation}}
		m.{{.FuncName}},{{end}}
	}
	for _, c := range customs {
		if ok, column, err := c(); !ok {
			result = false
			errors.Add(column, err)
		}
	}
        return result, errors
}
`}
