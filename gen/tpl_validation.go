package gen

var validation = &Template{
	Name: "Validation",
	Text: `
func (m {{.Name}}) IsValid() (bool, error) {
        result := true
	errors := &ar.Errors{}
	var on ar.On
	if m.IsNewRecord() {
		on = ar.OnCreate()
	} else {
		on = ar.OnUpdate()
	}
        rules := map[string]*ar.Validation{ {{range .Validation}}
		"{{.ColumnName}}": m.{{.FuncName}}().Rule(),{{end}}
        }
        for name, rule:= range rules {
                if ok, errs := ar.NewValidator(rule).On(on).IsValid(m.fieldValueByName(name)); !ok {
                        result = false
			errors.SetErrors(name, errs)
                }
        }
	customs := []*ar.Validation{ {{range .CustomValidation}}
		m.{{.FuncName}}().Rule(),{{end}}
	}
	for _, rule := range customs {
		custom := ar.NewValidator(rule).On(on).Custom()
		custom(errors)
	}
	if len(errors.Messages) > 0 {
		result = false
	}
        return result, errors
}
`}
