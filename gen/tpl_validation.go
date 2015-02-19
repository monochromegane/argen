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
			errors.SetErrors(name, errs)
                }
        }
	customs := []*ar.Validation{ {{range .CustomValidation}}
		m.{{.FuncName}}().Rule(),{{end}}
	}
	for _, rule := range customs {
		custom := ar.NewValidator(rule).Custom()
		custom(errors)
	}
        return result, errors
}
`}
