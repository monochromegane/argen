package gen

var save = &Template{
	Name: "Save",
	Text: `
func (m *{{.Name}}) IsNewRecord() bool {
        return ar.IsZero(m.{{.PrimaryKeyField}})
}

func (m *{{.Name}}) IsPersistent() bool {
        return !m.IsNewRecord()
}

func (m *{{.Name}}) Save() (bool, *ar.Errors) {
	if ok, errs := m.IsValid(); !ok {
		return false, errs
	}
	errs := &ar.Errors{}
        if m.IsNewRecord() {
		ins := ar.NewInsert()
                q, b := ins.Table("{{.TableName}}").Params(map[string]interface{}{ {{range .FieldsWithoutPrimaryKey}}
			"{{.ColumnName}}": m.{{.Name}},{{end}}
                }).Build()

                if _, err := db.Exec(q, b...); err != nil {
			errs.Add("base", err)
                        return false, errs
                }
                return true, nil
        }else{
                upd := ar.NewUpdate()
                q, b := upd.Table("{{.TableName}}").Params(map[string]interface{}{ {{range .Fields}}
			"{{.ColumnName}}": m.{{.Name}},{{end}}
                }).Where("{{.PrimaryKeyColumn}}", m.{{.PrimaryKeyField}}).Build()

                if _, err := db.Exec(q, b...); err != nil {
			errs.Add("base", err)
                        return false, errs
                }
                return true, nil
        }
}
`}
