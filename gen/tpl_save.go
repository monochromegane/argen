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

func (m *{{.Name}}) Save() error {
        if m.IsNewRecord() {
                ins := &ar.Insert{}
                q, b := ins.Table("{{.TableName}}").Params(map[string]interface{}{ {{range .Fields}}
			"{{.ColumnName}}": m.{{.Name}},{{end}}
                }).Build()

                if _, err := db.Exec(q, b...); err != nil {
                        return err
                }
                return nil
        }else{
                upd := &ar.Update{}
                q, b := upd.Table("{{.TableName}}").Params(map[string]interface{}{ {{range .Fields}}
			"{{.ColumnName}}": m.{{.Name}},{{end}}
                }).Where("{{.PrimaryKeyColumn}}", m.{{.PrimaryKeyField}}).Build()

                if _, err := db.Exec(q, b...); err != nil {
                        return err
                }
                return nil
        }
        return nil
}
`}
