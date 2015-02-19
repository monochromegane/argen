package gen

var update = &Template{
	Name: "Update",
	Text: `
func (m *{{.Name}}) Update(p {{.Name}}Params) (bool, *ar.Errors) {
        if ok, errs := m.IsValid(); !ok {
                return false, errs
        }
        return m.UpdateColumns(p)
}

func (m *{{.Name}}) UpdateColumns(p {{.Name}}Params) (bool, *ar.Errors) {
        errs := &ar.Errors{}
        params := map[string]interface{}{}

{{range .Fields}}
        if !ar.IsZero(p.{{.Name}}) && m.{{.Name}} != p.{{.Name}} {
                params["{{.ColumnName}}"] = p.{{.Name}}
        }
{{end}}

        if _, err := m.updateColumnsByMap(params); err != nil {
                errs.AddError("base", err)
                return false, errs
        }
        return true, nil
}

func (m *{{.Name}}) updateColumnsByMap(params map[string]interface{}) (bool, error) {
        upd := ar.NewUpdate()
        q, b := upd.Table("{{.TableName}}").Params(params).Where("{{.PrimaryKeyColumn}}", m.{{.PrimaryKeyField}}).Build()
        if _, err := db.Exec(q, b...); err != nil {
                return false, err
        }
        return true, nil
}
`}
