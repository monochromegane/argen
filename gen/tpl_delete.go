package gen

var delete = &Template{
	Name: "Delete",
	Text: `
func (m *{{.Name}}) Delete() (bool, error) {
        errs := &ar.Errors{}
        if _, err := ar.NewDelete(db, logger).Table("{{.TableName}}").Where("{{.PrimaryKeyColumn}}", m.{{.PrimaryKeyField}}).Exec(); err != nil {
                errs.AddError("base", err)
                return false, errs
        }
        return true, nil
}

func (m {{.Name}}) DeleteAll() (bool, error) {
        errs := &ar.Errors{}
        if _, err := ar.NewDelete(db, logger).Table("{{.TableName}}").Exec(); err != nil {
                errs.AddError("base", err)
                return false, errs
        }
        return true, nil
}
`}
