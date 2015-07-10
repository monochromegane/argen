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

func (m *{{.Name}}) Save(validate ...bool) (bool, *ar.Errors) {
	if len(validate) == 0 || len(validate) > 0 && validate[0] {
		if ok, errs := m.IsValid(); !ok {
			return false, errs
		}
	}
	errs := &ar.Errors{}
        if m.IsNewRecord() {
                ins := ar.NewInsert(db, logger).Table("{{.TableName}}").Params(map[string]interface{}{ {{range .FieldsWithoutPrimaryKey}}
			"{{.ColumnName}}": m.{{.Name}},{{end}}
                })

                if result, err := ins.Exec(); err != nil {
			errs.AddError("base", err)
                        return false, errs
                } else {
			if lastId, err := result.LastInsertId(); err == nil {
				m.{{.PrimaryKeyField}} = {{.PrimaryKeyType}}(lastId)
			}
		}
                return true, nil
        }else{
		upd := ar.NewUpdate(db, logger).Table("{{.TableName}}").Params(map[string]interface{}{ {{range .Fields}}
		"{{.ColumnName}}": m.{{.Name}},{{end}}
		}).Where("{{.PrimaryKeyColumn}}", m.{{.PrimaryKeyField}})

		if _, err := upd.Exec(); err != nil {
			errs.AddError("base", err)
			return false, errs
		}
                return true, nil
        }
}
`}
